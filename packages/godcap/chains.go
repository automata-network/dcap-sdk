package godcap

import (
	"encoding/json"

	_ "embed"

	"github.com/automata-network/dcap-sdk/packages/godcap/pccs"
	"github.com/chzyer/logex"
	"github.com/ethereum/go-ethereum/common"
)

type ChainConfig struct {
	ChainId                    int64             `json:"chain_id"`
	Name                       string            `json:"name"`
	Testnet                    bool              `json:"testnet"`
	OneRpc                     string            `json:"one_rpc"`
	Endpoint                   string            `json:"endpoint"`
	EIP1559                    bool              `json:"eip_1559"`
	Explorer                   string            `json:"explorer"`
	DcapPortal                 common.Address    `json:"dcap_portal"`
	AutomataDcapAttestationFee common.Address    `json:"automata_dcap_attestation_fee"`
	PCCSRouter                 common.Address    `json:"pccs_router"`
	V3QuoteVerifier            common.Address    `json:"v3_quote_verifier"`
	V4QuoteVerifier            common.Address    `json:"v4_quote_verifier"`
	PCCS                       *pccs.ChainConfig `json:"pccs"`
}

func parseChainConfig(data []byte) *ChainConfig {
	var chainConfig ChainConfig
	err := json.Unmarshal(data, &chainConfig)
	if err != nil {
		logex.Info(string(data))
		logex.Fatal(err)
	}
	return &chainConfig
}

func ChainConfigFromChainId(chainId int64) *ChainConfig {
	for _, item := range Chains {
		if item.ChainId == chainId {
			return item
		}
	}
	return nil
}

func AddChainConfig(chain *ChainConfig) bool {
	for _, item := range Chains {
		if item.ChainId == chain.ChainId {
			return false
		}
	}
	Chains = append(Chains, chain)
	return true
}

var Chains = []*ChainConfig{
	ChainAutomataMainnet,
	ChainEthereumMainnet,
	ChainBaseMainnet,
	ChainOPMainnet,
	ChainWorldMainnet,
	ChainArbitrumMainnet,
	ChainAutomataTestnet,
	ChainEthereumSepolia,
	ChainEthereumHolesky,
	ChainBaseSepolia,
	ChainOPSepolia,
	ChainWorldSepolia,
	ChainArbitrumSepolia,
}

//go:embed chains_config/automata_mainnet.json
var automataMainnet []byte
var ChainAutomataMainnet = parseChainConfig(automataMainnet)

//go:embed chains_config/ethereum_mainnet.json
var ethereumMainnet []byte
var ChainEthereumMainnet = parseChainConfig(ethereumMainnet)

//go:embed chains_config/base_mainnet.json
var baseMainnet []byte
var ChainBaseMainnet = parseChainConfig(baseMainnet)

//go:embed chains_config/op_mainnet.json
var opMainnet []byte
var ChainOPMainnet = parseChainConfig(opMainnet)

//go:embed chains_config/world_mainnet.json
var worldMainnet []byte
var ChainWorldMainnet = parseChainConfig(worldMainnet)

//go:embed chains_config/arbitrum_mainnet.json
var arbitrumMainnet []byte
var ChainArbitrumMainnet = parseChainConfig(arbitrumMainnet)

//go:embed chains_config/automata_testnet.json
var automataTestnet []byte
var ChainAutomataTestnet = parseChainConfig(automataTestnet)

//go:embed chains_config/ethereum_sepolia.json
var ethereumSepolia []byte
var ChainEthereumSepolia = parseChainConfig(ethereumSepolia)

//go:embed chains_config/ethereum_holesky.json
var ethereumHolesky []byte
var ChainEthereumHolesky = parseChainConfig(ethereumHolesky)

//go:embed chains_config/base_sepolia.json
var baseSepolia []byte
var ChainBaseSepolia = parseChainConfig(baseSepolia)

//go:embed chains_config/op_sepolia.json
var opSepolia []byte
var ChainOPSepolia = parseChainConfig(opSepolia)

//go:embed chains_config/world_sepolia.json
var worldSepolia []byte
var ChainWorldSepolia = parseChainConfig(worldSepolia)

//go:embed chains_config/arbitrum_sepolia.json
var arbitrumSepolia []byte
var ChainArbitrumSepolia = parseChainConfig(arbitrumSepolia)
