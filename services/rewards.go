package services

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) RetrieveRewardClaimedLimit(limit uint64) ([]*models.RewardClaimed, error) {
	return s.RetrieveRewardClaimed(0, 0, limit)
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
	return s.RetrieveRewardDistributed(0, 0, limit)
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

func (s *Service) RetrieveRewardClaimed(fromBlock uint64, toBlock1 uint64, limit uint64) ([]*models.RewardClaimed, error) {
	iterations, lastBlock, err := s.getIterationsForQuery(fromBlock, toBlock1, limit, ContractCore)
	if err != nil {
		return nil, fmt.Errorf("cant getIterationsForQuery: %w", err)
	}

	var claims []*models.RewardClaimed

	logger.Log().WithField("layer", "Service-RetrieveRewardClaimed").Infof(
		"fetching RewardClaimed with limit: %v to block: %v total iterations: %v...",
		limit, lastBlock, iterations,
	)

	toBlock2 := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveRewardClaimed").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock2)

		res, err := s.retrieveRewardClaimed(opts)
		if err != nil {
			return nil, err
		}

		claims = append(claims, res...)

		fromBlock = toBlock2 + 1

		if i == iterations-1 {
			toBlock2 = lastBlock
		} else {
			toBlock2 = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveRewardClaimed").Infof("task completed successfully")

	return claims, nil
}

func (s *Service) RetrieveRewardDistributed(fromBlock uint64, toBlock1 uint64, limit uint64) ([]*models.RewardDistributed, error) {
	iterations, lastBlock, err := s.getIterationsForQuery(fromBlock, toBlock1, limit, ContractCore)
	if err != nil {
		return nil, fmt.Errorf("cant getIterationsForQuery: %w", err)
	}

	var distributions []*models.RewardDistributed

	logger.Log().WithField("layer", "Service-RetrieveRewardDistributed").Infof(
		"fetching RewardDistributed with limit: %v to block: %v total iterations: %v...",
		limit, lastBlock, iterations,
	)

	toBlock2 := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveRewardDistributed").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock2)

		res, err := s.retrieveRewardDistributed(opts)
		if err != nil {
			return nil, err
		}

		distributions = append(distributions, res...)

		fromBlock = toBlock2 + 1

		if i == iterations-1 {
			toBlock2 = lastBlock
		} else {
			toBlock2 = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveRewardDistributed").Infof("task completed successfully")

	return distributions, nil
}
