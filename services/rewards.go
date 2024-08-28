package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) RetrieveRewardClaimedLimit(limit uint64) ([]*models.RewardClaimed, error) {
	iterations, last, err := s.getIterationsForLimitQueryCore(limit)
	if err != nil {
		return nil, err
	}

	var claims []*models.RewardClaimed

	logger.Log().WithField("layer", "Service-RetrieveRewardClaimedLimit").Infof(
		"fetching RewardClaimed with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveRewardClaimedLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveRewardClaimed(opts)
		if err != nil {
			return nil, err
		}

		claims = append(claims, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveRewardClaimedLimit").Infof("task completed successfully")

	return claims, nil
}

func (s *Service) retrieveRewardClaimed(opts *bind.FilterOpts) ([]*models.RewardClaimed, error) {
	iterator, err := s.core.FilterRewardsClaimed(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveRewardClaimedLimit").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var claims []*models.RewardClaimed

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveRewardClaimedLimit").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		claim, err := s.getRewardClaimed(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		claims = append(claims, claim)
	}

	return claims, nil
}

// getRewardClaimed is used to get models.RewardClaimed from given event and block number
func (s *Service) getRewardClaimed(event *core.CoreRewardsClaimed, blockN uint64) (*models.RewardClaimed, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveRewardClaimedLimit").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetRewardClaimedFromEvent(event, block.Time), nil
}

func (s *Service) RetrieveRewardDistributedLimit(limit uint64) ([]*models.RewardDistributed, error) {
	iterations, last, err := s.getIterationsForLimitQueryCore(limit)
	if err != nil {
		return nil, err
	}

	var distributions []*models.RewardDistributed

	logger.Log().WithField("layer", "Service-RetrieveRewardDistributedLimit").Infof(
		"fetching RewardDistributed with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveRewardDistributedLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveRewardDistributed(opts)
		if err != nil {
			return nil, err
		}

		distributions = append(distributions, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveRewardDistributedLimit").Infof("task completed successfully")

	return distributions, nil
}

func (s *Service) retrieveRewardDistributed(opts *bind.FilterOpts) ([]*models.RewardDistributed, error) {
	iterator, err := s.core.FilterRewardsDistributed(opts, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveRewardDistributedLimit").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var distributions []*models.RewardDistributed

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveRewardDistributedLimit").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		distribution, err := s.getRewardDistributed(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		distributions = append(distributions, distribution)
	}

	return distributions, nil
}

// getRewardDistributed is used to get models.RewardDistributed from given event and block number
func (s *Service) getRewardDistributed(event *core.CoreRewardsDistributed, blockN uint64) (*models.RewardDistributed, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveRewardDistributedLimit").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetRewardDistributedFromEvent(event, block.Time), nil
}
