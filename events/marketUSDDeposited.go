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

// MarketUSDDepositedSubscription is a struct for listening to all 'MarketUSDDeposited' contract events and return them as models.MarketUSDDeposited struct
type MarketUSDDepositedSubscription struct {
	*basicSubscription
	MarketUSDDepositedChan chan *models.MarketUSDDeposited
	contractEventChan      chan *core.CoreMarketUsdDeposited
}

func (e *Events) ListenMarketUSDDeposited() (*MarketUSDDepositedSubscription, error) {
	contractEventChan := make(chan *core.CoreMarketUsdDeposited)

	contractSub, err := e.core.WatchMarketUsdDeposited(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenMarketUSDDeposited").Errorf("error watch deposit: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "MarketUSDDeposited")
	}

	depositSub := newMarketUSDDepositedSubscription(contractSub, contractEventChan)

	go depositSub.listen(e.rpcClient)

	return depositSub, nil
}

// newMarketUSDDepositedSubscription is used to create new MarketUSDDepositedSubscription instance
func newMarketUSDDepositedSubscription(eventSub event.Subscription, contractEventChan chan *core.CoreMarketUsdDeposited) *MarketUSDDepositedSubscription {
	return &MarketUSDDepositedSubscription{
		basicSubscription:      newBasicSubscription(eventSub),
		contractEventChan:      contractEventChan,
		MarketUSDDepositedChan: make(chan *models.MarketUSDDeposited),
	}
}

// listen is used to run a goroutine
func (s *MarketUSDDepositedSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.MarketUSDDepositedChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-MarketUSDDeposited").Errorf("error listening deposit: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventMarketUSDDeposited := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventMarketUSDDeposited.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-MarketUSDDeposited").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					eventMarketUSDDeposited.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			deposited := models.GetMarketUSDDepositedFromEvent(eventMarketUSDDeposited, time)

			s.MarketUSDDepositedChan <- deposited
		}
	}
}
