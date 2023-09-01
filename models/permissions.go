package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
	"strings"
)

// Permission is a permission enum
type Permission int

// UserPermissions is a struct for permissions granted by account owner to User with a list of Permissions
type UserPermissions struct {
	User        common.Address
	Permissions []Permission
}

const (
	ADMIN Permission = iota
	WITHDRAW
	DELEGATE
	MINT
	REWARDS
	PERPS_MODIFY_COLLATERAL
	PERPS_COMMIT_ASYNC_ORDER
	unsupported
)

// permissionsS is mapping Permission to its string value
var permissionsS = [...]string{
	ADMIN:                    "ADMIN",
	WITHDRAW:                 "WITHDRAW",
	DELEGATE:                 "DELEGATE",
	MINT:                     "MINT",
	REWARDS:                  "REWARDS",
	PERPS_MODIFY_COLLATERAL:  "PERPS_MODIFY_COLLATERAL",
	PERPS_COMMIT_ASYNC_ORDER: "PERPS_COMMIT_ASYNC_ORDER",
}

// String is used to return Permission string value
func (p Permission) String() string {
	return permissionsS[p]
}

// PermissionFromString is used to get Permission from string value mapped in permissionsS
func PermissionFromString(s string) (Permission, error) {
	for i, r := range permissionsS {
		if s == r {
			return Permission(i), nil
		}
	}

	logger.Log().WithField("layer", "Model-PermissionFromString").Warningf("got usupported value: %v", s)
	return unsupported, errors.GetUnsupportedErr("permissions")
}

// getPermissions is used to get UserPermissions slice from given contract user permissions slice
func getPermissions(perms []perpsMarketGoerli.IAccountModuleAccountPermissions) (res []*UserPermissions) {
	for _, p := range perms {
		perm := &UserPermissions{
			User:        p.User,
			Permissions: decodePermissions(p),
		}

		res = append(res, perm)
	}

	return res
}

// decodePermissions is used to decode given contract permissions to Permission slice
func decodePermissions(perm perpsMarketGoerli.IAccountModuleAccountPermissions) (res []Permission) {
	for _, b := range perm.Permissions {
		p, err := PermissionFromString(strings.TrimRight(string(b[:]), string(rune(0))))
		if err != nil {
			logger.Log().WithField("layer", "Model-decodePermissions").Warningf("received unsupported bytes value %v", string(b[:]))
		} else {
			res = append(res, p)
		}
	}

	return res
}
