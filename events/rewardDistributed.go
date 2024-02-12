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

// RewardDistributedSubscription is a struct for listening to all 'RewardDistributed' contract events and return them as models.RewardDistributed struct
type RewardDistributedSubscription struct {
	*basicSubscription
	RewardDistributedChan chan *models.RewardDistributed
	contractEventChan     chan *core.CoreRewardsDistributed
}

func (e *Events) ListenRewardDistributed() (*RewardDistributedSubscription, error) {
	contractEventChan := make(chan *core.CoreRewardsDistributed)

	contractSub, err := e.core.WatchRewardsDistributed(nil, contractEventChan, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenRewardDistributed").Errorf("error watch reward claimed: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "RewardDistributed")
	}

	rewardSub := newRewardDistributedSubscription(contractSub, contractEventChan)

	go rewardSub.listen(e.rpcClient)

	return rewardSub, nil
}

// newRewardDistributedSubscription is used to create new RewardDistributedSubscription instance
func newRewardDistributedSubscription(eventSub event.Subscription, contractEventChan chan *core.CoreRewardsDistributed) *RewardDistributedSubscription {
	return &RewardDistributedSubscription{
		basicSubscription:     newBasicSubscription(eventSub),
		contractEventChan:     contractEventChan,
		RewardDistributedChan: make(chan *models.RewardDistributed),
	}
}

// listen is used to run a goroutine
func (s *RewardDistributedSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.RewardDistributedChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-RewardDistributed").Errorf("error listening reward claimed: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventRewardDistributed := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventRewardDistributed.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-RewardDistributed").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					eventRewardDistributed.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			distributed := models.GetRewardDistributedFromEvent(eventRewardDistributed, time)

			s.RewardDistributedChan <- distributed
		}
	}
}
