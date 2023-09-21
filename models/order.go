package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// Order is an order event model
//   - MarketID: ID of the market used for the order.
//   - AccountID: ID of the account used for the order.
//   - OrderType: Represents the transaction type (0 at the time of writing).
//   - SizeDelta: Requested change in size of the order.
//   - AcceptablePrice: Maximum or minimum accepted price to settle the order.
//   - SettlementTime: Time at which the order can be settled.
//   - ExpirationTime: Time at which the order expired.
//   - TrackingCode: Optional code for integrator tracking purposes.
//   - Sender: Address of the sender of the order.
//   - BlockNumber: Block number where the order was committed.
//   - BlockTimestamp: Timestamp of the block where the order was committed.
type Order struct {
	MarketID        uint64
	AccountID       *big.Int
	OrderType       uint8
	SizeDelta       *big.Int
	AcceptablePrice *big.Int
	SettlementTime  uint64
	ExpirationTime  uint64
	TrackingCode    [32]byte
	Sender          common.Address
	BlockNumber     uint64
	BlockTimestamp  uint64
}

// GetOrderFromEvent is used to get Order struct from given event and block timestamp
func GetOrderFromEvent(event *perpsMarketGoerli.PerpsMarketGoerliOrderCommitted, time uint64) *Order {
	if event == nil {
		logger.Log().WithField("layer", "Models-Order").Warning("nil event received")
		return &Order{}
	}

	marketID := uint64(0)
	if event.MarketId != nil {
		marketID = event.MarketId.Uint64()
	}

	settlementTime := uint64(0)
	if event.SettlementTime != nil {
		settlementTime = event.SettlementTime.Uint64()
	}

	expirationTime := uint64(0)
	if event.ExpirationTime != nil {
		expirationTime = event.ExpirationTime.Uint64()
	}

	return &Order{
		MarketID:        marketID,
		AccountID:       event.AccountId,
		OrderType:       event.OrderType,
		SizeDelta:       event.SizeDelta,
		AcceptablePrice: event.AcceptablePrice,
		SettlementTime:  settlementTime,
		ExpirationTime:  expirationTime,
		TrackingCode:    event.TrackingCode,
		Sender:          event.Sender,
		BlockNumber:     event.Raw.BlockNumber,
		BlockTimestamp:  time,
	}
}
