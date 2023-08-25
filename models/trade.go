package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
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
	AccountID        uint64
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

// ParseFromEvent is used to parse Trade model data from given event and block timestamp
func (t *Trade) ParseFromEvent(event *perpsMarketGoerli.PerpsMarketGoerliOrderSettled, time uint64) {
	t.MarketID = event.MarketId.Uint64()
	t.AccountID = event.AccountId.Uint64()
	t.FillPrice = event.FillPrice
	t.PnL = event.Pnl
	t.AccruedFunding = event.AccruedFunding
	t.SizeDelta = event.SizeDelta
	t.NewSize = event.NewSize
	t.TotalFees = event.TotalFees
	t.ReferralFees = event.ReferralFees
	t.CollectedFees = event.CollectedFees
	t.SettlementReward = event.SettlementReward
	t.TrackingCode = event.TrackingCode
	t.Settler = event.Settler
	t.BlockNumber = event.Raw.BlockNumber
	t.BlockTimestamp = time
	t.TransactionHash = event.Raw.TxHash.Hex()
}
