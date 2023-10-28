package events

import (
	"strings"

	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// AccountPermissionGrantedSubscription is a struct for listening to all 'PermissionGranted' contract events and return them as models.PermissionChanged struct
type AccountPermissionGrantedSubscription struct {
	*basicSubscription
	PermissionChangeChan chan *models.PermissionChanged
	contractEventChan    chan *perpsMarketGoerli.PerpsMarketGoerliPermissionGranted
}

func (e *Events) ListenAccountPermissionGranted() (*AccountPermissionGrantedSubscription, error) {
	createdChan := make(chan *perpsMarketGoerli.PerpsMarketGoerliPermissionGranted)

	createdSub, err := e.perpsMarket.WatchPermissionGranted(nil, createdChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-AccountPermissionGranted").Errorf("error watch permission granted: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "AccountPermissionGranted")
	}

	accountsSub := newAccountPermissionGrantedSubscription(createdSub, createdChan)

	go accountsSub.listen(e.perpsMarket)

	return accountsSub, nil
}

// newAccountPermissionGrantedSubscription is used to get new AccountPermissionGrantedSubscription instance
func newAccountPermissionGrantedSubscription(
	eventSub event.Subscription,
	created chan *perpsMarketGoerli.PerpsMarketGoerliPermissionGranted,
) *AccountPermissionGrantedSubscription {
	return &AccountPermissionGrantedSubscription{
		basicSubscription:    newBasicSubscription(eventSub),
		PermissionChangeChan: make(chan *models.PermissionChanged),
		contractEventChan:    created,
	}
}

// listen is used to run events listen goroutine
func (s *AccountPermissionGrantedSubscription) listen(perpsMarket *perpsMarketGoerli.PerpsMarketGoerli) {
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
				logger.Log().WithField("layer", "Events-AccountPermissionGranted").Errorf(
					"error listening account permission granted: %v", err.Error(),
				)
				s.ErrChan <- err
			}
			return
		case contractEvent := <-s.contractEventChan:
			p, err := models.PermissionFromString(strings.TrimRight(string(contractEvent.Permission[:]), string(rune(0))))
			if err != nil {
				logger.Log().WithField("layer", "Events-AccountPermissionGranted").Errorf(
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
