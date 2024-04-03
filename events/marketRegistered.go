package events

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// MarketRegisteredSubscription is a struct for listening to all 'MarketRegistered' contract events and return them as models.MarketRegistered struct
type MarketRegisteredSubscription struct {
	*basicSubscription
	MarketRegisteredChan chan *models.MarketRegistered
	contractEventChan    chan *core.CoreMarketRegistered
}

func (e *Events) ListenMarketRegistered() (*MarketRegisteredSubscription, error) {
	contractEventChan := make(chan *core.CoreMarketRegistered)

	contractSub, err := e.core.WatchMarketRegistered(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenMarketRegistered").Errorf("error watch reward claimed: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "MarketRegistered")
	}

	rewardSub := newMarketRegisteredSubscription(contractSub, contractEventChan)

	go rewardSub.listen(e.rpcClient)

	return rewardSub, nil
}

// newMarketRegisteredSubscription is used to create new MarketRegisteredSubscription instance
func newMarketRegisteredSubscription(eventSub event.Subscription, contractEventChan chan *core.CoreMarketRegistered) *MarketRegisteredSubscription {
	return &MarketRegisteredSubscription{
		basicSubscription:    newBasicSubscription(eventSub),
		contractEventChan:    contractEventChan,
		MarketRegisteredChan: make(chan *models.MarketRegistered),
	}
}

// listen is used to run a goroutine
func (s *MarketRegisteredSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.MarketRegisteredChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-MarketRegistered").Errorf("error listening reward claimed: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventMarketRegistered := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventMarketRegistered.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-MarketRegistered").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					eventMarketRegistered.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			claimed := models.GetMarketRegisteredFromEvent(eventMarketRegistered, time)

			s.MarketRegisteredChan <- claimed
		}
	}
}
