package models

import (
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

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
