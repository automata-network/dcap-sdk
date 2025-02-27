package godcap

import (
	"context"
	"encoding/hex"
	"math/big"
	"os"
	"testing"

	"github.com/automata-network/dcap-sdk/packages/godcap/mock"
	"github.com/automata-network/dcap-sdk/packages/godcap/stubs/VerifiedCounter"
	"github.com/automata-network/dcap-sdk/packages/godcap/zkdcap"
	"github.com/chzyer/logex"
	"github.com/chzyer/test"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var verifiedCounterAddr = common.HexToAddress("0xc2FfB783e36c5F4718B96D527d8983222FAF4680")

func TestDcapPortalOnChain(t *testing.T) {
	defer test.New(t)

	ctx := context.Background()
	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		logex.Info("skip testing DcapPortal because env var PRIVATE_KEY is empty")
		return
	}

	portal, err := NewDcapPortal(ctx, WithPrivateKey(privateKey), WithEndpoint(ChainAutomataTestnet.Endpoint))
	test.Nil(err)

	counter, err := VerifiedCounter.NewVerifiedCounterCaller(verifiedCounterAddr, portal.Client())
	test.Nil(err)

	originCounter, err := counter.Number(nil)
	test.Nil(err)

	// deposit 10 wei to increase the counter, check the logic from ../dcap-portal/src/examples/VerifiedCounter.sol
	callback := NewCallbackFromAbiJSON(VerifiedCounter.VerifiedCounterABI).
		WithParams("deposit").
		WithTo(verifiedCounterAddr).
		WithValue(big.NewInt(10))

	tx, err := portal.VerifyAndAttestOnChain(nil, mock.Quotes[0], callback)
	test.Nil(err)

	receipt := <-portal.WaitTx(ctx, tx)

	portal.PrintAttestationFee(tx, callback, receipt)

	newCounter, err := counter.Number(&bind.CallOpts{BlockNumber: new(big.Int).Set(receipt.BlockNumber)})
	if err != nil {
		t.Fatal(err)
	}

	if new(big.Int).Sub(newCounter, originCounter).Cmp(big.NewInt(10)) != 0 {
		t.Fatalf("counter mismatch: origin=%v, new=%v", originCounter, newCounter)
	}
}

func TestSp1(t *testing.T) {
	privateKey := os.Getenv("NETWORK_PRIVATE_KEY")
	if privateKey == "" {
		logex.Info("skip testing sp1 because env var NETWORK_PRIVATE_KEY is empty")
		return
	}

	ctx := context.Background()
	portal, err := NewDcapPortal(ctx, WithEndpoint(ChainAutomataTestnet.Endpoint), WithZkProof(nil))
	if err != nil {
		t.Fatal(err)
	}
	zkproof, err := portal.GenerateZkProof(ctx, zkdcap.ZkTypeSuccinct, mock.Quotes[1])
	if err != nil {
		logex.Error(err)
		t.Fatal(err)
	}
	succ, err := portal.CheckZkProof(ctx, zkproof)
	if err != nil {
		logex.Error(err)
		t.Fatal(err)
	}
	if !succ {
		t.Fatal("verify zkproof failed")
	}
}

func TestRisc0(t *testing.T) {
	privateKey := os.Getenv("BONSAI_API_KEY")
	if privateKey == "" {
		logex.Info("skip testing risc0 because env var BONSAI_API_KEY is empty")
		return
	}

	ctx := context.Background()
	portal, err := NewDcapPortal(ctx, WithEndpoint(ChainAutomataTestnet.Endpoint), WithZkProof(nil))
	if err != nil {
		t.Fatal(err)
	}
	zkproof, err := portal.GenerateZkProof(ctx, zkdcap.ZkTypeRiscZero, mock.Quotes[1])
	if err != nil {
		logex.Error(err)
		t.Fatal(err)
	}
	succ, err := portal.CheckZkProof(ctx, zkproof)
	if err != nil {
		logex.Error(err)
		t.Fatal(err)
	}
	if !succ {
		t.Fatal("verify zkproof failed")
	}
}

// we use a mock attestation contract to test the attestation fee
func TestDcapPortalWithFee(t *testing.T) {
	defer test.New(t)
	ctx := context.Background()

	privateKey := os.Getenv("PRIVATE_KEY")
	if privateKey == "" {
		logex.Info("skip testing DcapPortal because env var PRIVATE_KEY is empty")
		return
	}

	chain := *ChainAutomataTestnet
	chain.AutomataDcapAttestationFee = common.HexToAddress("0xA0c3a7C811e3B6b7D7a381b3aD29A7FCF9048DFf")
	chain.DcapPortal = common.HexToAddress("0x1aFedD4123494f83ADc166A4Fd6Da96321c88c41")

	mockVerifiedCounterAddr := common.HexToAddress("0x5BE14673A6d40C711F082D6f7e4796E2fC57d7b2")
	callback := NewCallbackFromAbiJSON(VerifiedCounter.VerifiedCounterABI).
		WithParams("deposit").
		WithTo(mockVerifiedCounterAddr).
		WithValue(big.NewInt(10))

	portal, err := NewDcapPortal(ctx, WithChainConfig(&chain), WithPrivateKey(privateKey))
	test.Nil(err)

	succ, err := portal.CheckQuote(ctx, mock.Quotes[0])
	test.Nil(err)
	test.True(succ)

	opt, err := portal.BuildTransactOpts(ctx)
	test.Nil(err)
	opt.NoSend = true

	_, err = portal.VerifyAndAttestOnChain(opt, mock.Quotes[0], callback)
	test.Nil(err)
}

func TestDcapPortalZkProof(t *testing.T) {
	defer test.New(t)
	ctx := context.Background()
	output, _ := hex.DecodeString("02550004000000810790c06f000000040102000000000000000000000000009790d89a10210ec6968a773cee2ca05b5aa97309f36727a968527be4606fc19e6f73acce350946c9d46a9bf7a63f843000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000080e702060000000000f2dd2696f69b950645832bdc095ffd11247eeff687eeacdb57a58d2ddb9a9f94fea40c961e19460c00ffa31420ecbc180000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000998204508d58dcbfebe5e11c48669f7a921ac2da744dfb7d014ecdff2acdff1c9f665fdad52aadacf296a1df9909eb2383d100224f1716aeb431f7cb3cf028197dbd872487f27b0f6329ab17647dc9953c7014109818634f879e6550bc60f93eecfc42ff4d49278bfdbb0c77e570f4490cff10a2ee1ac11fbd2c2b49fa6cfa3cf1a1cb755c72522dd8a689e9d47906a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000278e753482976c8a7351fe2113609c7350d491cdae3d449eefc202fa41b2ad6840239cc2ba084c2d594b4e6dabeae0fcbf71c96daf0d0c9ecf0e9810c04579000000000067c002df31109a57d5e3d5326395ad84591ca85f6ec68885ed1973188829b33378f1435cf544a2a578ae47b0407e9088073d56882bd7b5fe1e979033773ad177500d5bd80fa74a3f32c80b978c8ad671395dabf24283eef9091bc3919fd39b9915a87f1adf3061c165c0191e2658256a2855cac9267f179aafb1990c9e918d6452816adf9953f245d005b9d7d8e36a842a60b51e5cf85b2c2072ae397c178535c9985b77f360068c37da390ae1b67de639f755b83b1eb8bad6a5e8467cc1d4d42a521d32")
	seal, _ := hex.DecodeString("c101b42b0ce05176f0df453d3a5abb909ff6ba39700b3b85de9d087354196675f1a2a2b90653f9cdb08b13e23a44552a124da087ea1c57a9841c4882cc6dd1df0ecb8fa025d947b0c0a12576df4e9db3f2481772a8ff12322ea72769e42bd3d4894bf09203ed599a9e213bb2355e7c3405c86cf765610b68e8a94ecfb8cf256c8bd7b8281f898dbd64b44c98b9ccfdff664ab79f7ba62b02d67b05e01949fba6624937671b01076781fe07b169ae9748f59a2f3e3eabcb8089f197956e34ee205f37ea0011607cb34f566fa919ce19a07476aa8610ade05127a7f7b09d1db0429b76fd0c266ca86af3576105ba0261e419213804e817c76bc84a11fcee99d5156818688d")
	zkproof := &zkdcap.ZkProof{
		Type:   zkdcap.ZkTypeRiscZero,
		Output: output,
		Proof:  seal,
	}
	portal, err := NewDcapPortal(ctx, WithChainConfig(ChainAutomataTestnet))
	test.Nil(err)

	succ, err := portal.CheckZkProof(ctx, zkproof)
	test.Nil(err)
	if !succ {
		t.Fatal("verify zkproof failed")
	}
}
