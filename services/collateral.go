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

func (s *Service) RetrieveCollateralWithdrawnLimit(limit uint64) ([]*models.CollateralWithdrawn, error) {
	iterations, last, err := s.getIterationsForLimitQuery(limit)
	if err != nil {
		return nil, err
	}

	var withdraws []*models.CollateralWithdrawn

	logger.Log().WithField("layer", "Service-RetrieveCollateralWithdrawnLimit").Infof(
		"fetching CollateralWithdrawn with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveCollateralWithdrawnLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveCollateralWithdrawn(opts)
		if err != nil {
			return nil, err
		}

		withdraws = append(withdraws, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveCollateralWithdrawnLimit").Infof("task completed successfully")

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
	iterations, last, err := s.getIterationsForLimitQuery(limit)
	if err != nil {
		return nil, err
	}

	var deposits []*models.CollateralDeposited

	logger.Log().WithField("layer", "Service-RetrieveCollateralDepositedLimit").Infof(
		"fetching CollateralDeposited with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveCollateralDepositedLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveCollateralDeposited(opts)
		if err != nil {
			return nil, err
		}

		deposits = append(deposits, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveCollateralDepositedLimit").Infof("task completed successfully")

	return deposits, nil
}

func (s *Service) retrieveCollateralDeposited(opts *bind.FilterOpts) ([]*models.CollateralDeposited, error) {
	iterator, err := s.core.FilterDeposited(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveCollateralDepositedLimit").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var deposits []*models.CollateralDeposited

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveCollateralDepositedLimit").Errorf("iterator error: %v", iterator.Error().Error())
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
		logger.Log().WithField("layer", "Service-RetrieveCollateralDepositedLimit").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetCollateralDepositedFromEvent(event, block.Time), nil
}
