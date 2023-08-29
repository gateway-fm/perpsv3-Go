package models

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/stretchr/testify/require"
)

func TestGetOrderFromEvent(t *testing.T) {
	timeNow := time.Now()

	testCases := []struct {
		name  string
		event *perpsMarketGoerli.PerpsMarketGoerliOrderCommitted
		time  uint64
		want  *Order
	}{
		{
			name: "nil event",
			want: &Order{},
		},
		{
			name: "only market ID",
			event: &perpsMarketGoerli.PerpsMarketGoerliOrderCommitted{
				MarketId: big.NewInt(1),
			},
			want: &Order{
				MarketID: uint64(1),
			},
		},
		{
			name: "only account ID",
			event: &perpsMarketGoerli.PerpsMarketGoerliOrderCommitted{
				AccountId: big.NewInt(1),
			},
			want: &Order{
				AccountID: uint64(1),
			},
		},
		{
			name: "only settlement time",
			event: &perpsMarketGoerli.PerpsMarketGoerliOrderCommitted{
				SettlementTime: big.NewInt(1),
			},
			want: &Order{
				SettlementTime: uint64(1),
			},
		},
		{
			name: "only expiration time",
			event: &perpsMarketGoerli.PerpsMarketGoerliOrderCommitted{
				ExpirationTime: big.NewInt(1),
			},
			want: &Order{
				ExpirationTime: uint64(1),
			},
		},
		{
			name: "full event",
			event: &perpsMarketGoerli.PerpsMarketGoerliOrderCommitted{
				MarketId:        big.NewInt(1),
				AccountId:       big.NewInt(2),
				OrderType:       uint8(3),
				SizeDelta:       big.NewInt(5),
				AcceptablePrice: big.NewInt(6),
				SettlementTime:  big.NewInt(7),
				ExpirationTime:  big.NewInt(8),
				TrackingCode:    crypto.Keccak256Hash([]byte("tracking_code")),
				Sender:          common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				Raw: types.Log{
					BlockNumber: 9,
				},
			},
			time: uint64(timeNow.Unix()),
			want: &Order{
				MarketID:        uint64(1),
				AccountID:       uint64(2),
				OrderType:       uint8(3),
				SizeDelta:       big.NewInt(5),
				AcceptablePrice: big.NewInt(6),
				SettlementTime:  uint64(7),
				ExpirationTime:  uint64(8),
				TrackingCode:    crypto.Keccak256Hash([]byte("tracking_code")),
				Sender:          common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				BlockNumber:     9,
				BlockTimestamp:  uint64(timeNow.Unix()),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := GetOrderFromEvent(tt.event, tt.time)

			require.Equal(t, tt.want, res)
		})
	}
}
