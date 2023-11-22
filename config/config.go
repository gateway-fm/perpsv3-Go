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
	ConnectionTimeout   time.Duration
	ReadTimeout         time.Duration
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

// GetBaseAndromedaDefaultConfig is used to get default lib config for base andromeda test net
func GetBaseAndromedaDefaultConfig(rpcURL string) *PerpsvConfig {
	if rpcURL == "" {
		rpcURL = "https://gateway.tenderly.co/public/base"
	}

	return &PerpsvConfig{
		ChainID: BaseAndromeda,
		RPC:     rpcURL,
		ContractAddresses: &ContractAddresses{
			Core:        "0xd7cFdcf6499896a1E15Db92629AAC8085af8B2ac",
			PerpsMarket: "0x9863Dae3f4b5F4Ffe3A841a21565d57F2BA10E87",
			Forwarder:   "0xAE788aaf52780741E12BF79Ad684B91Bb0EF4D92",
			ERC7412:     "0x9849832a1d8274aaeDb1112ad9686413461e7101",
		},
		FirstContractBlocks: &FirstContractBlocks{
			Core:        11920413,
			PerpsMarket: 9243743,
		},
		ConnectionTimeout: time.Second * 30,
		ReadTimeout:       time.Second * 15,
	}
}
