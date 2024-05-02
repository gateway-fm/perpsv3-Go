package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
)

type CoreVaultLiquidation struct {
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

func GetCoreVaultLiquidationFromEvent(event *core.CoreVaultLiquidation, time uint64) *CoreVaultLiquidation {
	c := &CoreVaultLiquidation{}

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
