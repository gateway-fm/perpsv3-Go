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

// USDBurnedSubscription is a struct for listening to all 'UsdBurned' contract events and return them as models.USDBurned struct
type USDBurnedSubscription struct {
	*basicSubscription
	USDBurnedChan     chan *models.USDBurned
	contractEventChan chan *core.CoreUsdBurned
}

func (e *Events) ListenUSDBurned() (*USDBurnedSubscription, error) {
	contractEventChan := make(chan *core.CoreUsdBurned)

	contractSub, err := e.core.WatchUsdBurned(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenUSDBurned").Errorf("error watch usd burned: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "USDBurned")
	}

	orderSub := newUSDBurnedSubscription(contractSub, contractEventChan)

	go orderSub.listen(e.rpcClient)

	return orderSub, nil
}

// newUSDBurnedSubscription is used to create new USDBurnedSubscription instance
func newUSDBurnedSubscription(eventSub event.Subscription, contractEventChan chan *core.CoreUsdBurned) *USDBurnedSubscription {
	return &USDBurnedSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		contractEventChan: contractEventChan,
		USDBurnedChan:     make(chan *models.USDBurned),
	}
}

// listen is used to run a goroutine
func (s *USDBurnedSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.USDBurnedChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-UsdBurned").Errorf("error listening usd burned: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventUSDBurned := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventUSDBurned.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-UsdBurned").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					eventUSDBurned.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			burn := models.GetUSDBurnedFromEvent(eventUSDBurned, time)

			s.USDBurnedChan <- burn
		}
	}
}
