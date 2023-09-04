package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
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

// FormatAccount is used to get account from given data
func FormatAccount(
	id *big.Int,
	owner common.Address,
	lastInteraction uint64,
	permissions []perpsMarketGoerli.IAccountModuleAccountPermissions,
) *Account {
	return &Account{
		ID:              id,
		Permissions:     getUserPermissions(permissions),
		Owner:           owner,
		LastInteraction: lastInteraction,
	}
}
