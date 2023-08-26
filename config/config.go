package config

import (
	"time"
)

// PerpsvConfig is a type for lib configuration schema
type PerpsvConfig struct {
	RPC                 string
	ContractAddresses   *ContractAddresses
	FirstContractBlocks *FirstContractBlocks
	ConnectionTimeout   time.Duration
	ReadTimeout         time.Duration
}

// ContractAddresses is a part of a PerpsvConfig struct with contract addresses
type ContractAddresses struct {
	Core        string
	SpotMarket  string
	PerpsMarket string
}

// FirstContractBlocks is a part of a config struct with default first block numbers used to filters contract logs
type FirstContractBlocks struct {
	Core        uint64
	SpotMarket  uint64
	PerpsMarket uint64
}

// GetGoerliDefaultPerpsvConfig is used to get default lib config for goerli optimism test net
func GetGoerliDefaultPerpsvConfig() *PerpsvConfig {
	return &PerpsvConfig{
		RPC: "https://rpc.goerli.optimism.gateway.fm",
		ContractAddresses: &ContractAddresses{
			Core:        "0x76490713314fCEC173f44e99346F54c6e92a8E42",
			SpotMarket:  "0x5FF4b3aacdeC86782d8c757FAa638d8790799E83",
			PerpsMarket: "0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b",
		},
		FirstContractBlocks: &FirstContractBlocks{
			Core:        11664658,
			SpotMarket:  10875051,
			PerpsMarket: 12708889,
		},
		ConnectionTimeout: time.Second * 30,
		ReadTimeout:       time.Second * 15,
	}
}

// GetOptimismDefaultPerpsvConfig is used to get default lib config for ompimism main net
func GetOptimismDefaultPerpsvConfig() *PerpsvConfig {
	return &PerpsvConfig{
		RPC: "https://rpc.optimism.gateway.fm",
		ContractAddresses: &ContractAddresses{
			Core:        "0xffffffaEff0B96Ea8e4f94b2253f31abdD875847",
			SpotMarket:  "0x38908Ee087D7db73A1Bd1ecab9AAb8E8c9C74595",
			PerpsMarket: "",
		},
		FirstContractBlocks: &FirstContractBlocks{
			Core:        94847041,
			SpotMarket:  94846457,
			PerpsMarket: 0,
		},
		ConnectionTimeout: time.Second * 30,
		ReadTimeout:       time.Second * 15,
	}
}
