package godcap

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/automata-network/dcap-sdk/packages/godcap/parser"
	"github.com/automata-network/dcap-sdk/packages/godcap/pccs"
	"github.com/automata-network/dcap-sdk/packages/godcap/stubs/DcapLibCallback"
	gen "github.com/automata-network/dcap-sdk/packages/godcap/stubs/DcapPortal"
	"github.com/automata-network/dcap-sdk/packages/godcap/stubs/IDcapAttestation"
	"github.com/automata-network/dcap-sdk/packages/godcap/zkdcap"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const basefeeMultiplier = 2

var (
	ErrValueShouldBeNil        = logex.Define("value in TransactOpts should be nil")
	ErrTransactOptsMissingFrom = logex.Define("TransactOpts missing from")
	ErrInsuccifientFunds       = logex.Define("InsuccifientFunds")
	DcapError                  = map[string]string{
		"0x1356a63b": "AutomataDcapAttestation: BP_Not_Valid()",
		"0x1a72054d": "AutomataDcapAttestation: Insuccifient_Funds()",
		"0xc40a532b": "AutomataDcapAttestation: Withdrawal_Failed()",
	}
)

type DcapPortal struct {
	client  *ethclient.Client
	Stub    *gen.DcapPortal
	abi     abi.ABI
	dcapAbi abi.ABI
	chain   *ChainConfig
	pccs    *pccs.Server

	zkProof *zkdcap.ZkProofClient
}

func NewDcapPortal(ctx context.Context, endpoint string) (*DcapPortal, error) {
	client, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, logex.Trace(err)
	}

	chainId, err := client.ChainID(ctx)
	if err != nil {
		return nil, logex.Trace(err)
	}
	chain := ChainConfigFromChainId(chainId.Int64())
	if chain == nil {
		return nil, logex.Trace(err)
	}
	return NewDcapPortalFromConfig(ctx, client, chain)
}

func NewDcapPortalFromConfig(ctx context.Context, client *ethclient.Client, chain *ChainConfig) (*DcapPortal, error) {
	stub, err := gen.NewDcapPortal(chain.DcapPortal, client)
	if err != nil {
		return nil, logex.Trace(err)
	}

	portalAbi, err := abi.JSON(strings.NewReader(gen.DcapPortalABI))
	if err != nil {
		return nil, logex.Trace(err)
	}
	libAbi, err := abi.JSON(strings.NewReader(DcapLibCallback.DcapLibCallbackABI))
	if err != nil {
		return nil, logex.Trace(err)
	}
	dcapAbi, err := abi.JSON(strings.NewReader(IDcapAttestation.IDcapAttestationABI))
	if err != nil {
		return nil, logex.Trace(err)
	}
	for name, err := range libAbi.Errors {
		portalAbi.Errors[name] = err
	}

	pccs, err := pccs.NewServer(client, chain.PCCS)
	if err != nil {
		return nil, logex.Trace(err)
	}

	portal := &DcapPortal{
		chain:   chain,
		client:  client,
		abi:     portalAbi,
		dcapAbi: dcapAbi,
		pccs:    pccs,
		Stub:    stub,
	}
	return portal, nil
}

func (d *DcapPortal) Pccs() *pccs.Server {
	return d.pccs
}

func (d *DcapPortal) EnableZkProof(cfg *zkdcap.ZkProofConfig) error {
	client, err := zkdcap.NewZkProofClient(cfg, d.pccs)
	if err != nil {
		return logex.Trace(err)
	}
	d.zkProof = client
	return nil
}

func (p *DcapPortal) BuildTransactOpts(ctx context.Context, key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(p.chain.ChainId))
	if err != nil {
		return nil, logex.Trace(err)
	}
	opts.Context = ctx
	return p.normalizeOpts(opts)
}

func (p *DcapPortal) WaitTx(ctx context.Context, tx *types.Transaction) <-chan *types.Receipt {
	result := make(chan *types.Receipt)
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				receipt, err := p.client.TransactionReceipt(ctx, tx.Hash())
				if err != nil {
					logex.Infof("waiting tx receipt for %v: %v", tx.Hash(), err)
					continue
				}

				logex.Infof("tx receipt %v comfirmed on %v", tx.Hash(), receipt.BlockNumber)
				result <- receipt
				return
			}
		}
	}()

	return result
}

func (p *DcapPortal) VerifyOnChain(opts *bind.TransactOpts, rawQuote []byte, callback *Callback) (*types.Transaction, error) {
	params, err := callback.Abi()
	if err != nil {
		return nil, logex.Trace(err)
	}
	opts, err = p.normalizeOpts(opts)
	if err != nil {
		return nil, logex.Trace(err)
	}
	balance, err := p.client.BalanceAt(opts.Context, opts.From, nil)
	if err != nil {
		return nil, logex.Trace(err)
	}
	feeBase, err := p.EstimateFeeBaseVerifyOnChain(opts.Context, opts.From, balance, rawQuote)
	if err != nil {
		return nil, logex.Trace(err)
	}
	opts.Value = new(big.Int).Add(p.attestationFee(opts, feeBase), params.Value)

	newTx, err := p.Stub.VerifyOnChain(opts, rawQuote, params)
	if err != nil {
		return nil, logex.Trace(p.decodeErr(err))
	}
	return newTx, nil
}

func (p *DcapPortal) GenerateZkProof(ctx context.Context, ty zkdcap.ZkType, quote []byte) (*zkdcap.ZkProof, error) {
	if p.zkProof == nil {
		return nil, logex.NewErrorf("DcapPortal not withZkProofClient")
	}
	parser := parser.NewQuoteParser(quote)
	collateral, err := zkdcap.NewCollateralFromQuoteParser(ctx, parser, p.pccs)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return p.zkProof.ProveQuote(ctx, ty, quote, collateral)
}

func (p *DcapPortal) VerifyAndAttestWithZKProof(opts *bind.TransactOpts, zkProof *zkdcap.ZkProof, callback *Callback) (*types.Transaction, error) {
	params, err := callback.Abi()
	if err != nil {
		return nil, logex.Trace(err)
	}
	opts, err = p.normalizeOpts(opts)
	if err != nil {
		return nil, logex.Trace(err)
	}
	balance, err := p.client.BalanceAt(opts.Context, opts.From, nil)
	if err != nil {
		return nil, logex.Trace(err)
	}
	feeBase, err := p.EstimateFeeBaseVerifyAndAttestWithZKProof(opts.Context, opts.From, balance, zkProof)
	if err != nil {
		return nil, logex.Trace(err)
	}
	opts.Value = new(big.Int).Add(p.attestationFee(opts, feeBase), params.Value)

	newTx, err := p.Stub.VerifyAndAttestWithZKProof(opts, zkProof.Output, uint8(zkProof.Type), zkProof.Proof, params)
	if err != nil {
		return nil, logex.Trace(p.decodeErr(err))
	}
	return newTx, nil
}

// Note: the fee = EstimateFeeBase * GasPrice
func (p *DcapPortal) EstimateFeeBaseVerifyOnChain(ctx context.Context, from common.Address, value *big.Int, rawQuote []byte) (*big.Int, error) {
	result, err := p.callContract(ctx, from, value, "estimateFeeBaseVerifyOnChain", rawQuote)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return result[0].(*big.Int), nil
}

// Note: the fee = EstimateFeeBase * GasPrice
func (p *DcapPortal) EstimateFeeBaseVerifyAndAttestWithZKProof(ctx context.Context, from common.Address, value *big.Int, zkProof *zkdcap.ZkProof) (*big.Int, error) {
	result, err := p.callContract(ctx, from, value, "estimateFeeBaseVerifyAndAttestWithZKProof", zkProof.Output, uint8(zkProof.Type), zkProof.Proof)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return result[0].(*big.Int), nil
}

func (p *DcapPortal) CheckQuote(ctx context.Context, from common.Address, quote []byte) (bool, error) {
	calldata, err := p.dcapAbi.Pack("verifyOnChain", quote)
	if err != nil {
		return false, logex.Trace(err)
	}
	balance, err := p.client.BalanceAt(ctx, from, nil)
	if err != nil {
		return false, logex.Trace(err)
	}
	ret, err := p.client.CallContract(ctx, ethereum.CallMsg{
		To:    &p.chain.AutomataDcapAttestationFee,
		From:  from,
		Data:  calldata,
		Value: balance,
	}, nil)
	if err != nil {
		return false, logex.Trace(err)
	}
	args, err := p.dcapAbi.Unpack("verifyOnChain", ret)
	if err != nil {
		return false, logex.Trace(err)
	}
	return args[0].(bool), nil
}

func (p *DcapPortal) CheckZkProof(ctx context.Context, from common.Address, proof *zkdcap.ZkProof) (bool, error) {
	calldata, err := p.dcapAbi.Pack("verifyAndAttestWithZKProof", proof.Output, proof.Type, proof.Proof)
	if err != nil {
		return false, logex.Trace(err)
	}
	balance, err := p.client.BalanceAt(ctx, from, nil)
	if err != nil {
		return false, logex.Trace(err)
	}
	ret, err := p.client.CallContract(ctx, ethereum.CallMsg{
		To:    &p.chain.AutomataDcapAttestationFee,
		From:  from,
		Data:  calldata,
		Value: balance,
	}, nil)
	if err != nil {
		return false, logex.Trace(err)
	}
	args, err := p.dcapAbi.Unpack("verifyAndAttestWithZKProof", ret)
	if err != nil {
		return false, logex.Trace(err)
	}
	return args[0].(bool), nil
}

func (p *DcapPortal) EstimateAttestationFee(tx *types.Transaction, callback *Callback) *big.Int {
	return new(big.Int).Sub(tx.Value(), callback.Value())
}

func (p *DcapPortal) CalculateAttestationFee(tx *types.Transaction, callback *Callback, receipt *types.Receipt) *big.Int {
	estimateFee := p.EstimateAttestationFee(tx, callback)
	feeBase := new(big.Int).Div(estimateFee, tx.GasFeeCap())
	fee := new(big.Int).Mul(feeBase, receipt.EffectiveGasPrice)
	return fee
}

func (p *DcapPortal) PrintAttestationFee(tx *types.Transaction, callback *Callback, receipt *types.Receipt) {
	fmt.Println("Tx GasPrice:", tx.GasPrice())
	fmt.Println("Callback Value:", callback.Value())
	fmt.Println("Receipt EffectiveGasPrice:", receipt.EffectiveGasPrice)
	estimateFee := p.EstimateAttestationFee(tx, callback)
	fmt.Println("EstimateAttestationFee:", estimateFee)
	feeBase := new(big.Int).Div(estimateFee, tx.GasFeeCap())
	fmt.Println("EstimateFeeBase(over actual value):", feeBase)
	fee := new(big.Int).Mul(feeBase, receipt.EffectiveGasPrice)
	fmt.Println("Estimate EffectiveAttestationFee(over actual value):", fee)
	fmt.Println("Total Sent:", tx.Value())
	refund := new(big.Int).Sub(tx.Value(), callback.Value())
	refund = refund.Sub(refund, fee)
	fmt.Println("Estimate Refund(under actual value):", refund)
}

func (p *DcapPortal) Client() *ethclient.Client {
	return p.client
}

func (p *DcapPortal) callContract(ctx context.Context, from common.Address, value *big.Int, method string, args ...interface{}) ([]interface{}, error) {
	calldata, err := p.abi.Pack(method, args...)
	if err != nil {
		return nil, logex.Trace(err)
	}
	msg := ethereum.CallMsg{
		From:  from,
		To:    &p.chain.DcapPortal,
		Value: value,
		Data:  calldata,
	}
	data, err := p.client.CallContract(ctx, msg, nil)
	if err != nil {
		return nil, logex.Trace(p.decodeErr(err))
	}
	result, err := p.abi.Unpack(method, data)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return result, nil
}

func (p *DcapPortal) decodeErrData(msg string, dataBytes []byte) error {
	sig := dataBytes
	if len(sig) > 4 {
		sig = sig[:4]
	}
	for name, er := range p.abi.Errors {
		if bytes.Equal(er.ID[:4], sig) {
			args, _ := er.Inputs.Unpack(dataBytes[len(sig):])
			for idx := range args {
				if b, ok := args[idx].([]byte); ok {
					args[idx] = hexutil.Bytes(b)
				}
			}
			if name == "CALLBACK_FAILED" {
				return logex.Trace(p.decodeErrData(msg, args[0].(hexutil.Bytes)), "callbackFailed")
			}
			return logex.NewErrorf("%v: %v(%v)", msg, name, args)
		}
	}
	return logex.NewErrorf("%v: %v", msg, hexutil.Bytes(dataBytes))
}

func (p *DcapPortal) decodeErr(err error) error {
	if err == nil {
		return nil
	}

	jerr, ok := err.(JsonError)
	if !ok {
		return err
	}

	data, ok := jerr.ErrorData().(string)
	if !ok {
		return err
	}
	if sig, ok := DcapError[data]; ok {
		return fmt.Errorf("%v: %v", jerr.Error(), sig)
	}
	dataBytes, er := hex.DecodeString(strings.TrimPrefix(data, "0x"))
	if er == nil {
		return p.decodeErrData(jerr.Error(), dataBytes)
	}
	return logex.NewErrorf("%v: %v", jerr.Error(), data)

}

func (p *DcapPortal) attestationFee(opts *bind.TransactOpts, feeBase *big.Int) *big.Int {
	if p.chain.EIP1559 {
		return new(big.Int).Mul(opts.GasFeeCap, feeBase)
	} else {
		return new(big.Int).Mul(opts.GasPrice, feeBase)
	}
}

func (p *DcapPortal) normalizeOpts(optsRef *bind.TransactOpts) (*bind.TransactOpts, error) {
	var opts bind.TransactOpts
	if optsRef != nil {
		opts = *optsRef
	}
	if opts.Context == nil {
		opts.Context = context.Background()
	}
	if opts.Value != nil {
		return nil, ErrValueShouldBeNil.Trace()
	}

	var head *types.Header
	var err error

	if p.chain.EIP1559 && opts.GasFeeCap == nil {
		head, err = p.client.HeaderByNumber(opts.Context, nil)
		if err != nil {
			return nil, logex.Trace(err)
		}
		if head.BaseFee == nil && p.chain.EIP1559 {
			p.chain.EIP1559 = false
		}
	}

	if p.chain.EIP1559 {
		if opts.GasTipCap == nil {
			tip, err := p.client.SuggestGasTipCap(opts.Context)
			if err != nil {
				return nil, logex.Trace(err)
			}
			opts.GasTipCap = tip
		}
		if opts.GasFeeCap == nil {
			opts.GasFeeCap = new(big.Int).Add(
				opts.GasTipCap,
				new(big.Int).Mul(head.BaseFee, big.NewInt(basefeeMultiplier)),
			)
		}
	} else {
		if opts.GasPrice == nil {
			price, err := p.client.SuggestGasPrice(opts.Context)
			if err != nil {
				return nil, logex.Trace(err)
			}
			opts.GasPrice = price
		}
	}

	return &opts, nil
}

type JsonError interface {
	Error() string
	ErrorCode() int
	ErrorData() interface{}
}
