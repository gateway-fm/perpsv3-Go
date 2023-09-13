package events

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// TradeSubscription is a struct for listening to all 'OrderSettled' contract events and return them as models.Trade struct
type TradeSubscription struct {
	*basicSubscription
	TradesChan        chan *models.Trade
	contractEventChan chan *perpsMarketGoerli.PerpsMarketGoerliOrderSettled
}

func (e *Events) ListenTrades() (*TradeSubscription, error) {
	contractEventChan := make(chan *perpsMarketGoerli.PerpsMarketGoerliOrderSettled)

	contractSub, err := e.perpsMarket.WatchOrderSettled(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-OrderSettled").Errorf("error watch order settled: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "OrderSettled")
	}

	tradesSub := newTradeSubscription(contractSub, contractEventChan)

	go tradesSub.listen(e.rpcClient)

	return tradesSub, nil
}

// newTradeSubscription is used to create new TradeSubscription instance
func newTradeSubscription(eventSub event.Subscription, contractEventChan chan *perpsMarketGoerli.PerpsMarketGoerliOrderSettled) *TradeSubscription {
	return &TradeSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		contractEventChan: contractEventChan,
		TradesChan:        make(chan *models.Trade),
	}
}

// listen is used to run a goroutine
func (s *TradeSubscription) listen(rpcClient *ethclient.Client) {
	for {
		select {
		case <-s.stop:
			close(s.TradesChan)
			close(s.contractEventChan)
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-OrderSettled").Errorf("error listening order settled: %v", err.Error())
				s.ErrChan <- err
			}
			continue
		case orderSettled := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(orderSettled.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-OrderSettled").Warningf(
					"error fetching block number %v: %v; trade event time set to 0 ",
					orderSettled.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			trade := models.GetTradeFromEvent(orderSettled, time)

			s.TradesChan <- trade
		}
	}
}
