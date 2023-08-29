package events

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/stretchr/testify/require"
)

func TestEvents_ListenOrders(t *testing.T) {
	rpc := os.Getenv("TEST_RPC_EVENTS")
	if rpc == "" {
		log.Fatal("no rpc in env vars")
	}

	rpcClient, _ := ethclient.Dial(rpc)

	coreC, _ := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	spot, _ := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	perps, _ := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)

	e := NewEvents(rpcClient, coreC, spot, perps)

	subs, err := e.ListenOrders()
	require.NoError(t, err)

	stopChan := make(chan struct{})

	// handle events
	go func() {
		for {
			select {
			case <-stopChan:
				subs.Close()
				return
			case err = <-subs.ErrChan:
				require.NoError(t, err)
			case trade := <-subs.OrdersChan:
				require.NotEqual(t, &models.Order{}, trade)
			}
		}
	}()

	time.Sleep(10 * time.Millisecond)

	// stop listening
	close(stopChan)

	time.Sleep(10 * time.Millisecond)
}
