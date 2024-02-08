package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// DelegationUpdated is a `DelegationUpdated` Core smart-contract event struct
type DelegationUpdated struct {
	AccountId      *big.Int
	PoolId         *big.Int
	CollateralType common.Address
	Amount         *big.Int
	Leverage       *big.Int
	Sender         common.Address
	BlockNumber    uint64
	BlockTimestamp uint64
}

// GetDelegationUpdatedFromEvent is used to get DelegationUpdated struct from given contract event
func GetDelegationUpdatedFromEvent(event *core.CoreDelegationUpdated, time uint64) *DelegationUpdated {
	if event == nil {
		logger.Log().WithField("layer", "Models-CoreDelegationUpdated").Warning("nil event received")
		return &DelegationUpdated{}
	}

	return &DelegationUpdated{
		AccountId:      event.AccountId,
		PoolId:         event.PoolId,
		CollateralType: event.CollateralType,
		Amount:         event.Amount,
		Leverage:       event.Leverage,
		Sender:         event.Sender,
		BlockNumber:    event.Raw.BlockNumber,
		BlockTimestamp: time,
	}
}
