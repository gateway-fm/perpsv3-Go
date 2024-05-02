package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

type CollateralConfiguration struct {
	DepositingEnabled    bool
	IssuanceRatioD18     *big.Int
	LiquidationRatioD18  *big.Int
	LiquidationRewardD18 *big.Int
	OracleNodeId         [32]byte
	TokenAddress         common.Address
	MinDelegationD18     *big.Int
}

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

//	DepositingEnabled    bool
//	IssuanceRatioD18     *big.Int
//	LiquidationRatioD18  *big.Int
//	LiquidationRewardD18 *big.Int
//	OracleNodeId         [32]byte
//	TokenAddress         common.Address
//	MinDelegationD18     *big.Int

func GetCollateralConfigurations(data []core.CollateralConfigurationData) []*CollateralConfiguration {
	cc := make([]*CollateralConfiguration, len(data))

	for i, c := range data {
		cc[i] = GetCollateralConfiguration(c)
	}

	return cc
}

func GetCollateralConfiguration(data core.CollateralConfigurationData) *CollateralConfiguration {
	c := &CollateralConfiguration{}

	c.DepositingEnabled = data.DepositingEnabled
	c.IssuanceRatioD18 = data.IssuanceRatioD18
	c.LiquidationRatioD18 = data.LiquidationRatioD18
	c.LiquidationRewardD18 = data.LiquidationRewardD18
	c.OracleNodeId = data.OracleNodeId
	c.TokenAddress = data.TokenAddress
	c.MinDelegationD18 = data.MinDelegationD18

	return c
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
