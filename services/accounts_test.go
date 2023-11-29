package services

import (
	"github.com/gateway-fm/perpsv3-Go/config"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/models"
)

func TestService_FormatAccount_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	idPer := new(big.Int)
	idPer.SetString("170141183460469231731687303715884105754", 10)

	idNoPer := new(big.Int)
	idNoPer.SetString("170141183460469231731687303715884105753", 10)

	testCases := []struct {
		name            string
		id              *big.Int
		wantOwner       common.Address
		wantPermissions []*models.UserPermissions
		wantErr         error
	}{
		{
			name:      "no permissions",
			id:        idNoPer,
			wantOwner: common.HexToAddress("0x8382Be7cc5C2Cd8b14F44108444ced6745c5feCb"),
		},
		{
			name:      "one permission",
			id:        idNoPer,
			wantOwner: common.HexToAddress("0x8382Be7cc5C2Cd8b14F44108444ced6745c5feCb"),
			wantPermissions: []*models.UserPermissions{{
				User:        common.HexToAddress("0xE9B7BabCbE0a57b18ce9f6F6181193F65D80c999"),
				Permissions: []models.Permission{0},
			}},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewService(rpcClient, conf, coreC, perps)

			res, err := s.FormatAccount(tt.id)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantOwner, res.Owner)
			} else {
				require.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}

func TestService_FormatAccounts_OnChain_Limit(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	s, _ := NewService(rpcClient, conf, coreC, perps)

	_, err := s.FormatAccountsLimit(20000)

	require.NoError(t, err)
}
