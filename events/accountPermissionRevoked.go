package events

import (
	"strings"

	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// AccountPermissionRevokedSubscription is a struct for listening to all 'PermissionRevoked' contract events and return them as models.PermissionChanged struct
type AccountPermissionRevokedSubscription struct {
	*basicSubscription
	PermissionChangeChan chan *models.PermissionChanged
	contractEventChan    chan *perpsMarket.PerpsMarketPermissionRevoked
}

func (e *Events) ListenAccountPermissionRevoked() (*AccountPermissionRevokedSubscription, error) {
	revokedChan := make(chan *perpsMarket.PerpsMarketPermissionRevoked)

	revokedSub, err := e.perpsMarket.WatchPermissionRevoked(nil, revokedChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-AccountPermissionRevoked").Errorf("error watch permission revoked: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "AccountPermissionRevoked")
	}

	accountsSub := newAccountPermissionRevokedSubscription(revokedSub, revokedChan)

	go accountsSub.listen(e.perpsMarket)

	return accountsSub, nil
}

// newAccountPermissionRevokedSubscription is used to get new AccountPermissionRevokedSubscription instance
func newAccountPermissionRevokedSubscription(
	eventSub event.Subscription,
	revoked chan *perpsMarket.PerpsMarketPermissionRevoked,
) *AccountPermissionRevokedSubscription {
	return &AccountPermissionRevokedSubscription{
		basicSubscription:    newBasicSubscription(eventSub),
		PermissionChangeChan: make(chan *models.PermissionChanged),
		contractEventChan:    revoked,
	}
}

// listen is used to run events listen goroutine
func (s *AccountPermissionRevokedSubscription) listen(perps *perpsMarket.PerpsMarket) {
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
				logger.Log().WithField("layer", "Events-AccountPermissionRevoked").Errorf(
					"error listening account permission revoked: %v", err.Error(),
				)
				s.ErrChan <- err
			}
			return
		case contractEvent := <-s.contractEventChan:
			p, err := models.PermissionFromString(strings.TrimRight(string(contractEvent.Permission[:]), string(rune(0))))
			if err != nil {
				logger.Log().WithField("layer", "Events-AccountPermissionRevoked").Errorf(
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
