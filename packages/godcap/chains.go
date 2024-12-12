package godcap

import (
	"github.com/automata-network/dcap-sdk/packages/godcap/pccs"
	"github.com/ethereum/go-ethereum/common"
)

type ChainConfig struct {
	ChainId                    int64
	Name                       string
	Testnet                    bool
	OneRpc                     string
	Endpoint                   string
	EIP1559                    bool
	DcapPortal                 common.Address
	AutomataDcapAttestationFee common.Address
	PCCSRouter                 common.Address
	V3QuoteVerifier            common.Address
	V4QuoteVerifier            common.Address
	PCCS                       *pccs.ChainConfig
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

var (
	ChainAutomataMainnet = &ChainConfig{
		Name:                       "Automata Mainnet",
		ChainId:                    65536,
		Testnet:                    false,
		OneRpc:                     "https://1rpc.io/ata",
		Endpoint:                   "https://rpc.ata.network",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xaEd8bF5907fC8690b1cb70DFD459Ca5Ed1529246"),
		PCCSRouter:                 common.HexToAddress("0x722525B96b62e182F8A095af0a79d4EA2037795C"),
		V3QuoteVerifier:            common.HexToAddress("0x4613038C93aF8963dc9E5e46c9fb3cbc68724df1"),
		V4QuoteVerifier:            common.HexToAddress("0xdE13b52a02Bd0a48AcF4FCaefccb094b41135Ee2"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x9c54C72867b07caF2e6255CE32983c28aFE40F26"),
			AutomataPcsDao:             common.HexToAddress("0xcf171ACd6c0a776f9d3E1F6Cac8067c982Ac6Ce1"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x45f91C0d9Cf651785d93fcF7e9E97dE952CdB910"),
		},
	}

	ChainEthereumMainnet = &ChainConfig{
		Name:                       "Ethereum Mainnet",
		ChainId:                    1,
		Testnet:                    false,
		OneRpc:                     "https://1rpc.io/eth",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xE26E11B257856B0bEBc4C759aaBDdea72B64351F"),
		PCCSRouter:                 common.HexToAddress("0x09bBC921be046726bb5b694A49888e4e2e7AA9C3"),
		V3QuoteVerifier:            common.HexToAddress("0xF38a49322cAA0Ead71D4B1cF2afBb6d02BE5FC96"),
		V4QuoteVerifier:            common.HexToAddress("0xC86EE37Ee5030B9fF737F3E71f7611Abf5dfD9B7"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x868c18869f68E0E0b0b7B2B4439f7fDDd0421e6b"),
			AutomataPcsDao:             common.HexToAddress("0x86f8865BCe8BE62CB8096b5B94fA3fB3a6ED330c"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x28111536292b34f37120861A46B39BF39187d73a"),
		},
	}
	ChainBaseMainnet = &ChainConfig{
		Name:                       "Base Mainnet",
		ChainId:                    8453,
		Testnet:                    false,
		OneRpc:                     "https://1rpc.io/base",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xaEd8bF5907fC8690b1cb70DFD459Ca5Ed1529246"),
		PCCSRouter:                 common.HexToAddress("0x722525B96b62e182F8A095af0a79d4EA2037795C"),
		V3QuoteVerifier:            common.HexToAddress("0x4613038C93aF8963dc9E5e46c9fb3cbc68724df1"),
		V4QuoteVerifier:            common.HexToAddress("0xdE13b52a02Bd0a48AcF4FCaefccb094b41135Ee2"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x9c54C72867b07caF2e6255CE32983c28aFE40F26"),
			AutomataPcsDao:             common.HexToAddress("0xcf171ACd6c0a776f9d3E1F6Cac8067c982Ac6Ce1"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x45f91C0d9Cf651785d93fcF7e9E97dE952CdB910"),
		},
	}
	ChainOPMainnet = &ChainConfig{
		Name:                       "OP Mainnet",
		ChainId:                    10,
		Testnet:                    false,
		OneRpc:                     "https://1rpc.io/op",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xaEd8bF5907fC8690b1cb70DFD459Ca5Ed1529246"),
		PCCSRouter:                 common.HexToAddress("0x722525B96b62e182F8A095af0a79d4EA2037795C"),
		V3QuoteVerifier:            common.HexToAddress("0x4613038C93aF8963dc9E5e46c9fb3cbc68724df1"),
		V4QuoteVerifier:            common.HexToAddress("0xdE13b52a02Bd0a48AcF4FCaefccb094b41135Ee2"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x9c54C72867b07caF2e6255CE32983c28aFE40F26"),
			AutomataPcsDao:             common.HexToAddress("0xcf171ACd6c0a776f9d3E1F6Cac8067c982Ac6Ce1"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x45f91C0d9Cf651785d93fcF7e9E97dE952CdB910"),
		},
	}
	ChainWorldMainnet = &ChainConfig{
		Name:                       "World Mainnet",
		ChainId:                    480,
		Testnet:                    false,
		OneRpc:                     "",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xE26E11B257856B0bEBc4C759aaBDdea72B64351F"),
		PCCSRouter:                 common.HexToAddress("0x09bBC921be046726bb5b694A49888e4e2e7AA9C3"),
		V3QuoteVerifier:            common.HexToAddress("0xF38a49322cAA0Ead71D4B1cF2afBb6d02BE5FC96"),
		V4QuoteVerifier:            common.HexToAddress("0xC86EE37Ee5030B9fF737F3E71f7611Abf5dfD9B7"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x868c18869f68E0E0b0b7B2B4439f7fDDd0421e6b"),
			AutomataPcsDao:             common.HexToAddress("0x86f8865BCe8BE62CB8096b5B94fA3fB3a6ED330c"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x28111536292b34f37120861A46B39BF39187d73a"),
		},
	}
	ChainArbitrumMainnet = &ChainConfig{
		Name:                       "Arbitrum Mainnet",
		ChainId:                    42161,
		Testnet:                    false,
		OneRpc:                     "https://1rpc.io/arb",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xaEd8bF5907fC8690b1cb70DFD459Ca5Ed1529246"),
		PCCSRouter:                 common.HexToAddress("0x722525B96b62e182F8A095af0a79d4EA2037795C"),
		V3QuoteVerifier:            common.HexToAddress("0x4613038C93aF8963dc9E5e46c9fb3cbc68724df1"),
		V4QuoteVerifier:            common.HexToAddress("0xdE13b52a02Bd0a48AcF4FCaefccb094b41135Ee2"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x9c54C72867b07caF2e6255CE32983c28aFE40F26"),
			AutomataPcsDao:             common.HexToAddress("0xcf171ACd6c0a776f9d3E1F6Cac8067c982Ac6Ce1"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x45f91C0d9Cf651785d93fcF7e9E97dE952CdB910"),
		},
	}
	ChainAutomataTestnet = &ChainConfig{
		Name:     "Automata Testnet",
		ChainId:  1398243,
		Testnet:  true,
		OneRpc:   "https://1rpc.io/ata/testnet",
		Endpoint: "https://rpc-testnet.ata.network",
		EIP1559:  true,
		// DcapPortal:                 common.HexToAddress("0xcbb758c7399cBa70Ca8B8f00D32733bC5cc89c48"), // original
		DcapPortal: common.HexToAddress("0xA1878A7Bc4B277e6802258339393dcf97ad2a044"),

		AutomataDcapAttestationFee: common.HexToAddress("0x6D67Ae70d99A4CcE500De44628BCB4DaCfc1A145"),
		PCCSRouter:                 common.HexToAddress("0x3095741175094128ae9F451fa3693B2d23719940"),
		V3QuoteVerifier:            common.HexToAddress("0x6cc70fDaB6248b374A7fD4930460F7b017190872"),
		V4QuoteVerifier:            common.HexToAddress("0x015E89a5fF935Fbc361DcB4Bac71e5cD8a5CeEe3"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x9c54C72867b07caF2e6255CE32983c28aFE40F26"),
			AutomataPcsDao:             common.HexToAddress("0xcf171ACd6c0a776f9d3E1F6Cac8067c982Ac6Ce1"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x45f91C0d9Cf651785d93fcF7e9E97dE952CdB910"),
		},
	}
	ChainEthereumSepolia = &ChainConfig{
		Name:                       "Ethereum Sepolia",
		ChainId:                    11155111,
		Testnet:                    true,
		OneRpc:                     "https://1rpc.io/sepolia",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xE28ea4E574871CA6A4331d6692bd3DD602Fb4f76"),
		PCCSRouter:                 common.HexToAddress("0xfFC62c8851F54723206235E24af1bf10b9ea1d47"),
		V3QuoteVerifier:            common.HexToAddress("0x6E64769A13617f528a2135692484B681Ee1a7169"),
		V4QuoteVerifier:            common.HexToAddress("0x90c14Bd25744d8b1E3971951BD56BfFf24dC053A"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0xB87a493684Bb643258Ae4887B444c6cB244db935"),
			AutomataPcsDao:             common.HexToAddress("0x980AEAdb3fa7c2c58A81091D93A819a24A103E6C"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x5eFDd14Bbfba36992f66a64653962BB0B8Ef1E26"),
		},
	}
	ChainEthereumHolesky = &ChainConfig{
		Name:                       "Ethereum Holesky",
		ChainId:                    17000,
		Testnet:                    true,
		OneRpc:                     "https://1rpc.io/holesky",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xaEd8bF5907fC8690b1cb70DFD459Ca5Ed1529246"),
		PCCSRouter:                 common.HexToAddress("0x729E3e7542E8A6630818E9a14A67e0Cb7008a5E5"),
		V3QuoteVerifier:            common.HexToAddress("0x4613038C93aF8963dc9E5e46c9fb3cbc68724df1"),
		V4QuoteVerifier:            common.HexToAddress("0xdE13b52a02Bd0a48AcF4FCaefccb094b41135Ee2"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x9c54C72867b07caF2e6255CE32983c28aFE40F26"),
			AutomataPcsDao:             common.HexToAddress("0xcf171ACd6c0a776f9d3E1F6Cac8067c982Ac6Ce1"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x45f91C0d9Cf651785d93fcF7e9E97dE952CdB910"),
		},
	}
	ChainBaseSepolia = &ChainConfig{
		Name:                       "Base Sepolia",
		ChainId:                    84532,
		Testnet:                    true,
		OneRpc:                     "",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xaEd8bF5907fC8690b1cb70DFD459Ca5Ed1529246"),
		PCCSRouter:                 common.HexToAddress("0x729E3e7542E8A6630818E9a14A67e0Cb7008a5E5"),
		V3QuoteVerifier:            common.HexToAddress("0x4613038C93aF8963dc9E5e46c9fb3cbc68724df1"),
		V4QuoteVerifier:            common.HexToAddress("0xdE13b52a02Bd0a48AcF4FCaefccb094b41135Ee2"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x9c54C72867b07caF2e6255CE32983c28aFE40F26"),
			AutomataPcsDao:             common.HexToAddress("0xcf171ACd6c0a776f9d3E1F6Cac8067c982Ac6Ce1"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x45f91C0d9Cf651785d93fcF7e9E97dE952CdB910"),
		},
	}
	ChainOPSepolia = &ChainConfig{
		Name:                       "OP Sepolia",
		ChainId:                    11155420,
		Testnet:                    true,
		OneRpc:                     "",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xaEd8bF5907fC8690b1cb70DFD459Ca5Ed1529246"),
		PCCSRouter:                 common.HexToAddress("0x729E3e7542E8A6630818E9a14A67e0Cb7008a5E5"),
		V3QuoteVerifier:            common.HexToAddress("0x4613038C93aF8963dc9E5e46c9fb3cbc68724df1"),
		V4QuoteVerifier:            common.HexToAddress("0xdE13b52a02Bd0a48AcF4FCaefccb094b41135Ee2"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x9c54C72867b07caF2e6255CE32983c28aFE40F26"),
			AutomataPcsDao:             common.HexToAddress("0xcf171ACd6c0a776f9d3E1F6Cac8067c982Ac6Ce1"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x45f91C0d9Cf651785d93fcF7e9E97dE952CdB910"),
		},
	}
	ChainWorldSepolia = &ChainConfig{
		Name:                       "World Sepolia",
		ChainId:                    4801,
		Testnet:                    true,
		OneRpc:                     "",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xaEd8bF5907fC8690b1cb70DFD459Ca5Ed1529246"),
		PCCSRouter:                 common.HexToAddress("0x729E3e7542E8A6630818E9a14A67e0Cb7008a5E5"),
		V3QuoteVerifier:            common.HexToAddress("0x4613038C93aF8963dc9E5e46c9fb3cbc68724df1"),
		V4QuoteVerifier:            common.HexToAddress("0xdE13b52a02Bd0a48AcF4FCaefccb094b41135Ee2"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x9c54C72867b07caF2e6255CE32983c28aFE40F26"),
			AutomataPcsDao:             common.HexToAddress("0xcf171ACd6c0a776f9d3E1F6Cac8067c982Ac6Ce1"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x45f91C0d9Cf651785d93fcF7e9E97dE952CdB910"),
		},
	}
	ChainArbitrumSepolia = &ChainConfig{
		Name:                       "Arbitrum Sepolia",
		ChainId:                    421614,
		Testnet:                    true,
		OneRpc:                     "",
		EIP1559:                    true,
		AutomataDcapAttestationFee: common.HexToAddress("0xaEd8bF5907fC8690b1cb70DFD459Ca5Ed1529246"),
		PCCSRouter:                 common.HexToAddress("0x729E3e7542E8A6630818E9a14A67e0Cb7008a5E5"),
		V3QuoteVerifier:            common.HexToAddress("0x4613038C93aF8963dc9E5e46c9fb3cbc68724df1"),
		V4QuoteVerifier:            common.HexToAddress("0xdE13b52a02Bd0a48AcF4FCaefccb094b41135Ee2"),
		PCCS: &pccs.ChainConfig{
			AutomataFmspcTcbDao:        common.HexToAddress("0x9c54C72867b07caF2e6255CE32983c28aFE40F26"),
			AutomataPcsDao:             common.HexToAddress("0xcf171ACd6c0a776f9d3E1F6Cac8067c982Ac6Ce1"),
			AutomataEnclaveIdentityDao: common.HexToAddress("0x45f91C0d9Cf651785d93fcF7e9E97dE952CdB910"),
		},
	}
)
