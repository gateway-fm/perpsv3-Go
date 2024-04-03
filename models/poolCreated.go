package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
)

type PoolCreated struct {
	PoolID         *big.Int
	Owner          common.Address
	Sender         common.Address
	BlockNumber    uint64
	BlockTimestamp uint64
}

func GetPoolCreatedFromEvent(event *core.CorePoolCreated, time uint64) *PoolCreated {
	p := &PoolCreated{}

	p.PoolID = event.PoolId
	p.Owner = event.Owner
	p.Sender = event.Sender
	p.BlockNumber = event.Raw.BlockNumber
	p.BlockTimestamp = time

	return p
}
