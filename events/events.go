package events

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

type IEvents interface {
	//
}

type Events struct {
	rpcClient  *ethclient.Client
	core       *coreGoerli.CoreGoerli
	spotMarket *spotMarketGoerli.SpotMarketGoerli
}

func NewEvents(client *ethclient.Client, core *coreGoerli.CoreGoerli, spotMarket *spotMarketGoerli.SpotMarketGoerli) IEvents {
	return &Events{
		rpcClient:  client,
		core:       core,
		spotMarket: spotMarket,
	}
}

func (e *Events) listenToOrderSettled() error {
	trades := make(chan *spotMarketGoerli.SpotMarketGoerliOrderSettled)

	sub, err := e.spotMarket.WatchOrderSettled(nil, trades, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-OrderSettled").Errorf("error watch order settled: %v", err.Error())
		return errors.GetEventListenErr(err, "OrderSettled")
	}

	for {
		select {
		case err = <-sub.Err():
			logger.Log().WithField("layer", "Events-OrderSettled").Errorf("error listening order settled: %v", err.Error())
			return errors.GetEventListenErr(err, "OrderSettled")
		case trade := <-trades:
			log.Println("new trade:", trade)
		}
	}
}
