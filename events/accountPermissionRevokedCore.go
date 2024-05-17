package events

import (
	"strings"

	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// AccountCorePermissionRevokedSubscription is a struct for listening to all 'PermissionRevoked' contract events and return them as models.PermissionChanged struct
type AccountCorePermissionRevokedSubscription struct {
	*basicSubscription
	PermissionChangeChan chan *models.PermissionChanged
	contractEventChan    chan *core.CorePermissionRevoked
}

func (e *Events) ListenAccountCorePermissionRevoked() (*AccountCorePermissionRevokedSubscription, error) {
	revokedChan := make(chan *core.CorePermissionRevoked)

	revokedSub, err := e.core.WatchPermissionRevoked(nil, revokedChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-AccountCorePermissionRevoked").Errorf("error watch permission revoked: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "AccountCorePermissionRevoked")
	}

	accountsSub := newAccountCorePermissionRevokedSubscription(revokedSub, revokedChan)

	go accountsSub.listen(e.core)

	return accountsSub, nil
}

// newAccountCorePermissionRevokedSubscription is used to get new AccountCorePermissionRevokedSubscription instance
func newAccountCorePermissionRevokedSubscription(
	eventSub event.Subscription,
	revoked chan *core.CorePermissionRevoked,
) *AccountCorePermissionRevokedSubscription {
	return &AccountCorePermissionRevokedSubscription{
		basicSubscription:    newBasicSubscription(eventSub),
		PermissionChangeChan: make(chan *models.PermissionChanged),
		contractEventChan:    revoked,
	}
}

// listen is used to run events listen goroutine
func (s *AccountCorePermissionRevokedSubscription) listen(perps *core.Core) {
	defer func() {
		close(s.PermissionChangeChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-AccountCorePermissionRevoked").Errorf(
					"error listening account permission revoked: %v", err.Error(),
				)
				s.ErrChan <- err
			}
			return
		case contractEvent := <-s.contractEventChan:
			p, err := models.PermissionFromString(strings.TrimRight(string(contractEvent.Permission[:]), string(rune(0))))
			if err != nil {
				logger.Log().WithField("layer", "Events-AccountCorePermissionRevoked").Errorf(
					"error decode permission %v: %v", string(contractEvent.Permission[:]), err.Error(),
				)
				s.ErrChan <- err
				continue
			}

			change := &models.PermissionChanged{
				AccountID:  contractEvent.AccountId,
				User:       contractEvent.User,
				Permission: p,
			}

			s.PermissionChangeChan <- change
		}
	}
}
