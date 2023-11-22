package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// Trade is a trade event model
//   - MarketID         - ID of the market used for the trade
//   - AccountID        - ID of the account used for the trade
//   - FillPrice        - Price at which the order was settled
//   - PnL              - PL of the previous closed position
//   - AccruedFunding   - Accrued funding of the previous closed position
//   - SizeDelta        - Size delta from the order
//   - NewSize          - New size of the position after settlement
//   - TotalFees        - Amount of fees collected by the protocol
//   - ReferralFees     - Amount of fees collected by the referrer
//   - CollectedFees    - Amount of fees collected by the fee collector
//   - SettlementReward - Amount of fees collected by the settler
//   - TrackingCode     - Optional code for integrator tracking purposes
//   - Settler          - Address of the settler of the order
//   - BlockNumber      - Block number where the trade was settled.
//   - BlockTimestamp   - Timestamp of the block where the trade was settled.
//   - TransactionHash  - Hash of the transaction where the trade was settled.
type Trade struct {
	MarketID         uint64
	AccountID        *big.Int
	FillPrice        *big.Int
	PnL              *big.Int
	AccruedFunding   *big.Int
	SizeDelta        *big.Int
	NewSize          *big.Int
	TotalFees        *big.Int
	ReferralFees     *big.Int
	CollectedFees    *big.Int
	SettlementReward *big.Int
	TrackingCode     [32]byte
	Settler          common.Address
	BlockNumber      uint64
	BlockTimestamp   uint64
	TransactionHash  string
}

// GetTradeFromEvent is used to get new Trade from given event and block timestamp
func GetTradeFromEvent(event *perpsMarket.PerpsMarketOrderSettled, time uint64) *Trade {
	if event == nil {
		logger.Log().WithField("layer", "Models-Trade").Warning("nil event received")
		return &Trade{}
	}

	marketID := uint64(0)
	if event.MarketId != nil {
		marketID = event.MarketId.Uint64()
	}

	return &Trade{
		MarketID:         marketID,
		AccountID:        event.AccountId,
		FillPrice:        event.FillPrice,
		PnL:              event.Pnl,
		AccruedFunding:   event.AccruedFunding,
		SizeDelta:        event.SizeDelta,
		NewSize:          event.NewSize,
		TotalFees:        event.TotalFees,
		ReferralFees:     event.ReferralFees,
		CollectedFees:    event.CollectedFees,
		SettlementReward: event.SettlementReward,
		TrackingCode:     event.TrackingCode,
		Settler:          event.Settler,
		BlockNumber:      event.Raw.BlockNumber,
		BlockTimestamp:   time,
		TransactionHash:  event.Raw.TxHash.Hex(),
	}
}
