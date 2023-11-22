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

#### RetrieveTrades()

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

#### RetrieveTradesLimit()

To get all trades with RPC provider block limiration use the RetrieveTrades function:

```go
func RetrieveTradesLimit(limit uint64) ([]*models.Trade, error)
```

The function will query blocks several times from the first contract block to the latest block until all blocks are queried.
If the contract was deployed a long time ago the function can take more than 1 minute to work.

- Default value for `limit` is a 20 000 blocks per one query

#### ListenTrades()

To subscribe on the contract `OrederSettled` event use the ListenTrades function.

```go
func ListenTrades() (*events.TradeSubscription, error) {}
```

The goroutine will return events as a `Trade` model on the `TradesChan` chanel and errors on the `ErrChan` chanel. To
close the subscription use the `Close` function.

You can see an [example](examples/trades_events.go) of the usage here:

```go
package main

import (
	"fmt"
	"log"
	"time"

	perpsv3_Go "github.com/gateway-fm/perpsv3-Go"
	"github.com/gateway-fm/perpsv3-Go/config"
)

func main() {
	lib, err := perpsv3_Go.Create(config.GetGoerliDefaultPerpsvConfig())
	if err != nil {
		log.Fatal(err)
	}

	subs, err := lib.ListenTrades()
	if err != nil {
		log.Fatal(err)
	}

	stopChan := make(chan struct{})

	// handle events
	go func() {
		for {
			select {
			case <-stopChan:
				subs.Close()
				return
			case err = <-subs.ErrChan:
				log.Println(err.Error())
			case trade := <-subs.TradesChan:
				fmt.Println(trade.AccountID)
				fmt.Println(trade.AccruedFunding)
			}
		}
	}()

	time.Sleep(10 * time.Second)

	// stop listening
	close(stopChan)

	time.Sleep(5 * time.Second)

}
```

### Orders

#### Model

Using Orders services you operate with Orders model which represents a `OrderCommitted` event of Perps Market smart-contract
with some additional fields:

```go
type Order struct {
    // Event fields:
	MarketID        uint64          // ID of the market used for the trade
	AccountID       uint64          // ID of the account used for the trade
    OrderType       uint8           // Represents the transaction type (0 at the time of writing)
	SizeDelta       *big.Int        // Requested change in size of the order
    AcceptablePrice *big.Int        // Maximum or minimum accepted price to settle the order.
    SettlementTime  uint64          // Time at which the order can be settled.
    ExpirationTime  uint64          // Time at which the order expired.
    TrackingCode    [32]byte        // Optional code for integrator tracking purposes.
    Sender          common.Address  // Address of the sender of the order.
    // Additional fields:
    BlockNumber      uint64         // Block number where the trade was settled
    BlockTimestamp   uint64         // Timestamp of the block where the trade was settled
}
```

#### RetrieveOrders()

To get orders for specific block range use the RetrieveOrders function:

```go
func RetrieveOrders(fromBlock uint64, toBLock *uint64) ([]*models.Order, error) {}
```

- Default value for `fromBlock` is a first contract block that you give in the configs
- Default value for `toBlock` is a latest blockchain block
- To use default values set `0` for `fromBlock` and `nil` for `toBlock`
- For specific block data use same values for `fromBlock` and `toBlock`
- For all contract data use default values
- For a range of blocks use required block IDs

*Warning*

If you want to query more than 20 000 block or query old block be sure you use a private PRC provider

#### RetrieveOrdersLimit()

To get all orders with RPC provider block limitation use the RetrieveOrdersLimit function:

```go
func RetrieveOrdersLimit(limit uint64) ([]*models.Order, error) {}
```

The function will query blocks several times from the first contract block to the latest block until all blocks are queried.
If the contract was deployed a long time ago the function can take more than 1 minute to work.

- Default value for `limit` is a 20 000 blocks per one query


#### ListenOrders()

To subscribe on the contract `OrederCommitted` event use the ListenOrders function.

```go
func ListenOrders() (*events.OrderSubscription, error) {}
```

The goroutine will return events as a `Order` model on the `OrdersChan` chanel and errors on the `ErrChan` chanel. To
close the subscription use the `Close` function.

### Markets

You can query current market IDs, current market metadata and current market summary

#### Models

```go
type MarketMetadata struct {
    MarketID *big.Int // is a market ID value
    Name     string   // is a market name value
    Symbol   string   // is a market symbol value for example 'ETH'
}
```

```go
type MarketSummary struct {
	MarketID               *big.Int // Represents the ID of the market
	Skew                   *big.Int // Represents the skew of the market
	Size                   *big.Int // Represents the size of the market
	MaxOpenInterest        *big.Int // Represents the maximum open interest of the market
	CurrentFundingRate     *big.Int // Represents the current funding rate of the market
	CurrentFundingVelocity *big.Int // Represents the current funding velocity of the market
	IndexPrice             *big.Int // Represents the index price of the market
}
```

#### GetMarketIds()

To get available market IDs use GetMarketIDs function

```go
func GetMarketIDs() ([]*big.Int, error) {}
```

#### GetMarketMetadata()

To get current market metadata by given market ID use GetMarketMetadata function

```go
func GetMarketMetadata(marketID *big.Int) (*models.MarketMetadata, error) {}
```

#### GetMarketSummary()

To get current market summary by given market ID use GetMarketSummary function

```go
func GetMarketSummary(marketID *big.Int) (*models.MarketSummary, error) {}
```

#### GetFoundingRate()

To get current founding rate by given market ID use GetFoundingRate function

```go
func GetFoundingRate(id *big.Int) (*big.Int, error) {}
```

### MarketUpdate

#### Models

Using MarketData services you operate with MarketUpdate model which represents a `MarketUpdated` event of Perps Market smart-contract
with some additional fields:

```go
type MarketUpdate struct {
	// Event fields
    MarketID               uint64  // ID of the market.
    Price                  uint64  // Price at the time of the event.
    Skew                   int64   // Market skew at the time of the event. Positive values indicate more longs.
    Size                   uint64  // Size of the entire market after settlement.
    SizeDelta              int64   // Change in market size during the update.
    CurrentFundingRate     int64   // Current funding rate of the market.
    CurrentFundingVelocity int64   // Current rate of change of the funding rate.
	// Additional fields
    BlockNumber            uint64  // Block number at which the market data was fetched.
    BlockTimestamp         uint64  // Timestamp of the block at which the market data was fetched.
    TransactionHash        string  // Hash of the transaction where the market update occurred.
}
```

You can also use MarketDataBig model, it will operate with big.Int value types instead of uint64 and int64. Only methods
with `Big` suffix can operate with this model

```go
type MarketUpdateBig struct {
	MarketID               *big.Int
	Price                  *big.Int
	Skew                   *big.Int
	Size                   *big.Int
	SizeDelta              *big.Int
	CurrentFundingRate     *big.Int
	CurrentFundingVelocity *big.Int
	BlockNumber            uint64
	BlockTimestamp         uint64
	TransactionHash        string
}
```

#### RetrieveMarketUpdates() / RetrieveMarketUpdatesBig()

To get market update data for specific block range use the RetrieveMarketUpdates or RetrieveMarketUpdatesBig functions:

```go
func RetrieveMarketUpdates(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdate, error) {}
func RetrieveMarketUpdatesBig(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdateBig, error) {}
```

- Default value for `fromBlock` is a first contract block that you give in the configs
- Default value for `toBlock` is a latest blockchain block
- To use default values set `0` for `fromBlock` and `nil` for `toBlock`
- For specific block data use same values for `fromBlock` and `toBlock`
- For all contract data use default values
- For a range of blocks use required block IDs

*Warning*

If you want to query more than 20 000 block or query old block be sure you use a private PRC provider

#### RetrieveMarketUpdatesLimit() / RetrieveMarketUpdatesBigLimit()

To get all market updates with RPC provider block limitation use the RetrieveMarketUpdatesLimit or RetrieveMarketUpdatesBigLimit functions:

```go
func RetrieveMarketUpdatesLimit(limit uint64) ([]*models.MarketUpdate, error) {}
func RetrieveMarketUpdatesBigLimit(limit uint64) ([]*models.MarketUpdateBig, error) {}
```

The function will query blocks several times from the first contract block to the latest block until all blocks are queried.
If the contract was deployed a long time ago the function can take more than 1 minute to work.

- Default value for `limit` is a 20 000 blocks per one query


#### ListenMarketUpdate() / ListenMarketUpdatesBig()

To subscribe on the contract `MarketUpdated` event use the ListenMarketUpdates or ListenMarketUpdatesBig functions.

```go
func ListenMarketUpdates() (*events.MarketUpdateSubscription, error)
func ListenMarketUpdatesBig() (*MarketUpdateSubscriptionBig, error)
```

The goroutine will return events as a `MarketUpdate` or `MarketUpdateBig` model on the `MarketUpdateChan` chanel and errors on the `ErrChan` chanel. To
close the subscription use the `Close` function.

### Positions

#### Model

Using Positions services you operate with Position model which represents a `OpenPosition` data struct of Perps Market 
smart-contract with some additional fields:

```go
type Position struct {
	// Data from the contract
	TotalPnl       *big.Int // Represents the total profit and loss for the position
	AccruedFunding *big.Int // Represents the accrued funding for the position
	PositionSize   *big.Int // Represents the size of the position
	// Data from the latest block
	BlockNumber    uint64   // Represents the block number at which the position data was fetched
	BlockTimestamp uint64   // Represents the timestamp of the block at which the position data was fetched
}
```

#### GetPosition()

To get `Position` by reading contract with `getOpenPosition` method use the GetPosition function:

```go
func GetPosition(accountID *big.Int, marketID *big.Int) (*models.Position, error) {}
```

It will return data from the contract in the latest block. Function can return contract error if the market ID is invalid.
If account ID is invalid it will return model with blank fields.

### Liquidations

#### Model

Using Liquidations services you operate with Liquidation model which represents a `PositionLiquidated` event of Perps Market smart-contract
with some additional fields:

```go
type Liquidation struct {
	// Event fields
	MarketID            uint64   // ID of the market used for the order.
	AccountID           uint64   // ID of the account used for the order.
	AmountLiquidated    *big.Int // amount liquidated.
	CurrentPositionSize *big.Int // position size after liquidation.
	// Additional fields
	BlockNumber         uint64   // Block number where the order was committed.
	BlockTimestamp      uint64   // Timestamp of the block where the order was committed.
}
```

#### RetrieveLiquidations()

To get liquidations for specific block range use the RetrieveLiquidations function:

```go
func RetrieveLiquidations(fromBlock uint64, toBLock *uint64) ([]*models.Liquidation, error) {}
```

- Default value for `fromBlock` is a first contract block that you give in the configs
- Default value for `toBlock` is a latest blockchain block
- To use default values set `0` for `fromBlock` and `nil` for `toBlock`
- For specific block data use same values for `fromBlock` and `toBlock`
- For all contract data use default values

*Warning*

If you want to query more than 20 000 block or query old block be sure you use a private PRC provider

#### RetrieveLiquidationsLimit()

To get all luquidations with RPC provider block limitation use the RetrieveLiquidationsLimit function:

```go
func RetrieveLiquidationsLimit(limit uint64) ([]*models.Liquidation, error) {}
```

The function will query blocks several times from the first contract block to the latest block until all blocks are queried.
If the contract was deployed a long time ago the function can take more than 1 minute to work.

- Default value for `limit` is a 20 000 blocks per one query

#### ListenLiquidations()

To subscribe on the contract `PositionLiquidated` event use the ListenLiquidations function.

```go
func ListenLiquidations() (*LiquidationSubscription, error) {}
```

The goroutine will return events as a `Liquidation` model on the `LiquidationsChan` chanel and errors on the `ErrChan` chanel. To
close the subscription use the `Close` function.

## License
This project is licensed under the MIT License.# asatruPythonE2E
