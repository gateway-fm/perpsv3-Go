package main

import (
	"fmt"
	"log"
	"time"

	perpsv3_Go "github.com/gateway-fm/perpsv3-Go"
)

func main() {
	conf := perpsv3_Go.GetOptimismGoerliDefaultConfig("")

	lib, err := perpsv3_Go.Create(conf)
	if err != nil {
		log.Fatal(err)
	}

	subs, err := lib.ListenTrades()
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
