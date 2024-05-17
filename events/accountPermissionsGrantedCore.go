package events

import (
	"strings"

	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// AccountCorePermissionGrantedSubscription is a struct for listening to all 'PermissionGranted' contract events and return them as models.PermissionChanged struct
type AccountCorePermissionGrantedSubscription struct {
	*basicSubscription
	PermissionChangeChan chan *models.PermissionChanged
	contractEventChan    chan *core.CorePermissionGranted
}

func (e *Events) ListenAccountCorePermissionGranted() (*AccountCorePermissionGrantedSubscription, error) {
	createdChan := make(chan *core.CorePermissionGranted)

	createdSub, err := e.core.WatchPermissionGranted(nil, createdChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-AccountCorePermissionGranted").Errorf("error watch permission granted: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "AccountCorePermissionGranted")
	}

	accountsSub := newAccountCorePermissionGrantedSubscription(createdSub, createdChan)

	go accountsSub.listen(e.core)

	return accountsSub, nil
}

// newAccountCorePermissionGrantedSubscription is used to get new AccountCorePermissionGrantedSubscription instance
func newAccountCorePermissionGrantedSubscription(
	eventSub event.Subscription,
	created chan *core.CorePermissionGranted,
) *AccountCorePermissionGrantedSubscription {
	return &AccountCorePermissionGrantedSubscription{
		basicSubscription:    newBasicSubscription(eventSub),
		PermissionChangeChan: make(chan *models.PermissionChanged),
		contractEventChan:    created,
	}
}

// listen is used to run events listen goroutine
func (s *AccountCorePermissionGrantedSubscription) listen(perps *core.Core) {
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
				logger.Log().WithField("layer", "Events-AccountCorePermissionGranted").Errorf(
					"error listening account permission granted: %v", err.Error(),
				)
				s.ErrChan <- err
			}
			return
		case contractEvent := <-s.contractEventChan:
			p, err := models.PermissionFromString(strings.TrimRight(string(contractEvent.Permission[:]), string(rune(0))))
			if err != nil {
				logger.Log().WithField("layer", "Events-AccountCorePermissionGranted").Errorf(
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
