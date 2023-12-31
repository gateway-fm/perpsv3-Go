package models

import (
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
)

func TestGetMarketUpdateFromEvent(t *testing.T) {
	timeNow := time.Now()

	testCases := []struct {
		name  string
		event *perpsMarket.PerpsMarketMarketUpdated
		time  uint64
		want  *MarketUpdate
	}{
		{
			name: "nil event",
			want: &MarketUpdate{},
		},
		{
			name: "only market ID",
			event: &perpsMarket.PerpsMarketMarketUpdated{
				MarketId: big.NewInt(1),
			},
			want: &MarketUpdate{
				MarketID:        uint64(1),
				TransactionHash: common.BytesToHash([]byte("")).Hex(),
			},
		},
		{
			name: "only price",
			event: &perpsMarket.PerpsMarketMarketUpdated{
				Price: big.NewInt(1),
			},
			want: &MarketUpdate{
				Price:           uint64(1),
				TransactionHash: common.BytesToHash([]byte("")).Hex(),
			},
		},
		{
			name: "only skew",
			event: &perpsMarket.PerpsMarketMarketUpdated{
				Skew: big.NewInt(1),
			},
			want: &MarketUpdate{
				Skew:            int64(1),
				TransactionHash: common.BytesToHash([]byte("")).Hex(),
			},
		},
		{
			name: "only size",
			event: &perpsMarket.PerpsMarketMarketUpdated{
				Size: big.NewInt(1),
			},
			want: &MarketUpdate{
				Size:            uint64(1),
				TransactionHash: common.BytesToHash([]byte("")).Hex(),
			},
		},
		{
			name: "only size delta",
			event: &perpsMarket.PerpsMarketMarketUpdated{
				SizeDelta: big.NewInt(1),
			},
			want: &MarketUpdate{
				SizeDelta:       int64(1),
				TransactionHash: common.BytesToHash([]byte("")).Hex(),
			},
		},
		{
			name: "only size delta",
			event: &perpsMarket.PerpsMarketMarketUpdated{
				SizeDelta: big.NewInt(1),
			},
			want: &MarketUpdate{
				SizeDelta:       int64(1),
				TransactionHash: common.BytesToHash([]byte("")).Hex(),
			},
		},
		{
			name: "only current funding rate",
			event: &perpsMarket.PerpsMarketMarketUpdated{
				CurrentFundingRate: big.NewInt(1),
			},
			want: &MarketUpdate{
				CurrentFundingRate: int64(1),
				TransactionHash:    common.BytesToHash([]byte("")).Hex(),
			},
		},
		{
			name: "only current funding velocity",
			event: &perpsMarket.PerpsMarketMarketUpdated{
				CurrentFundingVelocity: big.NewInt(1),
			},
			want: &MarketUpdate{
				CurrentFundingVelocity: int64(1),
				TransactionHash:        common.BytesToHash([]byte("")).Hex(),
			},
		},

		{
			name: "full event",
			event: &perpsMarket.PerpsMarketMarketUpdated{
				MarketId:               big.NewInt(1),
				Price:                  big.NewInt(2),
				Skew:                   big.NewInt(3),
				Size:                   big.NewInt(4),
				SizeDelta:              big.NewInt(5),
				CurrentFundingRate:     big.NewInt(6),
				CurrentFundingVelocity: big.NewInt(7),
				Raw: types.Log{
					BlockNumber: 8,
					TxHash:      common.BytesToHash([]byte("tx hash")),
				},
			},
			time: uint64(timeNow.Unix()),
			want: &MarketUpdate{
				MarketID:               uint64(1),
				Price:                  uint64(2),
				Skew:                   int64(3),
				Size:                   uint64(4),
				SizeDelta:              int64(5),
				CurrentFundingRate:     int64(6),
				CurrentFundingVelocity: int64(7),
				BlockNumber:            8,
				TransactionHash:        common.BytesToHash([]byte("tx hash")).Hex(),
				BlockTimestamp:         uint64(timeNow.Unix()),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := GetMarketUpdateFromEvent(tt.event, tt.time)

			require.Equal(t, tt.want, res)
		})
	}
}

func TestGetMarketUpdateBigFromEvent(t *testing.T) {
	timeNow := time.Now()

	testCases := []struct {
		name  string
		event *perpsMarket.PerpsMarketMarketUpdated
		time  uint64
		want  *MarketUpdateBig
	}{
		{
			name: "nil event",
			want: &MarketUpdateBig{},
		},
		{
			name: "full event",
			event: &perpsMarket.PerpsMarketMarketUpdated{
				MarketId:               big.NewInt(1),
				Price:                  big.NewInt(2),
				Skew:                   big.NewInt(3),
				Size:                   big.NewInt(4),
				SizeDelta:              big.NewInt(5),
				CurrentFundingRate:     big.NewInt(6),
				CurrentFundingVelocity: big.NewInt(7),
				Raw: types.Log{
					BlockNumber: 8,
					TxHash:      common.BytesToHash([]byte("tx hash")),
				},
			},
			time: uint64(timeNow.Unix()),
			want: &MarketUpdateBig{
				MarketID:               big.NewInt(1),
				Price:                  big.NewInt(2),
				Skew:                   big.NewInt(3),
				Size:                   big.NewInt(4),
				SizeDelta:              big.NewInt(5),
				CurrentFundingRate:     big.NewInt(6),
				CurrentFundingVelocity: big.NewInt(7),
				BlockNumber:            8,
				TransactionHash:        common.BytesToHash([]byte("tx hash")).Hex(),
				BlockTimestamp:         uint64(timeNow.Unix()),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := GetMarketUpdateBigFromEvent(tt.event, tt.time)

			require.Equal(t, tt.want, res)
		})
	}
}

func TestGetMarketMetadataFromContractResponse(t *testing.T) {
	testCases := []struct {
		name     string
		marketID *big.Int
		nameM    string
		symbol   string
		want     *MarketMetadata
	}{
		{
			marketID: big.NewInt(100),
			nameM:    "Ethereum",
			symbol:   "ETH",
			want: &MarketMetadata{
				MarketID: big.NewInt(100),
				Name:     "Ethereum",
				Symbol:   "ETH",
			},
		},
		{
			marketID: big.NewInt(200),
			nameM:    "Bitcoin",
			symbol:   "BTC",
			want: &MarketMetadata{
				MarketID: big.NewInt(200),
				Name:     "Bitcoin",
				Symbol:   "BTC",
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := GetMarketMetadataFromContractResponse(tt.marketID, tt.nameM, tt.symbol)

			require.Equal(t, tt.want, res)
		})
	}
}

func TestGetMarketSummaryFromContractModel(t *testing.T) {
	testCases := []struct {
		name     string
		summary  perpsMarket.IPerpsMarketModuleMarketSummary
		marketID *big.Int
		time     uint64
		want     *MarketSummary
	}{
		{
			name: "nil object",
			want: &MarketSummary{},
		},
		{
			name: "full object",
			summary: perpsMarket.IPerpsMarketModuleMarketSummary{
				Skew:                   big.NewInt(100),
				Size:                   big.NewInt(200),
				MaxOpenInterest:        big.NewInt(300),
				CurrentFundingRate:     big.NewInt(400),
				CurrentFundingVelocity: big.NewInt(500),
				IndexPrice:             big.NewInt(600),
			},
			marketID: big.NewInt(10),
			time:     uint64(1000),
			want: &MarketSummary{
				MarketID:               big.NewInt(10),
				Skew:                   big.NewInt(100),
				Size:                   big.NewInt(200),
				MaxOpenInterest:        big.NewInt(300),
				CurrentFundingRate:     big.NewInt(400),
				CurrentFundingVelocity: big.NewInt(500),
				IndexPrice:             big.NewInt(600),
				BlockTimestamp:         uint64(1000),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := GetMarketSummaryFromContractModel(tt.summary, tt.marketID, tt.time)

			require.Equal(t, tt.want, res)
		})
	}
}
