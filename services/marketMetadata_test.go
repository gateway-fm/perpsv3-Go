package services

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/stretchr/testify/require"
	"log"
	"math/big"
	"os"
	"testing"
)

func TestService_GetMarketMetadata_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	perps, _ := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	testCases := []struct {
		name     string
		marketID *big.Int
		want     *models.MarketMetadata
		wantErr  error
	}{
		{
			name:     "id 200",
			marketID: big.NewInt(200),
			want: &models.MarketMetadata{
				MarketID: big.NewInt(200),
				Name:     "Bitcoin",
				Symbol:   "BTC",
			},
		},
		{
			name:     "id 100",
			marketID: big.NewInt(100),
			want: &models.MarketMetadata{
				MarketID: big.NewInt(100),
				Name:     "Ethereum",
				Symbol:   "ETH",
			},
		},
		{
			name:     "id 300",
			marketID: big.NewInt(300),
			wantErr:  errors.InvalidArgumentErr,
		},
		{
			name:    "nil id",
			wantErr: errors.InvalidArgumentErr,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(rpcClient, coreC, 11664658, spot, 10875051, perps, 0)

			res, err := s.GetMarketMetadata(tt.marketID)

			if tt.wantErr == nil {
				require.NoError(t, err)
				require.Equal(t, tt.want, res)
			} else {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}
