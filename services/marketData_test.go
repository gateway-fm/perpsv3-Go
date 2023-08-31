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
	"os"
	"testing"
)

func TestService_RetrieveMarketUpdates_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	perps, _ := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	want := &models.MarketUpdate{
		MarketID:               200,
		Price:                  3780527432113118208,
		Skew:                   527000000000000000,
		Size:                   1049000000000000000,
		SizeDelta:              -255300000000000000,
		CurrentFundingRate:     62031943958317,
		CurrentFundingVelocity: 47430000000000,
		BlockNumber:            13739029,
		BlockTimestamp:         1692906126,
		TransactionHash:        "0x16704162005c11d71c745f7392a71a5ede8eb5f042e7fa917f210748773c57bf",
	}

	testCases := []struct {
		name       string
		startBlock uint64
		endBlock   uint64
		wantLength int
		wantAssert *models.MarketUpdate
	}{
		{
			name:       "one event",
			startBlock: uint64(13739029),
			endBlock:   uint64(13739029),
			wantLength: 1,
			wantAssert: want,
		},
		{
			name:       "five events",
			startBlock: uint64(13672986),
			endBlock:   uint64(13690512),
			wantLength: 5,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(rpcClient, coreC, 11664658, spot, 10875051, perps, 0)
			res, err := s.RetrieveMarketUpdates(tt.startBlock, &tt.endBlock)

			require.NoError(t, err)
			require.Equal(t, tt.wantLength, len(res))

			for i := 0; i < tt.wantLength; i++ {
				require.NotEqual(t, &models.MarketUpdate{}, res[i])
			}

			if tt.wantAssert != nil {
				require.Equal(t, tt.wantAssert, res[0])
			}
		})
	}
}
