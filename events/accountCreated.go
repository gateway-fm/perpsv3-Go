package events

import (
	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// AccountCreatedSubscription is a struct for listening to all 'AccountCreated' contract events and return them as models.Account struct
type AccountCreatedSubscription struct {
	*basicSubscription
	NewAccountChan    chan *models.Account
	contractEventChan chan *perpsMarket.PerpsMarketAccountCreated
}

func (e *Events) ListenAccountCreated() (*AccountCreatedSubscription, error) {
	createdChan := make(chan *perpsMarket.PerpsMarketAccountCreated)

	createdSub, err := e.perpsMarket.WatchAccountCreated(nil, createdChan, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-Accounts").Errorf("error watch account created: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "AccountCreated")
	}

	accountsSub := newAccountCreatedSubscription(createdSub, createdChan)

	go accountsSub.listen(e.perpsMarket)

	return accountsSub, nil
}

// newAccountCreatedSubscription is used to get new AccountCreatedSubscription instance
func newAccountCreatedSubscription(
	eventSub event.Subscription,
	created chan *perpsMarket.PerpsMarketAccountCreated,
) *AccountCreatedSubscription {
	return &AccountCreatedSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		NewAccountChan:    make(chan *models.Account),
		contractEventChan: created,
	}
}

// listen is used to run events listen goroutine
func (s *AccountCreatedSubscription) listen(perps *perpsMarket.PerpsMarket) {
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
				logger.Log().WithField("layer", "Events-AccountCreated").Errorf(
					"error listening account created info: %v", err.Error(),
				)
				s.ErrChan <- err
			}
			return
		case newAccount := <-s.contractEventChan:
			lastInteraction := getAccountLastInteraction(newAccount.AccountId, perps)

			account := models.FormatAccount(
				newAccount.AccountId,
				newAccount.Owner,
				lastInteraction.Uint64(),
				[]perpsMarket.IAccountModuleAccountPermissions{},
			)

			s.NewAccountChan <- account
		}
	}
}
