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

// PoolCreatedSubscription is a struct for listening to all 'PoolCreated' contract events and return them as models.PoolCreated struct
type PoolCreatedSubscription struct {
	*basicSubscription
	PoolCreatedChan   chan *models.PoolCreated
	contractEventChan chan *core.CorePoolCreated
}

func (e *Events) ListenPoolCreated() (*PoolCreatedSubscription, error) {
	contractEventChan := make(chan *core.CorePoolCreated)

	contractSub, err := e.core.WatchPoolCreated(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenPoolCreated").Errorf("error watch pool created: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "PoolCreated")
	}

	rewardSub := newPoolCreatedSubscription(contractSub, contractEventChan)

	go rewardSub.listen(e.rpcClient)

	return rewardSub, nil
}

// newPoolCreatedSubscription is used to create new PoolCreatedSubscription instance
func newPoolCreatedSubscription(eventSub event.Subscription, contractEventChan chan *core.CorePoolCreated) *PoolCreatedSubscription {
	return &PoolCreatedSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		contractEventChan: contractEventChan,
		PoolCreatedChan:   make(chan *models.PoolCreated),
	}
}

// listen is used to run a goroutine
func (s *PoolCreatedSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.PoolCreatedChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-PoolCreated").Errorf("error listening pool created: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventPoolCreated := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventPoolCreated.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-PoolCreated").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					eventPoolCreated.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			created := models.GetPoolCreatedFromEvent(eventPoolCreated, time)

			s.PoolCreatedChan <- created
		}
	}
}
