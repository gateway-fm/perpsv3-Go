package services

import (
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/stretchr/testify/require"
)

func TestService_RetrieveTrades_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	perps, _ := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	fillPrice := new(big.Int)
	fillPrice.SetString("26050753699652653732215", 10)

	pnl := new(big.Int)
	pnl.SetString("36901526133944874", 10)

	accruedF := new(big.Int)
	accruedF.SetString("-3064923447473062", 10)

	sizeD := new(big.Int)
	sizeD.SetString("-255300000000000000", 10)

	sizeN := new(big.Int)
	sizeN.SetString("-261000000000000000", 10)

	feesT := new(big.Int)
	feesT.SetString("4655530193664925748", 10)

	want := &models.Trade{
		MarketID:         200,
		AccountID:        1692895537802,
		FillPrice:        fillPrice,
		PnL:              pnl,
		AccruedFunding:   accruedF,
		SizeDelta:        sizeD,
		NewSize:          sizeN,
		TotalFees:        feesT,
		ReferralFees:     big.NewInt(0),
		CollectedFees:    big.NewInt(0),
		SettlementReward: big.NewInt(0),
		TrackingCode:     [32]byte{75, 87, 69, 78, 84, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Settler:          common.HexToAddress("0x4A58e0d29558111bfDc07Dc12Ca0fF7fcD0d0d75"),
		BlockNumber:      13739029,
		BlockTimestamp:   1692906126,
		TransactionHash:  "0x16704162005c11d71c745f7392a71a5ede8eb5f042e7fa917f210748773c57bf",
	}

	testCases := []struct {
		name       string
		startBlock uint64
		endBlock   uint64
		wantLength int
		wantAssert *models.Trade
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
			res, err := s.RetrieveTrades(tt.startBlock, &tt.endBlock)

			require.NoError(t, err)
			require.Equal(t, tt.wantLength, len(res))

			for i := 0; i < tt.wantLength; i++ {
				require.NotEqual(t, &models.Trade{}, res[i])
			}

			if tt.wantAssert != nil {
				require.Equal(t, tt.wantAssert.MarketID, res[0].MarketID)
				require.Equal(t, tt.wantAssert.AccountID, res[0].AccountID)
				require.Equal(t, tt.wantAssert.FillPrice, res[0].FillPrice)
				require.Equal(t, tt.wantAssert.PnL, res[0].PnL)
				require.Equal(t, tt.wantAssert.AccruedFunding, res[0].AccruedFunding)
				require.Equal(t, tt.wantAssert.SizeDelta, res[0].SizeDelta)
				require.Equal(t, tt.wantAssert.NewSize, res[0].NewSize)
				require.Equal(t, tt.wantAssert.TotalFees, res[0].TotalFees)
				require.Equal(t, tt.wantAssert.ReferralFees.Bytes(), res[0].ReferralFees.Bytes())
				require.Equal(t, tt.wantAssert.CollectedFees.Bytes(), res[0].CollectedFees.Bytes())
				require.Equal(t, tt.wantAssert.SettlementReward.Bytes(), res[0].SettlementReward.Bytes())
				require.Equal(t, tt.wantAssert.TrackingCode, res[0].TrackingCode)
				require.Equal(t, tt.wantAssert.Settler, res[0].Settler)
				require.Equal(t, tt.wantAssert.BlockNumber, res[0].BlockNumber)
				require.Equal(t, tt.wantAssert.BlockTimestamp, res[0].BlockTimestamp)
				require.Equal(t, tt.wantAssert.TransactionHash, res[0].TransactionHash)
			}
		})
	}
}
