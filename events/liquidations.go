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

// LiquidationSubscription is a struct for listening to all 'PositionLiquidated' contract events and return them as models.Liquidation struct
type LiquidationSubscription struct {
	*basicSubscription
	LiquidationsChan  chan *models.Liquidation
	contractEventChan chan *perpsMarketGoerli.PerpsMarketGoerliPositionLiquidated
}

func (e *Events) ListenLiquidations() (*LiquidationSubscription, error) {
	contractEventChan := make(chan *perpsMarketGoerli.PerpsMarketGoerliPositionLiquidated)

	contractSub, err := e.perpsMarket.WatchPositionLiquidated(nil, contractEventChan, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-PositionLiquidated").Errorf("error watch order committed: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "PositionLiquidated")
	}

	orderSub := newLiquidationSubscription(contractSub, contractEventChan)

	go orderSub.listen(e.rpcClient)

	return orderSub, nil
}

// newLiquidationSubscription is used to create new LiquidationSubscription instance
func newLiquidationSubscription(eventSub event.Subscription, contractEventChan chan *perpsMarketGoerli.PerpsMarketGoerliPositionLiquidated) *LiquidationSubscription {
	return &LiquidationSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		contractEventChan: contractEventChan,
		LiquidationsChan:  make(chan *models.Liquidation),
	}
}

// listen is used to run a goroutine
func (s *LiquidationSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.LiquidationsChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-PositionLiquidated").Errorf("error listening position liquidated: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case positionLiquidated := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(positionLiquidated.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-PositionLiquidated").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					positionLiquidated.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			order := models.GetLiquidationFromEvent(positionLiquidated, time)

			s.LiquidationsChan <- order
		}
	}
}
