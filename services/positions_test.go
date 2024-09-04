package services

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/gateway-fm/perpsv3-Go/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"

	"github.com/gateway-fm/perpsv3-Go/contracts/Account"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
)

func TestService_GetPosition_OnChain(t *testing.T) {
	rpc := os.Getenv("TEST_RPC")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	conf := config.GetBaseAndromedaDefaultConfig(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)
	acc, _ := Account.NewAccount(common.HexToAddress("0x63f4Dd0434BEB5baeCD27F3778a909278d8cf5b8"), rpcClient)

	testCases := []struct {
		name      string
		accountID *big.Int
		marketID  *big.Int
		wantErr   error
	}{
		{
			name:      "good data",
			accountID: big.NewInt(28006),
			marketID:  big.NewInt(100),
		},
		{
			name:      "0 account ID no errors",
			accountID: big.NewInt(0),
			marketID:  big.NewInt(100),
		},
		{
			name:      "bad market id",
			accountID: big.NewInt(28006),
			marketID:  big.NewInt(0),
			wantErr:   errors.GetReadContractErr(fmt.Errorf("execution reverted"), "PerpsMarket", "getOpenPosition"),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewService(rpcClient, conf, coreC, perps, acc)

			res, err := s.GetPosition(tt.accountID, tt.marketID)

			if tt.wantErr == nil {
				require.NotNil(t, res)
				require.NotNil(t, res.PositionSize)
				require.NotNil(t, res.AccruedFunding)
				require.NotNil(t, res.TotalPnl)
				require.NotEqual(t, res.BlockTimestamp, 0)
				require.NotEqual(t, res.BlockNumber, 0)
			} else {
				require.EqualError(t, tt.wantErr, err.Error())
			}
		})
	}
}
