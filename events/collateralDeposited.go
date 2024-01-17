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

// CollateralDepositedSubscription is a struct for listening to all 'Deposited' contract events and return them as models.CollateralDeposited struct
type CollateralDepositedSubscription struct {
	*basicSubscription
	CollateralDepositedChan chan *models.CollateralDeposited
	contractEventChan       chan *core.CoreDeposited
}

func (e *Events) ListenCollateralDeposited() (*CollateralDepositedSubscription, error) {
	contractEventChan := make(chan *core.CoreDeposited)

	contractSub, err := e.core.WatchDeposited(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenCollateralDeposited").Errorf("error watch withdrawn: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "CollateralDeposited")
	}

	depositSub := newCollateralDepositedSubscription(contractSub, contractEventChan)

	go depositSub.listen(e.rpcClient)

	return depositSub, nil
}

// newCollateralDepositedSubscription is used to create new CollateralDepositedSubscription instance
func newCollateralDepositedSubscription(eventSub event.Subscription, contractEventChan chan *core.CoreDeposited) *CollateralDepositedSubscription {
	return &CollateralDepositedSubscription{
		basicSubscription:       newBasicSubscription(eventSub),
		contractEventChan:       contractEventChan,
		CollateralDepositedChan: make(chan *models.CollateralDeposited),
	}
}

// listen is used to run a goroutine
func (s *CollateralDepositedSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.CollateralDepositedChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-CollateralDeposited").Errorf("error listening withdrawn: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventCollateralDeposited := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventCollateralDeposited.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-CollateralDeposited").Warningf(
					"error fetching block number %v: %v; deposit event time set to 0 ",
					eventCollateralDeposited.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			deposit := models.GetCollateralDepositedFromEvent(eventCollateralDeposited, time)

			s.CollateralDepositedChan <- deposit
		}
	}
}
