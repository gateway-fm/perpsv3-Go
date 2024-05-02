package models

import (
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
)

// Account is a struct for account model
//   - ID is an account NFT id
//   - Permissions is a slice of UserPermissions struct with granted permissions data
//   - Owner is a contract owner address
//   - LastInteraction is a unix timestamp for last accounts contract interaction
type Account struct {
	ID              *big.Int
	Permissions     []*UserPermissions
	Owner           common.Address
	LastInteraction uint64
}

// AccountLiquidated is a struct for `AccountLiquidated` event
//   - ID is an account NFT id
//   - Reward is a liquidation reward transferred to caller
//   - FullLiquidated is a filed for define is account fully liquidated or not
type AccountLiquidated struct {
	ID             *big.Int
	Reward         *big.Int
	FullLiquidated bool
}

type AccountCollateral struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}

func GetAccountCollateralFromContract(res struct {
	TotalDeposited *big.Int
	TotalAssigned  *big.Int
	TotalLocked    *big.Int
}) *AccountCollateral {
	ac := &AccountCollateral{}

	ac.TotalLocked = res.TotalLocked
	ac.TotalAssigned = res.TotalAssigned
	ac.TotalDeposited = res.TotalDeposited

	return ac
}

// FormatAccount is used to get account from given data
func FormatAccount(
	id *big.Int,
	owner common.Address,
	lastInteraction uint64,
	permissions []perpsMarket.IAccountModuleAccountPermissions,
) *Account {
	return &Account{
		ID:              id,
		Permissions:     getUserPermissions(permissions),
		Owner:           owner,
		LastInteraction: lastInteraction,
	}
}

func FormatAccountCore(
	id *big.Int,
	owner common.Address,
	lastInteraction uint64,
	permissions []core.IAccountModuleAccountPermissions,
) *Account {
	return &Account{
		ID:              id,
		Permissions:     getUserPermissionsCore(permissions),
		Owner:           owner,
		LastInteraction: lastInteraction,
	}
}
