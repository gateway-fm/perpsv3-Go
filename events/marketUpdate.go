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

// MarketUpdateSubscription is a struct for listening to all 'MarketUpdated' contract events and return them as models.MarketUpdate struct
type MarketUpdateSubscription struct {
	*basicSubscription
	MarketUpdatesChan chan *models.MarketUpdate
	contractEventChan chan *perpsMarketGoerli.PerpsMarketGoerliMarketUpdated
}

// MarketUpdateSubscriptionBig is a struct for listening to all 'MarketUpdated' contract events and return them as models.MarketUpdateBig struct
type MarketUpdateSubscriptionBig struct {
	*basicSubscription
	MarketUpdatesChan chan *models.MarketUpdateBig
	contractEventChan chan *perpsMarketGoerli.PerpsMarketGoerliMarketUpdated
}

func (e *Events) ListenMarketUpdatesBig() (*MarketUpdateSubscriptionBig, error) {
	contractEventChan := make(chan *perpsMarketGoerli.PerpsMarketGoerliMarketUpdated)

	contractSub, err := e.perpsMarket.WatchMarketUpdated(nil, contractEventChan)
	if err != nil {
		logger.Log().WithField("layer", "Events-MarketUpdated").Errorf("error watch order committed: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "MarketUpdated")
	}

	marketUpdateSub := newMarketUpdateSubscriptionBig(contractSub, contractEventChan)

	go marketUpdateSub.listen(e.rpcClient)

	return marketUpdateSub, nil
}

func (e *Events) ListenMarketUpdates() (*MarketUpdateSubscription, error) {
	contractEventChan := make(chan *perpsMarketGoerli.PerpsMarketGoerliMarketUpdated)

	contractSub, err := e.perpsMarket.WatchMarketUpdated(nil, contractEventChan)
	if err != nil {
		logger.Log().WithField("layer", "Events-MarketUpdated").Errorf("error watch order committed: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "MarketUpdated")
	}

	marketUpdateSub := newMarketUpdateSubscription(contractSub, contractEventChan)

	go marketUpdateSub.listen(e.rpcClient)

	return marketUpdateSub, nil
}

// newMarketUpdateSubscription is used to create new MarketUpdateSubscription instance
func newMarketUpdateSubscription(eventSub event.Subscription, contractEventChan chan *perpsMarketGoerli.PerpsMarketGoerliMarketUpdated) *MarketUpdateSubscription {
	return &MarketUpdateSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		contractEventChan: contractEventChan,
		MarketUpdatesChan: make(chan *models.MarketUpdate),
	}
}

// newMarketUpdateSubscription is used to create new MarketUpdateSubscription instance
func newMarketUpdateSubscriptionBig(eventSub event.Subscription, contractEventChan chan *perpsMarketGoerli.PerpsMarketGoerliMarketUpdated) *MarketUpdateSubscriptionBig {
	return &MarketUpdateSubscriptionBig{
		basicSubscription: newBasicSubscription(eventSub),
		contractEventChan: contractEventChan,
		MarketUpdatesChan: make(chan *models.MarketUpdateBig),
	}
}

// listen is used to run a goroutine
func (s *MarketUpdateSubscriptionBig) listen(rpcClient *ethclient.Client) {
	for {
		select {
		case <-s.stop:
			close(s.MarketUpdatesChan)
			close(s.contractEventChan)
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-MarketUpdated").Errorf("error listening market update: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case marketUpdate := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(marketUpdate.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-MarketUpdated").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					marketUpdate.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			update := models.GetMarketUpdateBigFromEvent(marketUpdate, time)

			s.MarketUpdatesChan <- update
		}
	}
}

// listen is used to run a goroutine
func (s *MarketUpdateSubscription) listen(rpcClient *ethclient.Client) {
	for {
		select {
		case <-s.stop:
			close(s.MarketUpdatesChan)
			close(s.contractEventChan)
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-MarketUpdated").Errorf("error listening market update: %v", err.Error())
				s.ErrChan <- err
			}
			continue
		case marketUpdate := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(marketUpdate.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-MarketUpdated").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					marketUpdate.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			update := models.GetMarketUpdateFromEvent(marketUpdate, time)

			s.MarketUpdatesChan <- update
		}
	}
}
