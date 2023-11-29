package events

import (
	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// AccountLiquidatedSubscription is a struct for listening to all 'AccountLiquidated' contract events and return them as models.AccountLiquidated struct
type AccountLiquidatedSubscription struct {
	*basicSubscription
	AccountLiquidated chan *models.AccountLiquidated
	contractEventChan chan *perpsMarket.PerpsMarketAccountLiquidationAttempt
}

func (e *Events) ListenAccountLiquidated() (*AccountLiquidatedSubscription, error) {
	createdChan := make(chan *perpsMarket.PerpsMarketAccountLiquidationAttempt)

	liquidatedSub, err := e.perpsMarket.WatchAccountLiquidationAttempt(nil, createdChan, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-AccountLiquidated").Errorf("error watch account liquidated: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "AccountLiquidated")
	}

	accountsSub := newAccountLiquidatedSubscription(liquidatedSub, createdChan)

	go accountsSub.listen(e.perpsMarket)

	return accountsSub, nil
}

// newAccountsSubscription
func newAccountLiquidatedSubscription(
	eventSub event.Subscription,
	created chan *perpsMarket.PerpsMarketAccountLiquidationAttempt,
) *AccountLiquidatedSubscription {
	return &AccountLiquidatedSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		AccountLiquidated: make(chan *models.AccountLiquidated),
		contractEventChan: created,
	}
}

// listen is used to run events listen goroutine
func (s *AccountLiquidatedSubscription) listen(perps *perpsMarket.PerpsMarket) {
	defer func() {
		close(s.AccountLiquidated)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-AccountLiquidated").Errorf(
					"error listening account created info: %v", err.Error(),
				)
				s.ErrChan <- err
			}
			return
		case contractEvent := <-s.contractEventChan:
			liquidated := &models.AccountLiquidated{
				ID:             contractEvent.AccountId,
				Reward:         contractEvent.Reward,
				FullLiquidated: contractEvent.FullLiquidation,
			}

			s.AccountLiquidated <- liquidated
		}
	}
}
