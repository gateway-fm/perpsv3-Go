package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) RetrieveLiquidations(fromBlock uint64, toBLock *uint64) ([]*models.Liquidation, error) {
	opts := s.getFilterOptsPerpsMarket(fromBlock, toBLock)
	return s.retrieveLiquidations(opts)
}

func (s *Service) RetrieveLiquidationsLimit(limit uint64) ([]*models.Liquidation, error) {
	iterations, last, err := s.getIterationsForLimitQueryPerpsMarket(limit)
	if err != nil {
		return nil, err
	}

	var liquidations []*models.Liquidation

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsLimit").Infof(
		"fetching liquidations with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.perpsMarketFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveLiquidationsLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsPerpsMarket(fromBlock, &toBlock)

		res, err := s.retrieveLiquidations(opts)
		if err != nil {
			return nil, err
		}

		liquidations = append(liquidations, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsLimit").Infof("task completed successfully")

	return liquidations, nil
}

func (s *Service) retrieveLiquidations(opts *bind.FilterOpts) ([]*models.Liquidation, error) {
	iterator, err := s.perpsMarket.FilterPositionLiquidated(opts, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveLiquidations").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "perps market")
	}

	var liquidations []*models.Liquidation

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveLiquidations").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "perps market")
		}

		liquidation, err := s.getLiquidation(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		liquidations = append(liquidations, liquidation)
	}

	return liquidations, nil
}

// getLiquidation is used to get models.Liquidation from given event and block number
func (s *Service) getLiquidation(event *perpsMarket.PerpsMarketPositionLiquidated, blockN uint64) (*models.Liquidation, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveLiquidations").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetLiquidationFromEvent(event, block.Time), nil
}
