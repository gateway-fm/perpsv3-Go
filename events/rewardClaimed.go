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

// RewardClaimedSubscription is a struct for listening to all 'RewardClaimed' contract events and return them as models.RewardClaimed struct
type RewardClaimedSubscription struct {
	*basicSubscription
	RewardClaimedChan chan *models.RewardClaimed
	contractEventChan chan *core.CoreRewardsClaimed
}

func (e *Events) ListenRewardClaimed() (*RewardClaimedSubscription, error) {
	contractEventChan := make(chan *core.CoreRewardsClaimed)

	contractSub, err := e.core.WatchRewardsClaimed(nil, contractEventChan, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Events-ListenRewardClaimed").Errorf("error watch reward claimed: %v", err.Error())
		return nil, errors.GetEventListenErr(err, "RewardClaimed")
	}

	rewardSub := newRewardClaimedSubscription(contractSub, contractEventChan)

	go rewardSub.listen(e.rpcClient)

	return rewardSub, nil
}

// newRewardClaimedSubscription is used to create new RewardClaimedSubscription instance
func newRewardClaimedSubscription(eventSub event.Subscription, contractEventChan chan *core.CoreRewardsClaimed) *RewardClaimedSubscription {
	return &RewardClaimedSubscription{
		basicSubscription: newBasicSubscription(eventSub),
		contractEventChan: contractEventChan,
		RewardClaimedChan: make(chan *models.RewardClaimed),
	}
}

// listen is used to run a goroutine
func (s *RewardClaimedSubscription) listen(rpcClient *ethclient.Client) {
	defer func() {
		close(s.RewardClaimedChan)
		close(s.contractEventChan)
	}()

	for {
		select {
		case <-s.stop:
			return
		case err := <-s.eventSub.Err():
			if err != nil {
				logger.Log().WithField("layer", "Events-RewardClaimed").Errorf("error listening reward claimed: %v", err.Error())
				s.ErrChan <- err
			}
			return
		case eventRewardClaimed := <-s.contractEventChan:
			block, err := rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(eventRewardClaimed.Raw.BlockNumber)))
			time := uint64(0)
			if err != nil {
				logger.Log().WithField("layer", "Events-RewardClaimed").Warningf(
					"error fetching block number %v: %v; order event time set to 0 ",
					eventRewardClaimed.Raw.BlockNumber, err.Error(),
				)
				s.ErrChan <- err
			} else {
				time = block.Time
			}

			claimed := models.GetRewardClaimedFromEvent(eventRewardClaimed, time)

			s.RewardClaimedChan <- claimed
		}
	}
}
