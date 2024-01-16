package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// USDMinted is a `usdMinted` Core smart-contract event struct
type USDMinted struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	Amount         *big.Int
	Sender         common.Address
	BlockNumber    uint64
	BlockTimestamp uint64
}

// GetUSDMintedFromEvent is used to get USDMinted struct from given contract event
func GetUSDMintedFromEvent(event *core.CoreUsdMinted, time uint64) *USDMinted {
	if event == nil {
		logger.Log().WithField("layer", "Models-USDMinted").Warning("nil event received")
		return &USDMinted{}
	}

	return &USDMinted{
		AccountId:      event.AccountId,
		PoolId:         event.PoolId,
		CollateralType: event.CollateralType,
		Amount:         event.Amount,
		Sender:         event.Sender,
		BlockNumber:    event.Raw.BlockNumber,
		BlockTimestamp: time,
	}
}
