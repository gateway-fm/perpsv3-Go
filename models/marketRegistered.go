package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
)

type MarketRegistered struct {
	Market         common.Address
	MarketID       *big.Int
	Sender         common.Address
	BlockNumber    uint64
	BlockTimestamp uint64
}

func GetMarketRegisteredFromEvent(event *core.CoreMarketRegistered, time uint64) *MarketRegistered {
	m := &MarketRegistered{}

	m.Market = event.Market
	m.MarketID = event.MarketId
	m.Sender = event.Sender
	m.BlockNumber = event.Raw.BlockNumber
	m.BlockTimestamp = time

	return m
}
