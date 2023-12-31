package models

import (
	"math/big"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// Liquidation is a position liquidation model
//   - MarketID: ID of the market used for the order.
//   - AccountID: ID of the account used for the order.
//   - AmountLiquidated: amount liquidated.
//   - CurrentPositionSize: position size after liquidation.
//   - BlockNumber: Block number where the order was committed.
//   - BlockTimestamp: Timestamp of the block where the order was committed.
type Liquidation struct {
	MarketID            uint64
	AccountID           *big.Int
	AmountLiquidated    *big.Int
	CurrentPositionSize *big.Int
	BlockNumber         uint64
	BlockTimestamp      uint64
}

// GetLiquidationFromEvent is used to get Liquidation struct from given contract event
func GetLiquidationFromEvent(event *perpsMarket.PerpsMarketPositionLiquidated, time uint64) *Liquidation {
	if event == nil {
		logger.Log().WithField("layer", "Models-Liquidation").Warning("nil event received")
		return &Liquidation{}
	}

	marketID := uint64(0)
	if event.MarketId != nil {
		marketID = event.MarketId.Uint64()
	}

	return &Liquidation{
		MarketID:            marketID,
		AccountID:           event.AccountId,
		AmountLiquidated:    event.AmountLiquidated,
		CurrentPositionSize: event.CurrentPositionSize,
		BlockNumber:         event.Raw.BlockNumber,
		BlockTimestamp:      time,
	}
}
