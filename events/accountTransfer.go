package events

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/gateway-fm/perpsv3-Go/contracts/Account"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
	"math/big"
)

// AccountTransferSubscription is a struct for listening to all 'AccountTransfer' contract events and return them as models.AccountTransfer struct
type AccountTransferSubscription struct {
	*basicSubscription
	NewAccountChan    chan *models.AccountTransfer
	contractEventChan chan *Account.AccountTransfer
}

func (e *Events) ListenAccountTransfer() (*AccountTransferSubscription, error) {
	transferChan := make(chan *Account.AccountTransfer)

	transferSub, err := e.account.WatchTransfer(nil, transferChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-Accounts").Errorf("error watch account transfer: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "AccountTransfer")
	}

	accountsSub := newAccountTransferSubscription(transferSub, transferChan)

	go accountsSub.listen(e.rpcClient)

	return accountsSub, nil
}

// newAccountTransferSubscription is used to get new AccountTransferSubscription instance
func newAccountTransferSubscription(
	eventSub event.Subscription,
	transfer chan *Account.AccountTransfer,
) *AccountTransferSubscription {
	return &AccountTransferSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		NewAccountChan:    make(chan *models.AccountTransfer),
		contractEventChan: transfer,
	}
}

// listen is used to run events listen goroutine
func (s *AccountTransferSubscription) listen(rpcClient *ethclient.Client) {
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
				logger.Log().WithField("layer", "Events-AccountTransfer").Errorf(
					"error listening account transfer info: %v", err.Error(),
				)
				s.ErrChan <- err
			}
			return
		case transfer := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(transfer.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-AccountTransfer").Warningf(
					"error fetching block number %v: %v; transfer event time set to 0 ",
					transfer.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			s.NewAccountChan <- models.GetAccountTransferFromEvent(transfer, time)
		}
	}
}
