package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

type RewardDistributed struct {
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Amount         *big.Int
	Start          *big.Int
	Duration       *big.Int
	BlockNumber    uint64
	BlockTimestamp uint64
}

func GetRewardDistributedFromEvent(event *core.CoreRewardsDistributed, time uint64) *RewardDistributed {
	if event == nil {
		logger.Log().WithField("layer", "Models-RewardDistributed").Warning("nil event received")
		return &RewardDistributed{}
	}

	return &RewardDistributed{
		PoolId:         event.PoolId,
		CollateralType: event.CollateralType,
		Distributor:    event.Distributor,
		Amount:         event.Amount,
		Start:          event.Start,
		Duration:       event.Duration,
		BlockNumber:    event.Raw.BlockNumber,
		BlockTimestamp: time,
	}
}
