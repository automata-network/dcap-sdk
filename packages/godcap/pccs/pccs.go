package pccs

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"

	"github.com/automata-network/dcap-sdk/packages/godcap/stubs/AutomataEnclaveIdentityDao"
	"github.com/automata-network/dcap-sdk/packages/godcap/stubs/AutomataFmspcTcbDao"
	"github.com/automata-network/dcap-sdk/packages/godcap/stubs/AutomataPcsDao"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ChainConfig holds the addresses of the smart contracts
type ChainConfig struct {
	AutomataPcsDao             common.Address `json:"automata_pcs_dao"`
	AutomataFmspcTcbDao        common.Address `json:"automata_fmspc_tcb_dao"`
	AutomataEnclaveIdentityDao common.Address `json:"automata_enclave_identity_dao"`
}

// Constants for CA types
const (
	CA_ROOT uint8 = iota
	CA_PROCESSOR
	CA_PLATFORM
	CA_SIGNING
)

// Constants for Enclave ID types
const (
	ENCLAVE_ID_QE uint8 = iota
	ENCLAVE_ID_QVE
	ENCLAVE_ID_TDQE
)

// Server struct holds the Ethereum client and contract instances
type Client struct {
	client    *ethclient.Client
	pcs       *AutomataPcsDao.AutomataPcsDao
	fmspc     *AutomataFmspcTcbDao.AutomataFmspcTcbDao
	enclaveId *AutomataEnclaveIdentityDao.AutomataEnclaveIdentityDao
}

// NewClient initializes a new Server instance
func NewClient(client *ethclient.Client, chain *ChainConfig) (*Client, error) {
	// Initialize AutomataPcsDao contract
	pcs, err := AutomataPcsDao.NewAutomataPcsDao(chain.AutomataPcsDao, client)
	if err != nil {
		return nil, logex.Trace(err, chain.AutomataPcsDao)
	}
	// Initialize AutomataFmspcTcbDao contract
	fmspc, err := AutomataFmspcTcbDao.NewAutomataFmspcTcbDao(chain.AutomataFmspcTcbDao, client)
	if err != nil {
		return nil, logex.Trace(err, chain.AutomataFmspcTcbDao)
	}
	// Initialize AutomataEnclaveIdentityDao contract
	enclaveId, err := AutomataEnclaveIdentityDao.NewAutomataEnclaveIdentityDao(chain.AutomataEnclaveIdentityDao, client)
	if err != nil {
		return nil, logex.Trace(err, chain.AutomataEnclaveIdentityDao)
	}
	return &Client{
		client:    client,
		pcs:       pcs,
		fmspc:     fmspc,
		enclaveId: enclaveId,
	}, nil
}

// CertCrl holds certificate and CRL data
type CertCrl struct {
	Cert []byte
	Crl  []byte
}

// GetCertByID retrieves a certificate by its CA ID
func (p *Client) GetCertByID(ctx context.Context, ca uint8) (*CertCrl, error) {
	result, err := p.pcs.GetCertificateById(nil, ca)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return (*CertCrl)(&result), nil
}

// TcbInfo holds TCB information and its signature
type TcbInfo struct {
	TcbInfo   json.RawMessage `json:"tcbInfo"`
	Signature string          `json:"signature"`
}

// Encode serializes TcbInfo to JSON
func (t *TcbInfo) Encode() []byte {
	data, _ := json.Marshal(t)
	return data
}

// GetTcbInfo retrieves TCB information by type, FMSPC, and version
func (p *Client) GetTcbInfo(ctx context.Context, tcbType uint8, fmspc string, tcbVersion uint32) (*TcbInfo, error) {
	result, err := p.fmspc.GetTcbInfo(&bind.CallOpts{Context: ctx}, big.NewInt(int64(tcbType)), fmspc, big.NewInt(int64(tcbVersion)))
	if err != nil {
		return nil, logex.Trace(err)
	}
	var info TcbInfo
	if err := json.Unmarshal([]byte(result.TcbInfoStr), &info.TcbInfo); err != nil {
		return nil, logex.Trace(err)
	}
	info.Signature = hex.EncodeToString(result.Signature)
	return &info, nil
}

// EnclaveIdentityInfo holds enclave identity information and its signature
type EnclaveIdentityInfo struct {
	Identity  json.RawMessage `json:"enclaveIdentity"`
	Signature string          `json:"signature"`
}

// Encode serializes EnclaveIdentityInfo to JSON
func (e *EnclaveIdentityInfo) Encode() []byte {
	data, _ := json.Marshal(e)
	return data
}

// GetEnclaveID retrieves enclave identity information by ID and version
func (p *Client) GetEnclaveID(ctx context.Context, enclaveId uint8, version uint32) (*EnclaveIdentityInfo, error) {
	result, err := p.enclaveId.GetEnclaveIdentity(&bind.CallOpts{Context: ctx}, big.NewInt(int64(enclaveId)), big.NewInt(int64(version)))
	if err != nil {
		return nil, logex.Trace(err)
	}
	var info EnclaveIdentityInfo
	if err := json.Unmarshal([]byte(result.IdentityStr), &info.Identity); err != nil {
		return nil, logex.Trace(err)
	}
	info.Signature = hex.EncodeToString(result.Signature)
	return &info, nil
}
