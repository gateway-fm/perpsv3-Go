package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// CollateralDeposited is a `Deposited` Core smart-contract event struct
type CollateralDeposited struct {
	AccountId      *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	BlockNumber    uint64
	BlockTimestamp uint64
}

// CollateralWithdrawn is a `Withdrawn` Core smart-contract event struct
type CollateralWithdrawn struct {
	AccountId      *big.Int
	CollateralType common.Address
	TokenAmount    *big.Int
	Sender         common.Address
	BlockNumber    uint64
	BlockTimestamp uint64
}

// CollateralPrice is a collateral price data struct
type CollateralPrice struct {
	Price *big.Int
}

// GetCollateralDepositedFromEvent is used to get CollateralDeposited struct from given contract event
func GetCollateralDepositedFromEvent(event *core.CoreDeposited, time uint64) *CollateralDeposited {
	if event == nil {
		logger.Log().WithField("layer", "Models-CoreDeposited").Warning("nil event received")
		return &CollateralDeposited{}
	}

	return &CollateralDeposited{
		AccountId:      event.AccountId,
		CollateralType: event.CollateralType,
		TokenAmount:    event.TokenAmount,
		Sender:         event.Sender,
		BlockNumber:    event.Raw.BlockNumber,
		BlockTimestamp: time,
	}
}

// GetCollateralWithdrawnFromEvent is used to get CollateralWithdrawn struct from given contract event
func GetCollateralWithdrawnFromEvent(event *core.CoreWithdrawn, time uint64) *CollateralWithdrawn {
	if event == nil {
		logger.Log().WithField("layer", "Models-CoreWithdrawn").Warning("nil event received")
		return &CollateralWithdrawn{}
	}

	return &CollateralWithdrawn{
		AccountId:      event.AccountId,
		CollateralType: event.CollateralType,
		TokenAmount:    event.TokenAmount,
		Sender:         event.Sender,
		BlockNumber:    event.Raw.BlockNumber,
		BlockTimestamp: time,
	}
}
