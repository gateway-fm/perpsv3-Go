package events

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

type IEvents interface {
	//
}

type Events struct {
	rpcClient   *ethclient.Client
	core        *coreGoerli.CoreGoerli
	spotMarket  *spotMarketGoerli.SpotMarketGoerli
	perpsMarket *perpsMarketGoerli.PerpsMarketGoerli
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

type basicSubscription struct {
	stop                 chan struct{}
	contractSubscription event.Subscription
}

func newBasicSubscription(contractSubscription event.Subscription) *basicSubscription {
	return &basicSubscription{
		stop:                 make(chan struct{}),
		contractSubscription: contractSubscription,
	}
}

func (s *basicSubscription) Close() {
	s.contractSubscription.Unsubscribe()
	close(s.stop)
}

type TradeSubscription struct {
	*basicSubscription
	trades         chan *models.Trade
	contractTrades chan *perpsMarketGoerli.PerpsMarketGoerliOrderSettled
}

func newTradeSubscription(contractSubscription event.Subscription, contractTradesChannel chan *perpsMarketGoerli.PerpsMarketGoerliOrderSettled) *TradeSubscription {
	return &TradeSubscription{
		basicSubscription: newBasicSubscription(contractSubscription),
		contractTrades:    contractTradesChannel,
		trades:            make(chan *models.Trade),
	}
}

func (s *TradeSubscription) listen() {
	for {
		select {
		case <-s.stop:
			close(s.trades)
			close(s.contractTrades)
			return
		case err := <-s.contractSubscription.Err():
			logger.Log().WithField("layer", "Events-OrderSettled").Errorf("error listening order settled: %v", err.Error())
			// TODO should implement proper error handling
			continue
		case orderSettled := <-s.contractTrades:
			// TODO should fetch block number
			trade := models.GetTradeFromEvent(orderSettled, 0)

			s.trades <- trade
		}
	}
}

func (e *Events) ListenTrades() (*TradeSubscription, error) {
	contractTrades := make(chan *perpsMarketGoerli.PerpsMarketGoerliOrderSettled)

	contractSub, err := e.perpsMarket.WatchOrderSettled(nil, contractTrades, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-OrderSettled").Errorf("error watch order settled: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "OrderSettled")
	}

	tradesSub := newTradeSubscription(contractSub, contractTrades)

	return tradesSub, nil
}
