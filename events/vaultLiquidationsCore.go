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

// VaultLiquidationsCoreSubscription is a struct for listening to all 'VaultLiquidationsCore' contract events and return them as models.Account struct
type VaultLiquidationsCoreSubscription struct {
	*basicSubscription
	NewAccountChan    chan *models.CoreVaultLiquidation
	contractEventChan chan *core.CoreVaultLiquidation
}

func (e *Events) ListenVaultLiquidationsCore() (*VaultLiquidationsCoreSubscription, error) {
	liquidationChan := make(chan *core.CoreVaultLiquidation)

	liquidationSub, err := e.core.WatchVaultLiquidation(nil, liquidationChan, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-VaultLiquidationsCore").Errorf("error watch liquidation: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "VaultLiquidationsCore")
	}

	liquidationsSub := newVaultLiquidationsCoreSubscription(liquidationSub, liquidationChan)

	go liquidationsSub.listen(e.rpcClient)

	return liquidationsSub, nil
}

// newVaultLiquidationsCoreSubscription is used to get new VaultLiquidationsCoreSubscription instance
func newVaultLiquidationsCoreSubscription(
	eventSub event.Subscription,
	liquidation chan *core.CoreVaultLiquidation,
) *VaultLiquidationsCoreSubscription {
	return &VaultLiquidationsCoreSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		NewAccountChan:    make(chan *models.CoreVaultLiquidation),
		contractEventChan: liquidation,
	}
}

// listen is used to run events listen goroutine
func (s *VaultLiquidationsCoreSubscription) listen(rpcClient *ethclient.Client) {
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
				logger.Log().WithField("layer", "Events-VaultLiquidationsCore").Errorf(
					"error listening liquidation info: %v", err.Error(),
				)
				s.ErrChan <- err
			}
			return
		case liquidation := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(liquidation.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-VaultLiquidationsCore").Warningf(
					"error fetching block number %v: %v; liquidation event time set to 0 ",
					liquidation.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			s.NewAccountChan <- models.GetCoreVaultLiquidationFromEvent(liquidation, time)
		}
	}
}
