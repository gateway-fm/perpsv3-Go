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

func (s *Service) RetrieveLiquidationsCoreLimit(limit uint64) ([]*models.CoreLiquidation, error) {
	return s.RetrieveLiquidationsCore(0, 0, limit)
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
	return s.RetrieveVaultLiquidationsCore(0, 0, limit)
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

func (s *Service) RetrieveLiquidationsCore(fromBlock uint64, toBlock1 uint64, limit uint64) ([]*models.CoreLiquidation, error) {
	iterations, lastBlock, err := s.getIterationsForQuery(fromBlock, toBlock1, limit, ContractCore)
	if err != nil {
		return nil, fmt.Errorf("cant getIterationsForQuery: %w", err)
	}

	var liquidations []*models.CoreLiquidation

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsCore").Infof(
		"fetching liquidations with limit: %v to block: %v total iterations: %v...",
		limit, lastBlock, iterations,
	)

	toBlock2 := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveLiquidationsCore").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock2)

		res, err := s.retrieveLiquidationsCore(opts)
		if err != nil {
			return nil, err
		}

		liquidations = append(liquidations, res...)

		fromBlock = toBlock2 + 1

		if i == iterations-1 {
			toBlock2 = lastBlock
		} else {
			toBlock2 = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsCore").Infof("task completed successfully")

	return liquidations, nil
}

// RetrieveVaultLiquidationsCore is used to get all `DelegationUpdated` events with given start block, end block and block search
// limit. For most public RPC providers the value for limit is 20 000 blocks
func (s *Service) RetrieveVaultLiquidationsCore(fromBlock uint64, toBlock1 uint64, limit uint64) ([]*models.CoreVaultLiquidation, error) {
	iterations, lastBlock, err := s.getIterationsForQuery(fromBlock, toBlock1, limit, ContractCore)
	if err != nil {
		return nil, fmt.Errorf("cant getIterationsForQuery: %w", err)
	}

	var liquidations []*models.CoreVaultLiquidation

	logger.Log().WithField("layer", "Service-RetrieveVaultLiquidationsCore").Infof(
		"fetching liquidations with limit: %v to block: %v total iterations: %v...",
		limit, lastBlock, iterations,
	)

	toBlock2 := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveVaultLiquidationsCore").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock2)

		res, err := s.retrieveVaultLiquidationsCore(opts)
		if err != nil {
			return nil, err
		}

		liquidations = append(liquidations, res...)

		fromBlock = toBlock2 + 1

		if i == iterations-1 {
			toBlock2 = lastBlock
		} else {
			toBlock2 = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveVaultLiquidationsCore").Infof("task completed successfully")

	return liquidations, nil
}
