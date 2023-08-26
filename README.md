# perpsv3-Go

This repository contains a Go library for interacting with smart contracts related to a [Synthetix V3](https://docs.synthetix.io/v/v3/) 
DeFi protocol. It includes components to work with Core, Spot Market, and Perps Market contracts deployed on the Optimism 
mainnet and Optimistic Goerli testnet.

## Table of Contents

- [Getting started](#getting-started)
- [Configuration](#configuration)
- [API Reference](#api-reference)
- [License](#license)

## Getting started

To install a library use:

```
go get github.com/gateway-fm/perpsv3-Go
go mod tidy
```

A usage example:

```go
package main

import (
	"log"

	perpsv3_Go "github.com/gateway-fm/perpsv3-Go"
	"github.com/gateway-fm/perpsv3-Go/config"
)

func main() {
	perpsLib, err := perpsv3_Go.Create(config.GetGoerliDefaultPerpsvConfig())
	if err != nil {
		log.Fatal(err)
	}
	//...
}
```

## Configuration

You can use a default configurations. For now only two default configurations are available:

A configuration for Goerli Optimistic testnet:
```go
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
```

And a configuration for Optimism mainnet. Be informed that mainnet configs will return an error at this version due to 
the luck of Perps Market contract address on mainnnet.

```go
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
```

You can use you own configuration by creating a Config instance:

```go
package main

import (
	"log"
	perpsv3_Go "github.com/gateway-fm/perpsv3-Go"
	"github.com/gateway-fm/perpsv3-Go/config"
)

func main() {
	conf := &config.PerpsvConfig{
		RPC: "https://rpc.optimism.gateway.fm",
		//...
    }
	
	perpsLib, err := perpsv3_Go.Create(conf)
	if err != nil {
		log.Fatal(err)
	}
	//...
}
```

## API Reference

### Trades

#### Model

Using Trades services you operate with Trades model which represents a `OrderSettled` event of Perps Market smart-contract 
with some additional fields:

```go
type Trade struct {
	// Event fields:
	MarketID         uint64          // ID of the market used for the trade
	AccountID        uint64          // ID of the account used for the trade
	FillPrice        *big.Int        // Price at which the order was settled
	PnL              *big.Int        // PL of the previous closed position
	AccruedFunding   *big.Int        // Accrued funding of the previous closed position
	SizeDelta        *big.Int        // Size delta from the order
	NewSize          *big.Int        // New size of the position after settlement
	TotalFees        *big.Int        // Amount of fees collected by the protocol
	ReferralFees     *big.Int        // Amount of fees collected by the referrer
	CollectedFees    *big.Int        // Amount of fees collected by the fee collector
	SettlementReward *big.Int        // Amount of fees collected by the settler
	TrackingCode     [32]byte        // Optional code for integrator tracking purposes
	Settler          common.Address  // Address of the settler of the order
	// Additional fields:
	BlockNumber      uint64          // Block number where the trade was settled
	BlockTimestamp   uint64          // Timestamp of the block where the trade was settled
	TransactionHash  string          // Hash of the transaction where the trade was settled
}
```

#### RetrieveTrades

To get trades for specific block range use the RetrieveTrades function:

```go
func RetrieveTrades(fromBlock uint64, toBLock *uint64) ([]*models.Trade, error) {}
```

- Default value for `fromBlock` is a first contract block that you give in the configs
- Default value for `toBlock` is a latest blockchain block
- To use default values set `0` for `fromBlock` and `nil` for `toBlock`
- For specific block data use same values for `fromBlock` and `toBlock`
- For all contract data use default values
- For a range of blocks use required block IDs

*Warning*

If you want to query more than 20 000 block or query old block be sure you use a private PRC provider

## License
This project is licensed under the MIT License.