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

func TestService_RetrieveLiquidations_OnChain_Limit(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	perps, _ := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	s := NewService(rpcClient, coreC, 11664658, spot, 10875051, perps, 12708889)

	_, err := s.RetrieveLiquidationsLimit(20000)

	require.NoError(t, err)
}

func TestService_RetrieveLiquidations_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	perps, _ := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	id := new(big.Int)
	id.SetString("170141183460469231731687303715884105753", 10)

	want := &models.Liquidation{
		MarketID:            100,
		AccountID:           id,
		AmountLiquidated:    big.NewInt(200000000000000000),
		CurrentPositionSize: big.NewInt(0),
		BlockNumber:         14125002,
		BlockTimestamp:      1693678072,
	}

	testCases := []struct {
		name       string
		startBlock uint64
		endBlock   uint64
		wantLength int
		wantAssert *models.Liquidation
	}{
		{
			name:       "one event",
			startBlock: uint64(14125002),
			endBlock:   uint64(14125002),
			wantLength: 1,
			wantAssert: want,
		},
		{
			name:       "four events",
			startBlock: uint64(14001142),
			endBlock:   uint64(14001142),
			wantLength: 4,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(rpcClient, coreC, 11664658, spot, 10875051, perps, 0)
			res, err := s.RetrieveLiquidations(tt.startBlock, &tt.endBlock)

			if err != nil {
				log.Println("err:", err.Error())
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.wantLength, len(res))

			for i := 0; i < tt.wantLength; i++ {
				require.NotEqual(t, &models.Liquidation{}, res[i])
			}

			if tt.wantAssert != nil {
				require.Equal(t, tt.wantAssert.MarketID, res[0].MarketID)
				require.Equal(t, tt.wantAssert.AccountID, res[0].AccountID)
				require.Equal(t, tt.wantAssert.AmountLiquidated, res[0].AmountLiquidated)
				require.Equal(t, tt.wantAssert.CurrentPositionSize.Bytes(), res[0].CurrentPositionSize.Bytes())
				require.Equal(t, tt.wantAssert.BlockNumber, res[0].BlockNumber)
				require.Equal(t, tt.wantAssert.BlockTimestamp, res[0].BlockTimestamp)
			}
		})
	}
}
