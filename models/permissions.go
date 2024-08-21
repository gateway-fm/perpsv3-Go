package models

import (
	"strings"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// Permission is a permission enum
type Permission int

const (
	ADMIN Permission = iota
	WITHDRAW
	DELEGATE
	MINT
	REWARDS
	PERPS_MODIFY_COLLATERAL
	PERPS_COMMIT_ASYNC_ORDER
	BURN
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
	BURN:                     "BURN",
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

// decodePermissions is used to decode given contract permissions to Permission slice
func decodePermissions(perm perpsMarket.IAccountModuleAccountPermissions) (res []Permission) {
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

func decodePermissionsCore(perm core.IAccountModuleAccountPermissions) (res []Permission) {
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
