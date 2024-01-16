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

// DelegationUpdatedSubscription is a struct for listening to all 'DelegationUpdated' contract events and return them as models.DelegationUpdated struct
type DelegationUpdatedSubscription struct {
	*basicSubscription
	DelegationUpdatedChan chan *models.DelegationUpdated
	contractEventChan     chan *core.CoreDelegationUpdated
}

func (e *Events) ListenDelegationUpdated() (*DelegationUpdatedSubscription, error) {
	contractEventChan := make(chan *core.CoreDelegationUpdated)

	contractSub, err := e.core.WatchDelegationUpdated(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenDelegationUpdated").Errorf("error watch usd burned: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "DelegationUpdated")
	}

	delegationSub := newDelegationUpdatedSubscription(contractSub, contractEventChan)

	go delegationSub.listen(e.rpcClient)

	return delegationSub, nil
}

// newDelegationUpdatedSubscription is used to create new DelegationUpdatedSubscription instance
func newDelegationUpdatedSubscription(eventSub event.Subscription, contractEventChan chan *core.CoreDelegationUpdated) *DelegationUpdatedSubscription {
	return &DelegationUpdatedSubscription{
		basicSubscription:     newBasicSubscription(eventSub),
		contractEventChan:     contractEventChan,
		DelegationUpdatedChan: make(chan *models.DelegationUpdated),
	}
}

// listen is used to run a goroutine
func (s *DelegationUpdatedSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.DelegationUpdatedChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-DelegationUpdated").Errorf("error listening usd burned: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventDelegationUpdated := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventDelegationUpdated.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-DelegationUpdated").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					eventDelegationUpdated.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			delegation := models.GetDelegationUpdatedFromEvent(eventDelegationUpdated, time)

			s.DelegationUpdatedChan <- delegation
		}
	}
}
