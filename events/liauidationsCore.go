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

// LiquidationsCoreSubscription is a struct for listening to all 'LiquidationsCore' contract events and return them as models.Account struct
type LiquidationsCoreSubscription struct {
	*basicSubscription
	NewAccountChan    chan *models.CoreLiquidation
	contractEventChan chan *core.CoreLiquidation
}

func (e *Events) ListenLiquidationsCore() (*LiquidationsCoreSubscription, error) {
	liquidationChan := make(chan *core.CoreLiquidation)

	liquidationSub, err := e.core.WatchLiquidation(nil, liquidationChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-LiquidationsCore").Errorf("error watch liquidation: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "LiquidationsCore")
	}

	liquidationsSub := newLiquidationsCoreSubscription(liquidationSub, liquidationChan)

	go liquidationsSub.listen(e.rpcClient)

	return liquidationsSub, nil
}

// newLiquidationsCoreSubscription is used to get new LiquidationsCoreSubscription instance
func newLiquidationsCoreSubscription(
	eventSub event.Subscription,
	liquidation chan *core.CoreLiquidation,
) *LiquidationsCoreSubscription {
	return &LiquidationsCoreSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		NewAccountChan:    make(chan *models.CoreLiquidation),
		contractEventChan: liquidation,
	}
}

// listen is used to run events listen goroutine
func (s *LiquidationsCoreSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.NewAccountChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-LiquidationsCore").Errorf(
					"error listening liquidation info: %v", err.Error(),
				)
				s.ErrChan <- err
			}
			return
		case liquidation := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(liquidation.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-LiquidationsCore").Warningf(
					"error fetching block number %v: %v; liquidation event time set to 0 ",
					liquidation.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			s.NewAccountChan <- models.GetCoreLiquidationFromEvent(liquidation, time)
		}
	}
}
