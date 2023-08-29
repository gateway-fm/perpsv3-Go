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

func TestService_RetrieveOrders_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	perps, _ := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	sizeDelta := new(big.Int)
	sizeDelta.SetString("-500000000000000000", 10)

	acceptablePrice := new(big.Int)
	acceptablePrice.SetString("1000000000000000000", 10)

	want := &models.Order{
		MarketID:        100,
		AccountID:       8714,
		OrderType:       0,
		SizeDelta:       sizeDelta,
		AcceptablePrice: acceptablePrice,
		SettlementTime:  1693269991,
		ExpirationTime:  1693270591,
		TrackingCode:    [32]byte{},
		Sender:          common.HexToAddress("0xC47fF8a340dFc0605be060886F0B6AEea0db653f"),
		BlockNumber:     13920954,
		BlockTimestamp:  1693269976,
	}

	testCases := []struct {
		name       string
		startBlock uint64
		endBlock   uint64
		wantLength int
		wantAssert *models.Order
	}{
		{
			name:       "one event",
			startBlock: uint64(13920954),
			endBlock:   uint64(13920954),
			wantLength: 1,
			wantAssert: want,
		},
		{
			name:       "five events",
			startBlock: uint64(13907698),
			endBlock:   uint64(13920954),
			wantLength: 5,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(rpcClient, coreC, 11664658, spot, 10875051, perps, 0)
			res, err := s.RetrieveOrders(tt.startBlock, &tt.endBlock)

			if err != nil {
				log.Println("err:", err.Error())
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.wantLength, len(res))

			for i := 0; i < tt.wantLength; i++ {
				require.NotEqual(t, &models.Trade{}, res[i])
			}

			if tt.wantAssert != nil {
				require.Equal(t, tt.wantAssert, res[0])
			}
		})
	}
}
