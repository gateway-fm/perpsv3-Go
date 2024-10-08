package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) GetCollateralPrice(blockNumber *big.Int, collateralType common.Address) (*models.CollateralPrice, error) {
	var opts *bind.CallOpts
	if blockNumber != nil && blockNumber.Int64() > 0 {
		opts = &bind.CallOpts{BlockNumber: blockNumber}
	}

	price, err := s.core.GetCollateralPrice(opts, collateralType)
	if err != nil {
		return nil, errors.GetReadContractErr(err, "core", "GetCollateralPrice")
	}

	return &models.CollateralPrice{Price: price}, nil
}

func (s *Service) RetrieveCollateralWithdrawnLimit(limit uint64) ([]*models.CollateralWithdrawn, error) {
	return s.RetrieveCollateralWithdrawn(0, 0, limit)
}

func (s *Service) RetrieveCollateralWithdrawn(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CollateralWithdrawn, error) {
	iterations, lastBlock, err := s.getIterationsForQuery(fromBlock, toBlock, limit, ContractCore)
	if err != nil {
		return nil, err
	}

	if fromBlock == 0 {
		fromBlock = s.coreFirstBlock
	}

	var withdraws []*models.CollateralWithdrawn

	logger.Log().WithField("layer", "Service-RetrieveCollateralWithdrawn").Infof(
		"fetching CollateralWithdrawn with limit: %v from block %v to block: %v total iterations: %v...",
		limit, fromBlock, lastBlock, iterations,
	)

	startBlockOfIteration := fromBlock
	endBlockOfIteration := startBlockOfIteration + limit

	if endBlockOfIteration > toBlock {
		endBlockOfIteration = toBlock
	}

	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveCollateralWithdrawn").Infof("-- iteration %v", i)
		}

		opts := s.getFilterOptsCore(startBlockOfIteration, &endBlockOfIteration)

		res, err := s.retrieveCollateralWithdrawn(opts)
		if err != nil {
			return nil, err
		}

		withdraws = append(withdraws, res...)

		startBlockOfIteration = endBlockOfIteration + 1

		if i == iterations-1 {
			endBlockOfIteration = lastBlock
		} else {
			endBlockOfIteration = startBlockOfIteration + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveCollateralWithdrawn").Infof("task completed successfully")

	return withdraws, nil
}

func (s *Service) retrieveCollateralWithdrawn(opts *bind.FilterOpts) ([]*models.CollateralWithdrawn, error) {
	iterator, err := s.core.FilterWithdrawn(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveCollateralWithdrawnLimit").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var withdraws []*models.CollateralWithdrawn

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveCollateralWithdrawnLimit").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		withdraw, err := s.getCollateralWithdrawn(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		withdraws = append(withdraws, withdraw)
	}

	return withdraws, nil
}

// getCollateralWithdrawn is used to get models.CollateralWithdrawn from given event and block number
func (s *Service) getCollateralWithdrawn(event *core.CoreWithdrawn, blockN uint64) (*models.CollateralWithdrawn, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveCollateralWithdrawnLimit").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetCollateralWithdrawnFromEvent(event, block.Time), nil
}

func (s *Service) RetrieveCollateralDepositedLimit(limit uint64) ([]*models.CollateralDeposited, error) {
	return s.RetrieveCollateralDeposited(0, 0, limit)
}

func (s *Service) RetrieveCollateralDeposited(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CollateralDeposited, error) {
	iterations, lastBlock, err := s.getIterationsForQuery(fromBlock, toBlock, limit, ContractCore)
	if err != nil {
		return nil, err
	}

	if fromBlock == 0 {
		fromBlock = s.coreFirstBlock
	}

	var deposits []*models.CollateralDeposited

	logger.Log().WithField("layer", "Service-RetrieveCollateralDeposited").Infof(
		"fetching CollateralDeposited with limit: %v from block %v to block: %v total iterations: %v...",
		limit, fromBlock, lastBlock, iterations,
	)

	startBlockOfIteration := fromBlock
	endBlockOfIteration := startBlockOfIteration + limit

	if endBlockOfIteration > toBlock {
		endBlockOfIteration = toBlock
	}

	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveCollateralDeposited").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(startBlockOfIteration, &endBlockOfIteration)

		res, err := s.retrieveCollateralDeposited(opts)
		if err != nil {
			return nil, err
		}

		deposits = append(deposits, res...)

		startBlockOfIteration = endBlockOfIteration + 1

		if i == iterations-1 {
			endBlockOfIteration = lastBlock
		} else {
			endBlockOfIteration = startBlockOfIteration + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveCollateralDeposited").Infof("task completed successfully")

	return deposits, nil
}

func (s *Service) retrieveCollateralDeposited(opts *bind.FilterOpts) ([]*models.CollateralDeposited, error) {
	iterator, err := s.core.FilterDeposited(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveCollateralDeposited").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var deposits []*models.CollateralDeposited

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveCollateralDeposited").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		deposit, err := s.getCollateralDeposited(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		deposits = append(deposits, deposit)
	}

	return deposits, nil
}

// getCollateralDeposited is used to get models.CollateralDeposited from given event and block number
func (s *Service) getCollateralDeposited(event *core.CoreDeposited, blockN uint64) (*models.CollateralDeposited, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveCollateralDeposited").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetCollateralDepositedFromEvent(event, block.Time), nil
}

func (s *Service) GetCollateralConfigurations(hideDisabled bool) ([]*models.CollateralConfiguration, error) {
	data, err := s.core.GetCollateralConfigurations(nil, hideDisabled)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetCollateralConfigurations").Errorf("get collateral configurations error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "core", "GetCollateralConfigurations")
	}

	return models.GetCollateralConfigurations(data), nil
}
