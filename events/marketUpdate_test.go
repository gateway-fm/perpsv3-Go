package events

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	perps_test "github.com/gateway-fm/perpsv3-Go/utils/tesing-contracts/perps-test"
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

	rpc := os.Getenv("TEST_RPC_EVENTS")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	perps, _ := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	e := NewEvents(rpcClient, coreC, spot, perps)

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
