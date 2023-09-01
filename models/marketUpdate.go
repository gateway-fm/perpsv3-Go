package models

import (
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// MarketUpdate
//   - MarketID: ID of the market.
//   - Price: Price at the time of the event.
//   - Skew: Market skew at the time of the event. Positive values indicate more longs.
//   - Size: Size of the entire market after settlement.
//   - SizeDelta: Change in market size during the update.
//   - CurrentFundingRate: Current funding rate of the market.
//   - CurrentFundingVelocity: Current rate of change of the funding rate.
//   - BlockNumber: Block number at which the market data was fetched.
//   - BlockTimestamp: Timestamp of the block at which the market data was fetched.
//   - TransactionHash: Hash of the transaction where the market update occurred.
type MarketUpdate struct {
	MarketID               uint64
	Price                  uint64
	Skew                   int64
	Size                   uint64
	SizeDelta              int64
	CurrentFundingRate     int64
	CurrentFundingVelocity int64
	BlockNumber            uint64
	BlockTimestamp         uint64
	TransactionHash        string
}

// GetMarketUpdateFromEvent is used to get MarketUpdate struct from given event and block timestamp
func GetMarketUpdateFromEvent(event *perpsMarketGoerli.PerpsMarketGoerliMarketUpdated, time uint64) *MarketUpdate {
	if event == nil {
		logger.Log().WithField("layer", "Models-MarketUpdate").Warning("nil event received")
		return &MarketUpdate{}
	}

	marketID := uint64(0)
	if event.MarketId != nil {
		marketID = event.MarketId.Uint64()
	}

	price := uint64(0)
	if event.Price != nil {
		price = event.Price.Uint64()
	}

	skew := int64(0)
	if event.Skew != nil {
		skew = event.Skew.Int64()
	}

	size := uint64(0)
	if event.Size != nil {
		size = event.Size.Uint64()
	}

	sizeDelta := int64(0)
	if event.SizeDelta != nil {
		sizeDelta = event.SizeDelta.Int64()
	}

	currentFundingRate := int64(0)
	if event.CurrentFundingRate != nil {
		currentFundingRate = event.CurrentFundingRate.Int64()
	}

	currentFundingVelocity := int64(0)
	if event.CurrentFundingVelocity != nil {
		currentFundingVelocity = event.CurrentFundingVelocity.Int64()
	}

	return &MarketUpdate{
		MarketID:               marketID,
		Price:                  price,
		Skew:                   skew,
		Size:                   size,
		SizeDelta:              sizeDelta,
		CurrentFundingRate:     currentFundingRate,
		CurrentFundingVelocity: currentFundingVelocity,
		BlockNumber:            event.Raw.BlockNumber,
		BlockTimestamp:         time,
		TransactionHash:        event.Raw.TxHash.Hex(),
	}
}
