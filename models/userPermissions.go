package models

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
)

// UserPermissions is a struct for permissions granted by account owner to User with a list of Permissions
type UserPermissions struct {
	User        common.Address
	Permissions []Permission
}

// PermissionChanged is a struct for `PermissionRevoked` and `PermissionGranted` contract events
type PermissionChanged struct {
	AccountID  *big.Int
	User       common.Address
	Permission Permission
}

// getUserPermissions is used to get UserPermissions slice from given contract user permissions slice
func getUserPermissions(perms []perpsMarket.IAccountModuleAccountPermissions) (res []*UserPermissions) {
	for _, p := range perms {
		perm := &UserPermissions{
			User:        p.User,
			Permissions: decodePermissions(p),
		}

		res = append(res, perm)
	}

	return res
}
