package models

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetPositionFromContract(t *testing.T) {
	timeNow := time.Now()

	testCases := []struct {
		name   string
		obj    positionContract
		blockN uint64
		time   uint64
		want   *Position
	}{
		{
			name: "blank object",
			want: &Position{},
		},
		{
			name: "full object",
			obj: positionContract{
				TotalPnl:       big.NewInt(1),
				AccruedFunding: big.NewInt(2),
				PositionSize:   big.NewInt(3),
			},
			time:   uint64(timeNow.Unix()),
			blockN: uint64(4),
			want: &Position{
				TotalPnl:       big.NewInt(1),
				AccruedFunding: big.NewInt(2),
				PositionSize:   big.NewInt(3),
				BlockTimestamp: uint64(timeNow.Unix()),
				BlockNumber:    uint64(4),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			res := GetPositionFromContract(tt.obj, tt.blockN, tt.time)

			require.Equal(t, tt.want, res)
		})
	}
}
