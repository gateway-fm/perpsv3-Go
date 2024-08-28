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

func (s *Service) RetrieveMarketUSDDepositedLimit(limit uint64) ([]*models.MarketUSDDeposited, error) {
	iterations, last, err := s.getIterationsForLimitQueryCore(limit)
	if err != nil {
		return nil, err
	}

	var deposits []*models.MarketUSDDeposited

	logger.Log().WithField("layer", "Service-RetrieveMarketUSDDepositedLimit").Infof(
		"fetching MarketUSDDeposited with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveMarketUSDDepositedLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveMarketUSDDeposited(opts)
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

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsLimit").Infof("task completed successfully")

	return deposits, nil
}

func (s *Service) retrieveMarketUSDDeposited(opts *bind.FilterOpts) ([]*models.MarketUSDDeposited, error) {
	iterator, err := s.core.FilterMarketUsdDeposited(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveMarketUSDDepositedLimit").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var deposits []*models.MarketUSDDeposited

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveMarketUSDDepositedLimit").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		mint, err := s.getMarketUSDDeposited(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		deposits = append(deposits, mint)
	}

	return deposits, nil
}

func (s *Service) getMarketUSDDeposited(event *core.CoreMarketUsdDeposited, blockN uint64) (*models.MarketUSDDeposited, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveMarketUSDDepositedLimit").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetMarketUSDDepositedFromEvent(event, block.Time), nil
}

func (s *Service) RetrieveMarketUSDWithdrawnLimit(limit uint64) ([]*models.MarketUSDWithdrawn, error) {
	iterations, last, err := s.getIterationsForLimitQueryCore(limit)
	if err != nil {
		return nil, err
	}

	var deposits []*models.MarketUSDWithdrawn

	logger.Log().WithField("layer", "Service-RetrieveMarketUSDWithdrawnLimit").Infof(
		"fetching MarketUSDWithdrawn with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveMarketUSDWithdrawnLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveMarketUSDWithdrawn(opts)
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

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsLimit").Infof("task completed successfully")

	return deposits, nil
}

func (s *Service) retrieveMarketUSDWithdrawn(opts *bind.FilterOpts) ([]*models.MarketUSDWithdrawn, error) {
	iterator, err := s.core.FilterMarketUsdWithdrawn(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveMarketUSDWithdrawnLimit").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var deposits []*models.MarketUSDWithdrawn

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveMarketUSDWithdrawnLimit").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		mint, err := s.getMarketUSDWithdrawn(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		deposits = append(deposits, mint)
	}

	return deposits, nil
}

func (s *Service) getMarketUSDWithdrawn(event *core.CoreMarketUsdWithdrawn, blockN uint64) (*models.MarketUSDWithdrawn, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveMarketUSDWithdrawnLimit").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetMarketUSDWithdrawnFromEvent(event, block.Time), nil
}
