package models

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

func TestGetLiquidationFromEvent(t *testing.T) {
	timeNow := time.Now()

	testCases := []struct {
		name  string
		event *perpsMarket.PerpsMarketPositionLiquidated
		time  uint64
		want  *Liquidation
	}{
		{
			name: "nil event",
			want: &Liquidation{},
		},
		{
			name: "only market ID",
			event: &perpsMarket.PerpsMarketPositionLiquidated{
				MarketId: big.NewInt(1),
			},
			want: &Liquidation{
				MarketID: uint64(1),
			},
		},
		{
			name: "only account ID",
			event: &perpsMarket.PerpsMarketPositionLiquidated{
				AccountId: big.NewInt(1),
			},
			want: &Liquidation{
				AccountID: big.NewInt(1),
			},
		},
		{
			name: "full event",
			event: &perpsMarket.PerpsMarketPositionLiquidated{
				MarketId:            big.NewInt(1),
				AccountId:           big.NewInt(2),
				AmountLiquidated:    big.NewInt(3),
				CurrentPositionSize: big.NewInt(4),
				Raw: types.Log{
					BlockNumber: 5,
				},
			},
			time: uint64(timeNow.Unix()),
			want: &Liquidation{
				MarketID:            uint64(1),
				AccountID:           big.NewInt(2),
				AmountLiquidated:    big.NewInt(3),
				CurrentPositionSize: big.NewInt(4),
				BlockNumber:         5,
				BlockTimestamp:      uint64(timeNow.Unix()),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := GetLiquidationFromEvent(tt.event, tt.time)

			require.Equal(t, tt.want, res)
		})
	}
}
