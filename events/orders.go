package events

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// OrderSubscription is a struct for listening to all 'OrderCommitted' contract events and return them as models.Order struct
type OrderSubscription struct {
	*basicSubscription
	OrdersChan        chan *models.Order
	contractEventChan chan *perpsMarket.PerpsMarketOrderCommitted
}

func (e *Events) ListenOrders() (*OrderSubscription, error) {
	contractEventChan := make(chan *perpsMarket.PerpsMarketOrderCommitted)

	contractSub, err := e.perpsMarket.WatchOrderCommitted(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-OrderCommitted").Errorf("error watch order committed: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "OrderCommitted")
	}

	orderSub := newOrderSubscription(contractSub, contractEventChan)

	go orderSub.listen(e.rpcClient)

	return orderSub, nil
}

// newOrderSubscription is used to create new OrderSubscription instance
func newOrderSubscription(eventSub event.Subscription, contractEventChan chan *perpsMarket.PerpsMarketOrderCommitted) *OrderSubscription {
	return &OrderSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		contractEventChan: contractEventChan,
		OrdersChan:        make(chan *models.Order),
	}
}

// listen is used to run a goroutine
func (s *OrderSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.OrdersChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-OrderCommitted").Errorf("error listening order committed: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case orderCommitted := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(orderCommitted.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-OrderCommitted").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					orderCommitted.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			order := models.GetOrderFromEvent(orderCommitted, time)

			s.OrdersChan <- order
		}
	}
}
