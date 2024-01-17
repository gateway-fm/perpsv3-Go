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

// CollateralWithdrawnSubscription is a struct for listening to all 'Withdrawn' contract events and return them as models.CollateralWithdrawn struct
type CollateralWithdrawnSubscription struct {
	*basicSubscription
	CollateralWithdrawnChan chan *models.CollateralWithdrawn
	contractEventChan       chan *core.CoreWithdrawn
}

func (e *Events) ListenCollateralWithdrawn() (*CollateralWithdrawnSubscription, error) {
	contractEventChan := make(chan *core.CoreWithdrawn)

	contractSub, err := e.core.WatchWithdrawn(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenCollateralWithdrawn").Errorf("error watch withdrawn: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "CollateralWithdrawn")
	}

	withdrawnSub := newCollateralWithdrawnSubscription(contractSub, contractEventChan)

	go withdrawnSub.listen(e.rpcClient)

	return withdrawnSub, nil
}

// newCollateralWithdrawnSubscription is used to create new CollateralWithdrawnSubscription instance
func newCollateralWithdrawnSubscription(eventSub event.Subscription, contractEventChan chan *core.CoreWithdrawn) *CollateralWithdrawnSubscription {
	return &CollateralWithdrawnSubscription{
		basicSubscription:       newBasicSubscription(eventSub),
		contractEventChan:       contractEventChan,
		CollateralWithdrawnChan: make(chan *models.CollateralWithdrawn),
	}
}

// listen is used to run a goroutine
func (s *CollateralWithdrawnSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.CollateralWithdrawnChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-CollateralWithdrawn").Errorf("error listening withdrawn: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventCollateralWithdrawn := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventCollateralWithdrawn.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-CollateralWithdrawn").Warningf(
					"error fetching block number %v: %v; withdraw event time set to 0 ",
					eventCollateralWithdrawn.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			withdraw := models.GetCollateralWithdrawnFromEvent(eventCollateralWithdrawn, time)

			s.CollateralWithdrawnChan <- withdraw
		}
	}
}
