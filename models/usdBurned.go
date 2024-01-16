package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// USDBurned is a `usdBurned` Core smart-contract event struct
type USDBurned struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	Amount         *big.Int
	Sender         common.Address
	BlockNumber    uint64
	BlockTimestamp uint64
}

// GetUSDBurnedFromEvent is used to get USDBurned struct from given contract event
func GetUSDBurnedFromEvent(event *core.CoreUsdBurned, time uint64) *USDBurned {
	if event == nil {
		logger.Log().WithField("layer", "Models-USDBurned").Warning("nil event received")
		return &USDBurned{}
	}

	return &USDBurned{
		AccountId:      event.AccountId,
		PoolId:         event.PoolId,
		CollateralType: event.CollateralType,
		Amount:         event.Amount,
		Sender:         event.Sender,
		BlockNumber:    event.Raw.BlockNumber,
		BlockTimestamp: time,
	}
}
