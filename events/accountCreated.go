package events

import (
	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// AccountCreatedSubscription is a struct for listening to all 'AccountCreated' contract events and return them as models.Account struct
type AccountCreatedSubscription struct {
	*basicSubscription
	NewAccountChan    chan *models.Account
	contractEventChan chan *perpsMarketGoerli.PerpsMarketGoerliAccountCreated
}

func (e *Events) ListenAccountCreated() (*AccountCreatedSubscription, error) {
	createdChan := make(chan *perpsMarketGoerli.PerpsMarketGoerliAccountCreated)

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
	created chan *perpsMarketGoerli.PerpsMarketGoerliAccountCreated,
) *AccountCreatedSubscription {
	return &AccountCreatedSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		NewAccountChan:    make(chan *models.Account),
		contractEventChan: created,
	}
}

// listen is used to run events listen goroutine
func (s *AccountCreatedSubscription) listen(perpsMarket *perpsMarketGoerli.PerpsMarketGoerli) {
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
			lastInteraction := getAccountLastInteraction(newAccount.AccountId, perpsMarket)

			account := models.FormatAccount(
				newAccount.AccountId,
				newAccount.Owner,
				lastInteraction.Uint64(),
				[]perpsMarketGoerli.IAccountModuleAccountPermissions{},
			)

			s.NewAccountChan <- account
		}
	}
}
