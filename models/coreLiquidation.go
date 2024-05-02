package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
)

type CoreLiquidation struct {
	AccountId            *big.Int
	PoolId               *big.Int
	CollateralType       common.Address
	DebtLiquidated       *big.Int
	CollateralLiquidated *big.Int
	AmountRewarded       *big.Int
	LiquidateAsAccountId *big.Int
	Sender               common.Address
	BlockNumber          uint64
	BlockTimestamp       uint64
}

func GetCoreLiquidationFromEvent(event *core.CoreLiquidation, time uint64) *CoreLiquidation {
	c := &CoreLiquidation{}

	c.AccountId = event.AccountId
	c.PoolId = event.PoolId
	c.CollateralType = event.CollateralType
	c.DebtLiquidated = event.LiquidationData.DebtLiquidated
	c.CollateralLiquidated = event.LiquidationData.CollateralLiquidated
	c.AmountRewarded = event.LiquidationData.AmountRewarded
	c.LiquidateAsAccountId = event.LiquidateAsAccountId
	c.Sender = event.Sender
	c.BlockNumber = event.Raw.BlockNumber
	c.BlockTimestamp = time

	return c
}
