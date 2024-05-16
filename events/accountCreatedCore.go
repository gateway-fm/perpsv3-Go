package events

import (
	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// AccountCreatedCoreSubscription is a struct for listening to all 'AccountCreatedCore' contract events and return them as models.Account struct
type AccountCreatedCoreSubscription struct {
	*basicSubscription
	NewAccountChan    chan *models.Account
	contractEventChan chan *core.CoreAccountCreated
}

func (e *Events) ListenAccountCreatedCore() (*AccountCreatedCoreSubscription, error) {
	createdChan := make(chan *core.CoreAccountCreated)

	createdSub, err := e.core.WatchAccountCreated(nil, createdChan, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-Accounts").Errorf("error watch account created: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "AccountCreatedCore")
	}

	accountsSub := newAccountCreatedCoreSubscription(createdSub, createdChan)

	go accountsSub.listen(e.core)

	return accountsSub, nil
}

// newAccountCreatedCoreSubscription is used to get new AccountCreatedCoreSubscription instance
func newAccountCreatedCoreSubscription(
	eventSub event.Subscription,
	created chan *core.CoreAccountCreated,
) *AccountCreatedCoreSubscription {
	return &AccountCreatedCoreSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		NewAccountChan:    make(chan *models.Account),
		contractEventChan: created,
	}
}

// listen is used to run events listen goroutine
func (s *AccountCreatedCoreSubscription) listen(coreContract *core.Core) {
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
				logger.Log().WithField("layer", "Events-AccountCreatedCore").Errorf(
					"error listening account created info: %v", err.Error(),
				)
				s.ErrChan <- err
			}
			return
		case newAccount := <-s.contractEventChan:
			lastInteraction := getAccountLastInteractionCore(newAccount.AccountId, coreContract)

			account := models.FormatAccountCore(
				newAccount.AccountId,
				newAccount.Owner,
				lastInteraction.Uint64(),
				[]core.IAccountModuleAccountPermissions{},
			)

			s.NewAccountChan <- account
		}
	}
}
