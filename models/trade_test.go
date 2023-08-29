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

func TestGetTradeFromEvent(t *testing.T) {
	timeNow := time.Now()

	testCases := []struct {
		name  string
		event *perpsMarketGoerli.PerpsMarketGoerliOrderSettled
		time  uint64
		want  *Trade
	}{
		{
			name: "nil event",
			want: &Trade{},
		},
		{
			name: "only market ID",
			event: &perpsMarketGoerli.PerpsMarketGoerliOrderSettled{
				MarketId: big.NewInt(1),
			},
			want: &Trade{
				MarketID:        uint64(1),
				TransactionHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
			},
		},
		{
			name: "only account ID",
			event: &perpsMarketGoerli.PerpsMarketGoerliOrderSettled{
				AccountId: big.NewInt(1),
			},
			want: &Trade{
				AccountID:       uint64(1),
				TransactionHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
			},
		},
		{
			name: "full event",
			event: &perpsMarketGoerli.PerpsMarketGoerliOrderSettled{
				MarketId:         big.NewInt(1),
				AccountId:        big.NewInt(2),
				FillPrice:        big.NewInt(3),
				AccruedFunding:   big.NewInt(4),
				SizeDelta:        big.NewInt(5),
				NewSize:          big.NewInt(6),
				TotalFees:        big.NewInt(7),
				ReferralFees:     big.NewInt(8),
				CollectedFees:    big.NewInt(9),
				SettlementReward: big.NewInt(10),
				TrackingCode:     crypto.Keccak256Hash([]byte("tracking_code")),
				Settler:          common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				Raw: types.Log{
					BlockNumber: 11,
					TxHash:      crypto.Keccak256Hash([]byte("tx_hash")),
				},
			},
			time: uint64(timeNow.Unix()),
			want: &Trade{
				MarketID:         1,
				AccountID:        2,
				FillPrice:        big.NewInt(3),
				AccruedFunding:   big.NewInt(4),
				SizeDelta:        big.NewInt(5),
				NewSize:          big.NewInt(6),
				TotalFees:        big.NewInt(7),
				ReferralFees:     big.NewInt(8),
				CollectedFees:    big.NewInt(9),
				SettlementReward: big.NewInt(10),
				TrackingCode:     crypto.Keccak256Hash([]byte("tracking_code")),
				Settler:          common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"),
				BlockNumber:      11,
				TransactionHash:  crypto.Keccak256Hash([]byte("tx_hash")).Hex(),
				BlockTimestamp:   uint64(timeNow.Unix()),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := GetTradeFromEvent(tt.event, tt.time)

			require.Equal(t, tt.want, res)
		})
	}
}
