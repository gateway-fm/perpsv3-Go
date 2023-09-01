package services

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/stretchr/testify/require"
	"log"
	"math/big"
	"os"
	"testing"
)

func TestService_GetAccount_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	perps, _ := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

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
			s := NewService(rpcClient, coreC, 11664658, spot, 10875051, perps, 12708889)

			res, err := s.GetAccount(tt.id)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.wantOwner, res.Owner)
			} else {
				require.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}

func TestService_GetAccounts_OnChain_Limit(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	perps, _ := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	s := NewService(rpcClient, coreC, 11664658, spot, 10875051, perps, 12708889)

	_, err := s.GetAccountsLimit(20000)

	require.NoError(t, err)
}
