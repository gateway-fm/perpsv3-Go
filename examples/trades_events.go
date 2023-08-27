package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	events2 "github.com/gateway-fm/perpsv3-Go/events"
)

func main() {
	rpcClient, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}

	coreC, err := coreGoerli.NewCoreGoerli(common.HexToAddress("0x76490713314fCEC173f44e99346F54c6e92a8E42"), rpcClient)
	if err != nil {
		log.Fatal(err)
	}
	spot, err := spotMarketGoerli.NewSpotMarketGoerli(common.HexToAddress("0x5FF4b3aacdeC86782d8c757FAa638d8790799E83"), rpcClient)
	if err != nil {
		log.Fatal(err)
	}
	perps, err := perpsMarketGoerli.NewPerpsMarketGoerli(common.HexToAddress("0xf272382cB3BE898A8CdB1A23BE056fA2Fcf4513b"), rpcClient)
	if err != nil {
		log.Fatal(err)
	}

	eventsSer := events2.NewEvents(rpcClient, coreC, spot, perps)

	subs, err := eventsSer.ListenTrades()
	if err != nil {
		log.Fatal(err)
	}

	stopChan := make(chan struct{})

	// handle events
	go func() {
		for {
			select {
			case <-stopChan:
				subs.Close()
				return
			case err = <-subs.ErrChan:
				log.Println(err.Error())
			case trade := <-subs.TradesChan:
				fmt.Println(trade.AccountID)
				fmt.Println(trade.AccruedFunding)
			}
		}
	}()

	time.Sleep(10 * time.Second)

	// stop listening
	close(stopChan)

	time.Sleep(5 * time.Second)

}
