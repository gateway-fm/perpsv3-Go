package events

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	perps_test "github.com/gateway-fm/perpsv3-Go/utils/testing-contracts/perps-test"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
	"time"
)

func TestEvents_ListenMarketUpdates_OnChain(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	rpc := os.Getenv("TEST_RPC_EVENTS")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := core.NewCore(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	perps, _ := perpsMarket.NewPerpsMarket(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	e := NewEvents(rpcClient, coreC, perps)

	subs, err := e.ListenMarketUpdates()
	require.NoError(t, err)

	stopChan := make(chan struct{})

	perpsTest := perps_test.GetTestPerpsMarket(
		rpc,
		"0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b",
		"0xe487Ad4291019b33e2230F8E2FB1fb6490325260",
		420,
	)
	defer perpsTest.Close()

	go func() {
		for {
			select {
			case <-stopChan:
				subs.Close()
				return
			case err = <-subs.ErrChan:
				require.NoError(t, err)
			case marketUpdate := <-subs.MarketUpdatesChan:
				log.Printf("marketUpdate received, block: %v", marketUpdate.BlockNumber)
				require.NotNil(t, marketUpdate)
			}
		}
	}()

	perpsTest.CommitOrder("100", "10000000000000000")
	time.Sleep(40 * time.Second)

	close(stopChan)
}
