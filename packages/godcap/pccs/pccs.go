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

type ChainConfig struct {
	AutomataPcsDao             common.Address
	AutomataFmspcTcbDao        common.Address
	AutomataEnclaveIdentityDao common.Address
}

const (
	CA_ROOT uint8 = iota
	CA_PROCESSOR
	CA_PLATFORM
	CA_SIGNING
)

const (
	ENCLAVE_ID_QE uint8 = iota
	ENCLAVE_ID_QVE
	ENCLAVE_ID_TDQE
)

type Server struct {
	client    *ethclient.Client
	pcs       *AutomataPcsDao.AutomataPcsDao
	fmspc     *AutomataFmspcTcbDao.AutomataFmspcTcbDao
	enclaveId *AutomataEnclaveIdentityDao.AutomataEnclaveIdentityDao
}

func NewServer(client *ethclient.Client, chain *ChainConfig) (*Server, error) {
	pcs, err := AutomataPcsDao.NewAutomataPcsDao(chain.AutomataPcsDao, client)
	if err != nil {
		return nil, logex.Trace(err, chain.AutomataPcsDao)
	}
	fmspc, err := AutomataFmspcTcbDao.NewAutomataFmspcTcbDao(chain.AutomataFmspcTcbDao, client)
	if err != nil {
		return nil, logex.Trace(err, chain.AutomataFmspcTcbDao)
	}
	enclaveId, err := AutomataEnclaveIdentityDao.NewAutomataEnclaveIdentityDao(chain.AutomataEnclaveIdentityDao, client)
	if err != nil {
		return nil, logex.Trace(err, chain.AutomataEnclaveIdentityDao)
	}
	return &Server{
		client:    client,
		pcs:       pcs,
		fmspc:     fmspc,
		enclaveId: enclaveId,
	}, nil
}

type CertCrl struct {
	Cert []byte
	Crl  []byte
}

func (p *Server) GetCertByID(ctx context.Context, ca uint8) (*CertCrl, error) {
	result, err := p.pcs.GetCertificateById(nil, ca)
	if err != nil {
		return nil, logex.Trace(err)
	}
	return (*CertCrl)(&result), nil
}

type TcbInfo struct {
	TcbInfo   json.RawMessage `json:"tcbInfo"`
	Signature string          `json:"signature"`
}

func (t *TcbInfo) Encode() []byte {
	data, _ := json.Marshal(t)
	return data
}

func (p *Server) GetTcbInfo(ctx context.Context, tcbType uint8, fmspc string, tcbVersion uint32) (*TcbInfo, error) {
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

type EnclaveIdentityInfo struct {
	Identity  json.RawMessage `json:"enclaveIdentity"`
	Signature string          `json:"signature"`
}

func (e *EnclaveIdentityInfo) Encode() []byte {
	data, _ := json.Marshal(e)
	return data
}

func (p *Server) GetEnclaveID(ctx context.Context, enclaveId uint8, version uint32) (*EnclaveIdentityInfo, error) {
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
