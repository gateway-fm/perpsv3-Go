package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
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
		contractPerm perpsMarketGoerli.IAccountModuleAccountPermissions
		want         []Permission
	}{
		{
			name: "no permissions",
		},
		{
			name: "one permission",
			contractPerm: perpsMarketGoerli.IAccountModuleAccountPermissions{
				Permissions: [][32]byte{{65, 68, 77, 73, 78}},
			},
			want: []Permission{0},
		},
		{
			name: "one bad permission",
			contractPerm: perpsMarketGoerli.IAccountModuleAccountPermissions{
				Permissions: [][32]byte{{65, 68, 77, 73, 76}},
			},
		},
		{
			name: "several permissions",
			contractPerm: perpsMarketGoerli.IAccountModuleAccountPermissions{
				Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 82, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
			},
			want: []Permission{0, 1, 2},
		},
		{
			name: "several permissions one bad",
			contractPerm: perpsMarketGoerli.IAccountModuleAccountPermissions{
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

func TestGetPermissions(t *testing.T) {
	testCases := []struct {
		name          string
		contractPerms []perpsMarketGoerli.IAccountModuleAccountPermissions
		want          []*UserPermissions
	}{
		{
			name: "no permissions",
		},
		{
			name: "one user permission no permissions",
			contractPerms: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User: common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				},
			},
			want: []*UserPermissions{
				{
					User: common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				},
			},
		},
		{
			name: "one user permission one permission",
			contractPerms: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}},
				},
			},
			want: []*UserPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: []Permission{0},
				},
			},
		},
		{
			name: "one user permission one bad permission",
			contractPerms: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: [][32]byte{{65, 68, 77, 72, 78}},
				},
			},
			want: []*UserPermissions{
				{
					User: common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				},
			},
		},
		{
			name: "one user permission several permissions",
			contractPerms: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 82, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
				},
			},
			want: []*UserPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: []Permission{0, 1, 2},
				},
			},
		},
		{
			name: "one user permission several permissions one bad",
			contractPerms: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 81, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
				},
			},
			want: []*UserPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: []Permission{0, 2},
				},
			},
		},
		{
			name: "several user permission no permissions",
			contractPerms: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User: common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				},
				{
					User: common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
				},
				{
					User: common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
				},
			},
			want: []*UserPermissions{
				{
					User: common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				},
				{
					User: common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
				},
				{
					User: common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
				},
			},
		},
		{
			name: "several user permission one permission",
			contractPerms: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}},
				},
				{
					User:        common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
					Permissions: [][32]byte{{87, 73, 84, 72, 68, 82, 65, 87}},
				},
				{
					User:        common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
					Permissions: [][32]byte{{68, 69, 76, 69, 71, 65, 84, 69}},
				},
			},
			want: []*UserPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: []Permission{0},
				},
				{
					User:        common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
					Permissions: []Permission{1},
				},
				{
					User:        common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
					Permissions: []Permission{2},
				},
			},
		},
		{
			name: "several user permission several permissions",
			contractPerms: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 82, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
				},
				{
					User:        common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 82, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
				},
				{
					User:        common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 82, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
				},
			},
			want: []*UserPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: []Permission{0, 1, 2},
				},
				{
					User:        common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
					Permissions: []Permission{0, 1, 2},
				},
				{
					User:        common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
					Permissions: []Permission{0, 1, 2},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := getPermissions(tt.contractPerms)

			require.Equal(t, tt.want, res)
		})
	}
}
