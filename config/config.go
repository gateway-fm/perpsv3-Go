package config

import (
	"time"
)

//const (
//	GoerliID   = 420
//	OptimismID = 10
//)

// PerpsvConfig is a type for lib configuration schema
type PerpsvConfig struct {
	RPC                       string
	CoreContractAddress       string
	CoreContractFirstBlock    uint64
	SpotMarketContractAddress string
	SpotMarketFirstBlock      uint64
	//PerpsMarketContractAddress string
	//ChainID                    int
	ConnectionTimeout time.Duration
	ReadTimeout       time.Duration
}

// GetGoerliDefaultPerpsvConfig is used to get default lib config for goerli optimism test net
func GetGoerliDefaultPerpsvConfig() *PerpsvConfig {
	return &PerpsvConfig{
		RPC:                       "https://rpc.goerli.optimism.gateway.fm",
		CoreContractAddress:       "0x76490713314fCEC173f44e99346F54c6e92a8E42",
		CoreContractFirstBlock:    11664658,
		SpotMarketContractAddress: "0x5FF4b3aacdeC86782d8c757FAa638d8790799E83",
		SpotMarketFirstBlock:      10875051,
		//PerpsMarketContractAddress: "0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b",
		//ChainID:                    GoerliID,
		ConnectionTimeout: time.Second * 30,
		ReadTimeout:       time.Second * 15,
	}
}

// GetOptimismDefaultPerpsvConfig is used to get default lib config for ompimism main net
func GetOptimismDefaultPerpsvConfig() *PerpsvConfig {
	return &PerpsvConfig{
		RPC:                       "https://rpc.optimism.gateway.fm",
		CoreContractAddress:       "0xffffffaEff0B96Ea8e4f94b2253f31abdD875847",
		CoreContractFirstBlock:    94847041,
		SpotMarketContractAddress: "0x38908Ee087D7db73A1Bd1ecab9AAb8E8c9C74595",
		SpotMarketFirstBlock:      94846457,
		//PerpsMarketContractAddress: "",
		//ChainID:                    OptimismID,
		ConnectionTimeout: time.Second * 30,
		ReadTimeout:       time.Second * 15,
	}
}
