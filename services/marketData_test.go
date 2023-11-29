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
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
)

func TestService_RetrieveMarketUpdates_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

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
			s, _ := NewService(rpcClient, conf, coreC, perps)
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

func TestService_RetrieveMarketUpdatesBig_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	price := new(big.Int)
	price.SetString("26050583159510000000000", 10)

	want := &models.MarketUpdateBig{
		MarketID:               big.NewInt(200),
		Price:                  price,
		Skew:                   big.NewInt(527000000000000000),
		Size:                   big.NewInt(1049000000000000000),
		SizeDelta:              big.NewInt(-255300000000000000),
		CurrentFundingRate:     big.NewInt(62031943958317),
		CurrentFundingVelocity: big.NewInt(47430000000000),
		BlockNumber:            13739029,
		BlockTimestamp:         1692906126,
		TransactionHash:        "0x16704162005c11d71c745f7392a71a5ede8eb5f042e7fa917f210748773c57bf",
	}

	testCases := []struct {
		name       string
		startBlock uint64
		endBlock   uint64
		wantLength int
		wantAssert *models.MarketUpdateBig
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
			s, _ := NewService(rpcClient, conf, coreC, perps)
			res, err := s.RetrieveMarketUpdatesBig(tt.startBlock, &tt.endBlock)

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

func TestService_RetrieveMarketUpdates_OnChain_Limit(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	s, _ := NewService(rpcClient, conf, coreC, perps)

	_, err := s.RetrieveMarketUpdatesLimit(20000)

	require.NoError(t, err)
}

func TestService_RetrieveMarketUpdatesBig_OnChain_Limit(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	s, _ := NewService(rpcClient, conf, coreC, perps)

	_, err := s.RetrieveMarketUpdatesBigLimit(20000)

	require.NoError(t, err)
}

func TestService_GetMarketMetadata_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

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
			s, _ := NewService(rpcClient, conf, coreC, perps)

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

func TestService_GetMarketSummary(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	testCases := []struct {
		name     string
		marketID *big.Int
		wantErr  error
	}{
		{
			name:    "nil id",
			wantErr: errors.InvalidArgumentErr,
		},
		{
			name:     "id 100",
			marketID: big.NewInt(100),
		},
		{
			name:     "id 200",
			marketID: big.NewInt(200),
		},
		{
			name:     "id 300",
			marketID: big.NewInt(300),
			wantErr:  errors.InvalidArgumentErr,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewService(rpcClient, conf, coreC, perps)

			res, err := s.GetMarketSummary(tt.marketID)

			if tt.wantErr == nil {
				require.Equal(t, tt.marketID, res.MarketID)
				require.NotNil(t, res.Size)
				require.NotNil(t, res.Skew)
				require.NotNil(t, res.MaxOpenInterest)
				require.NotNil(t, res.CurrentFundingVelocity)
				require.NotNil(t, res.CurrentFundingRate)
				require.NotNil(t, res.IndexPrice)
				require.NotEqual(t, 0, res.BlockTimestamp)
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}

func TestService_GetMarketIDs(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	testCases := []struct {
		name    string
		want    []*big.Int
		wantErr error
	}{
		{
			name: "get ids",
			want: []*big.Int{big.NewInt(100), big.NewInt(200)},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewService(rpcClient, conf, coreC, perps)

			res, err := s.GetMarketIDs()

			if tt.wantErr == nil {
				require.Equal(t, tt.want, res)
			} else {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}

func TestService_GetFoundingRate(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	testCases := []struct {
		name    string
		id      *big.Int
		wantErr error
	}{
		{
			name: "id 100",
			id:   big.NewInt(100),
		},
		{
			name: "id 200",
			id:   big.NewInt(200),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewService(rpcClient, conf, coreC, perps)

			res, err := s.GetFoundingRate(tt.id)

			log.Println(res)

			if tt.wantErr == nil {
				require.NotNil(t, res)
			} else {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}
