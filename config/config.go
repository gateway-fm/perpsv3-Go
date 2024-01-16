package config

import (
	"time"
)

// PerpsvConfig is a type for lib configuration schema
type PerpsvConfig struct {
	ChainID             ChainID
	RPC                 string
	ContractAddresses   *ContractAddresses
	FirstContractBlocks *FirstContractBlocks
	Multicall           *Multicall
	ConnectionTimeout   time.Duration
	ReadTimeout         time.Duration
}

type Multicall struct {
	Retries int
	Wait    time.Duration
}

// ContractAddresses is a part of a PerpsvConfig struct with contract addresses
type ContractAddresses struct {
	Core        string
	PerpsMarket string
	ERC7412     string
	Forwarder   string
}

// FirstContractBlocks is a part of a config struct with default first block numbers used to filters contract logs
type FirstContractBlocks struct {
	Core        uint64
	PerpsMarket uint64
}

// GetOptimismGoerliDefaultConfig is used to get default lib config for goerli optimism test net
func GetOptimismGoerliDefaultConfig(rpcURL string) *PerpsvConfig {
	if rpcURL == "" {
		rpcURL = "https://rpc.goerli.optimism.gateway.fm"
	}

	return &PerpsvConfig{
		ChainID: OptimismGoerli,
		RPC:     rpcURL,
		ContractAddresses: &ContractAddresses{
			Core:        "0x76490713314fCEC173f44e99346F54c6e92a8E42",
			PerpsMarket: "0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b",
		},
		FirstContractBlocks: &FirstContractBlocks{
			Core:        11664658,
			PerpsMarket: 12708889,
		},
		ConnectionTimeout: time.Second * 30,
		ReadTimeout:       time.Second * 15,
	}
}

// GetBaseSepoliaDefaultConfig is used to get default lib config for base sepolia test net
func GetBaseSepoliaDefaultConfig(rpcURL string) *PerpsvConfig {
	if rpcURL == "" {
		rpcURL = "https://base-sepolia.blockpi.network/v1/rpc/public"
	}

	return &PerpsvConfig{
		ChainID: BaseAndromeda,
		RPC:     rpcURL,
		Multicall: &Multicall{
			Retries: 5,
			Wait:    time.Millisecond * 200,
		},
		ContractAddresses: &ContractAddresses{
			Core:        "0xF4Df9Dd327Fd30695d478c3c8a2fffAddcdD0d31",
			PerpsMarket: "0xE6C5f05C415126E6b81FCc3619f65Db2fCAd58D0",
			Forwarder:   "0xE2C5658cC5C448B48141168f3e475dF8f65A1e3e",
			ERC7412:     "0xBf01fE835b3315968bbc094f50AE3164e6d3D969",
		},
		FirstContractBlocks: &FirstContractBlocks{
			Core:        4548696,
			PerpsMarket: 4548969,
		},
		ConnectionTimeout: time.Second * 30,
		ReadTimeout:       time.Second * 15,
	}
}

// GetBaseAndromedaDefaultConfig is used to get default lib config for base goerli test net
func GetBaseAndromedaDefaultConfig(rpcURL string) *PerpsvConfig {
	if rpcURL == "" {
		rpcURL = "https://rpc.ankr.com/base_goerli/6259fa6541ffabb10ca241f7f437c2389ab7dda38c7be817ab0fb76992e73ae5"
	}

	return &PerpsvConfig{
		ChainID: BaseAndromeda,
		RPC:     rpcURL,
		Multicall: &Multicall{
			Retries: 5,
			Wait:    time.Millisecond * 200,
		},
		ContractAddresses: &ContractAddresses{
			Core:        "0xF4Df9Dd327Fd30695d478c3c8a2fffAddcdD0d31",
			PerpsMarket: "0x75c43165ea38cB857C45216a37C5405A7656673c",
			Forwarder:   "0xE2C5658cC5C448B48141168f3e475dF8f65A1e3e",
			ERC7412:     "0xEa7a8f0fDD16Ccd46BA541Fb657a0A7FD7E36261",
		},
		FirstContractBlocks: &FirstContractBlocks{
			Core:        13044276,
			PerpsMarket: 13044488,
		},
		ConnectionTimeout: time.Second * 30,
		ReadTimeout:       time.Second * 15,
	}
}

// GetBaseMainnetDefaultConfig is used to get default lib config for base main net
func GetBaseMainnetDefaultConfig(rpcURL string) *PerpsvConfig {
	if rpcURL == "" {
		rpcURL = "https://rpc.ankr.com/base/6259fa6541ffabb10ca241f7f437c2389ab7dda38c7be817ab0fb76992e73ae5"
	}

	return &PerpsvConfig{
		ChainID: BaseMainnet,
		RPC:     rpcURL,
		Multicall: &Multicall{
			Retries: 5,
			Wait:    time.Millisecond * 200,
		},
		ContractAddresses: &ContractAddresses{
			Core:        "0x32C222A9A159782aFD7529c87FA34b96CA72C696",
			PerpsMarket: "0x0A2AF931eFFd34b81ebcc57E3d3c9B1E1dE1C9Ce",
			Forwarder:   "0xE2C5658cC5C448B48141168f3e475dF8f65A1e3e",
			ERC7412:     "0xEb38e347F24ea04ffA945a475BdD949E0c383A0F",
		},
		FirstContractBlocks: &FirstContractBlocks{
			Core:        7889212,
			PerpsMarket: 7889389,
		},
		ConnectionTimeout: time.Second * 30,
		ReadTimeout:       time.Second * 15,
	}
}
