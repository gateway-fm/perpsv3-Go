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

// USDMintedSubscription is a struct for listening to all 'UsdMinted' contract events and return them as models.USDMinted struct
type USDMintedSubscription struct {
	*basicSubscription
	USDMintedsChan    chan *models.USDMinted
	contractEventChan chan *core.CoreUsdMinted
}

func (e *Events) ListenUSDMinted() (*USDMintedSubscription, error) {
	contractEventChan := make(chan *core.CoreUsdMinted)

	contractSub, err := e.core.WatchUsdMinted(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenUSDMinted").Errorf("error watch usd minted: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "USDMinted")
	}

	orderSub := newUSDMintedSubscription(contractSub, contractEventChan)

	go orderSub.listen(e.rpcClient)

	return orderSub, nil
}

// newUSDMintedSubscription is used to create new USDMintedSubscription instance
func newUSDMintedSubscription(eventSub event.Subscription, contractEventChan chan *core.CoreUsdMinted) *USDMintedSubscription {
	return &USDMintedSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		contractEventChan: contractEventChan,
		USDMintedsChan:    make(chan *models.USDMinted),
	}
}

// listen is used to run a goroutine
func (s *USDMintedSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.USDMintedsChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-UsdMinted").Errorf("error listening usd minted: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventUsdMinted := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventUsdMinted.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-UsdMinted").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					eventUsdMinted.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			mint := models.GetUSDMintedFromEvent(eventUsdMinted, time)

			s.USDMintedsChan <- mint
		}
	}
}
