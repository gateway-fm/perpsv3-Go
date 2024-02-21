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

// MarketUSDWithdrawnSubscription is a struct for listening to all 'MarketUSDWithdrawn' contract events and return them as models.MarketUSDWithdrawn struct
type MarketUSDWithdrawnSubscription struct {
	*basicSubscription
	MarketUSDWithdrawnChan chan *models.MarketUSDWithdrawn
	contractEventChan      chan *core.CoreMarketUsdWithdrawn
}

func (e *Events) ListenMarketUSDWithdrawn() (*MarketUSDWithdrawnSubscription, error) {
	contractEventChan := make(chan *core.CoreMarketUsdWithdrawn)

	contractSub, err := e.core.WatchMarketUsdWithdrawn(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenMarketUSDWithdrawn").Errorf("error watch withdraw: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "MarketUSDWithdrawn")
	}

	withdrawSub := newMarketUSDWithdrawnSubscription(contractSub, contractEventChan)

	go withdrawSub.listen(e.rpcClient)

	return withdrawSub, nil
}

// newMarketUSDWithdrawnSubscription is used to create new MarketUSDWithdrawnSubscription instance
func newMarketUSDWithdrawnSubscription(eventSub event.Subscription, contractEventChan chan *core.CoreMarketUsdWithdrawn) *MarketUSDWithdrawnSubscription {
	return &MarketUSDWithdrawnSubscription{
		basicSubscription:      newBasicSubscription(eventSub),
		contractEventChan:      contractEventChan,
		MarketUSDWithdrawnChan: make(chan *models.MarketUSDWithdrawn),
	}
}

// listen is used to run a goroutine
func (s *MarketUSDWithdrawnSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.MarketUSDWithdrawnChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-MarketUSDWithdrawn").Errorf("error listening withdraw: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventMarketUSDWithdrawn := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventMarketUSDWithdrawn.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-MarketUSDWithdrawn").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					eventMarketUSDWithdrawn.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			withdrawn := models.GetMarketUSDWithdrawnFromEvent(eventMarketUSDWithdrawn, time)

			s.MarketUSDWithdrawnChan <- withdrawn
		}
	}
}
