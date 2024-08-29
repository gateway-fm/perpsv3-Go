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

func (s *Service) RetrieveLiquidationsCoreLimit(limit uint64) ([]*models.CoreLiquidation, error) {
	//func (s *Service) RetrieveLiquidationsCore(limit uint64) ([]*models.CoreLiquidation, error) {
	iterations, last, err := s.getIterationsForLimitQueryCore(limit)
	if err != nil {
		return nil, err
	}

	var liquidations []*models.CoreLiquidation

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsCoreLimit").Infof(
		"fetching liquidations with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveLiquidationsCoreLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveLiquidationsCore(opts)
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

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsCoreLimit").Infof("task completed successfully")

	return liquidations, nil
}

func (s *Service) retrieveLiquidationsCore(opts *bind.FilterOpts) ([]*models.CoreLiquidation, error) {
	iterator, err := s.core.FilterLiquidation(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveLiquidationsCore").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var liquidations []*models.CoreLiquidation

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveLiquidationsCore").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		liquidation, err := s.getLiquidationCore(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		liquidations = append(liquidations, liquidation)
	}

	return liquidations, nil
}

// getLiquidationCore is used to get models.CoreLiquidation from given event and block number
func (s *Service) getLiquidationCore(event *core.CoreLiquidation, blockN uint64) (*models.CoreLiquidation, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveLiquidationsCore").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetCoreLiquidationFromEvent(event, block.Time), nil
}

func (s *Service) RetrieveVaultLiquidationsCoreLimit(limit uint64) ([]*models.CoreVaultLiquidation, error) {
	//func (s *Service) RetrieveVaultLiquidationsCore(limit uint64) ([]*models.CoreVaultLiquidation, error) {
	iterations, last, err := s.getIterationsForLimitQueryCore(limit)
	if err != nil {
		return nil, err
	}

	var liquidations []*models.CoreVaultLiquidation

	logger.Log().WithField("layer", "Service-RetrieveVaultLiquidationsCoreLimit").Infof(
		"fetching liquidations with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveVaultLiquidationsCoreLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveVaultLiquidationsCore(opts)
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

	logger.Log().WithField("layer", "Service-RetrieveVaultLiquidationsCoreLimit").Infof("task completed successfully")

	return liquidations, nil
}

func (s *Service) retrieveVaultLiquidationsCore(opts *bind.FilterOpts) ([]*models.CoreVaultLiquidation, error) {
	iterator, err := s.core.FilterVaultLiquidation(opts, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveVaultLiquidationsCore").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var liquidations []*models.CoreVaultLiquidation

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveVaultLiquidationsCore").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		liquidation, err := s.getVaultLiquidationCore(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		liquidations = append(liquidations, liquidation)
	}

	return liquidations, nil
}

// getVaultLiquidationCore is used to get models.CoreVaultLiquidation from given event and block number
func (s *Service) getVaultLiquidationCore(event *core.CoreVaultLiquidation, blockN uint64) (*models.CoreVaultLiquidation, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveVaultLiquidationsCore").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetCoreVaultLiquidationFromEvent(event, block.Time), nil
}

func (s *Service) RetrieveLiquidationsCore(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CoreLiquidation, error) {
}

// RetrieveVaultLiquidationsCore is used to get all `DelegationUpdated` events with given start block, end block and block search
// limit. For most public RPC providers the value for limit is 20 000 blocks
func (s *Service) RetrieveVaultLiquidationsCore(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CoreVaultLiquidation, error) {
}
