package models

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

func TestGetAccount(t *testing.T) {
	timeNow := time.Now()

	testCases := []struct {
		name            string
		id              *big.Int
		owner           common.Address
		lastInteraction uint64
		permissions     []perpsMarketGoerli.IAccountModuleAccountPermissions
		want            *Account
	}{
		{
			name: "blank account",
			want: &Account{},
		},
		{
			name:            "no permissions",
			id:              big.NewInt(1),
			owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
			lastInteraction: uint64(timeNow.Unix()),
			want: &Account{
				ID:              big.NewInt(1),
				Owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
				LastInteraction: uint64(timeNow.Unix()),
			},
		},
		{
			name:            "one permission",
			id:              big.NewInt(1),
			owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
			lastInteraction: uint64(timeNow.Unix()),
			permissions: []perpsMarketGoerli.IAccountModuleAccountPermissions{{
				User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				Permissions: [][32]byte{{65, 68, 77, 73, 78}},
			}},
			want: &Account{
				ID:              big.NewInt(1),
				Owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
				LastInteraction: uint64(timeNow.Unix()),
				Permissions: []*UserPermissions{{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: []Permission{0},
				}},
			},
		},
		{
			name:            "several permissions",
			id:              big.NewInt(1),
			owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
			lastInteraction: uint64(timeNow.Unix()),
			permissions: []perpsMarketGoerli.IAccountModuleAccountPermissions{{
				User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 82, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
			}},
			want: &Account{
				ID:              big.NewInt(1),
				Owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
				LastInteraction: uint64(timeNow.Unix()),
				Permissions: []*UserPermissions{{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: []Permission{0, 1, 2},
				}},
			},
		},
		{
			name:            "one permission several users",
			id:              big.NewInt(1),
			owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
			lastInteraction: uint64(timeNow.Unix()),
			permissions: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}},
				},
				{
					User:        common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
					Permissions: [][32]byte{{87, 73, 84, 72, 68, 82, 65, 87}},
				},
			},
			want: &Account{
				ID:              big.NewInt(1),
				Owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
				LastInteraction: uint64(timeNow.Unix()),
				Permissions: []*UserPermissions{
					{
						User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
						Permissions: []Permission{0},
					},
					{
						User:        common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
						Permissions: []Permission{1},
					},
				},
			},
		},
		{
			name:            "several permissions several users",
			id:              big.NewInt(1),
			owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
			lastInteraction: uint64(timeNow.Unix()),
			permissions: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 82, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
				},
				{
					User:        common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 82, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
				},
			},
			want: &Account{
				ID:              big.NewInt(1),
				Owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
				LastInteraction: uint64(timeNow.Unix()),
				Permissions: []*UserPermissions{
					{
						User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
						Permissions: []Permission{0, 1, 2},
					},
					{
						User:        common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
						Permissions: []Permission{0, 1, 2},
					},
				},
			},
		},
		{
			name:            "several permissions several users one bad",
			id:              big.NewInt(1),
			owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
			lastInteraction: uint64(timeNow.Unix()),
			permissions: []perpsMarketGoerli.IAccountModuleAccountPermissions{
				{
					User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 72, 84, 72, 68, 82, 65, 87}, {68, 69, 76, 69, 71, 65, 84, 69}},
				},
				{
					User:        common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
					Permissions: [][32]byte{{65, 68, 77, 73, 78}, {87, 73, 84, 72, 68, 82, 65, 87}, {63, 69, 76, 69, 71, 65, 84, 69}},
				},
			},
			want: &Account{
				ID:              big.NewInt(1),
				Owner:           common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"),
				LastInteraction: uint64(timeNow.Unix()),
				Permissions: []*UserPermissions{
					{
						User:        common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
						Permissions: []Permission{0, 2},
					},
					{
						User:        common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"),
						Permissions: []Permission{0, 1},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := GetAccount(tt.id, tt.owner, tt.lastInteraction, tt.permissions)

			require.Equal(t, tt.want, res)
		})
	}
}
