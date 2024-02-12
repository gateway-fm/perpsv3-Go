package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

type RewardClaimed struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	Distributor    common.Address
	Amount         *big.Int
	BlockNumber    uint64
	BlockTimestamp uint64
}

func GetRewardClaimedFromEvent(event *core.CoreRewardsClaimed, time uint64) *RewardClaimed {
	if event == nil {
		logger.Log().WithField("layer", "Models-RewardClaimed").Warning("nil event received")
		return &RewardClaimed{}
	}

	return &RewardClaimed{
		AccountId:      event.AccountId,
		PoolId:         event.PoolId,
		CollateralType: event.CollateralType,
		Distributor:    event.Distributor,
		Amount:         event.Amount,
		BlockNumber:    event.Raw.BlockNumber,
		BlockTimestamp: time,
	}
}
