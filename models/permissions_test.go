package models

import (
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPermissionFromString(t *testing.T) {
	testCases := []struct {
		name     string
		permName string
		want     Permission
		wantErr  error
	}{
		{
			name:     "admin",
			permName: "ADMIN",
			want:     Permission(0),
		},
		{
			name:     "withdraw",
			permName: "WITHDRAW",
			want:     Permission(1),
		},
		{
			name:     "delegate",
			permName: "DELEGATE",
			want:     Permission(2),
		},
		{
			name:     "mint",
			permName: "MINT",
			want:     Permission(3),
		},
		{
			name:     "rewards",
			permName: "REWARDS",
			want:     Permission(4),
		},
		{
			name:     "perps modify collateral",
			permName: "PERPS_MODIFY_COLLATERAL",
			want:     Permission(5),
		},
		{
			name:     "perps commit async order",
			permName: "PERPS_COMMIT_ASYNC_ORDER",
			want:     Permission(6),
		},
		{
			name:     "unsupported error",
			permName: "unsupported",
			want:     Permission(7),
			wantErr:  errors.EnumUnsupportedErr,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res, err := PermissionFromString(tt.permName)

			require.Equal(t, tt.want, res)

			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tt.wantErr)
			}

		})
	}
}

func TestDecodePermissions(t *testing.T) {
	testCases := []struct {
		name         string
		contractPerm perpsMarket.IAccountModuleAccountPermissions
		want         []Permission
	}{
		{
			name: "no permissions",
		},
		{
			name: "one permission",
			contractPerm: perpsMarket.IAccountModuleAccountPermissions{
				Permissions: [][32]byte{{65, 68, 77, 73, 78}},
			},
			want: []Permission{0},
		},
		{
			name: "one bad permission",
			contractPerm: perpsMarket.IAccountModuleAccountPermissions{
				Permissions: [][32]byte{{65, 68, 77, 73, 76}},
			},
		},
		{
			name: "several permissions",
			contractPerm: perpsMarket.IAccountModuleAccountPermissions{
				Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 82, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
			},
			want: []Permission{0, 1, 2},
		},
		{
			name: "several permissions one bad",
			contractPerm: perpsMarket.IAccountModuleAccountPermissions{
				Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 81, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
			},
			want: []Permission{0, 2},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := decodePermissions(tt.contractPerm)

			require.Equal(t, tt.want, res)
		})
	}
}
