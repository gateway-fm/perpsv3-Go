package models

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
)

// UserPermissions is a struct for permissions granted by account owner to User with a list of Permissions
type UserPermissions struct {
	User        common.Address
	Permissions []Permission
}

// getUserPermissions is used to get UserPermissions slice from given contract user permissions slice
func getUserPermissions(perms []perpsMarketGoerli.IAccountModuleAccountPermissions) (res []*UserPermissions) {
	for _, p := range perms {
		perm := &UserPermissions{
			User:        p.User,
			Permissions: decodePermissions(p),
		}

		res = append(res, perm)
	}

	return res
}