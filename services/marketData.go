package services

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) RetrieveMarketUpdatesLimit(limit uint64) ([]*models.MarketUpdate, error) {
	iterations, last, err := s.getIterationsForLimitQuery(limit)
	if err != nil {
		return nil, err
	}

	var marketUpdates []*models.MarketUpdate

	fromBlock := s.perpsMarketFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		opts := s.getFilterOptsPerpsMarket(fromBlock, &toBlock)

		res, err := s.retrieveMarketUpdates(opts)
		if err != nil {
			return nil, err
		}

		marketUpdates = append(marketUpdates, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	return marketUpdates, nil
}

func (s *Service) RetrieveMarketUpdates(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdate, error) {
	opts := s.getFilterOptsPerpsMarket(fromBlock, toBLock)
	return s.retrieveMarketUpdates(opts)
}

// retrieveMarketUpdates is used to get retrieve market updates with given filter options
func (s *Service) retrieveMarketUpdates(opts *bind.FilterOpts) ([]*models.MarketUpdate, error) {
	iterator, err := s.perpsMarket.FilterMarketUpdated(opts)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveMarketUpdates").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "perps market")
	}

	var marketUpdates []*models.MarketUpdate

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveMarketUpdates").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "perps market")
		}

		marketUpdate, err := s.getMarketUpdate(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		marketUpdates = append(marketUpdates, marketUpdate)
	}

	return marketUpdates, nil
}

// getMarketUpdate is used to get models.MarketUpdate from given event and block number
func (s *Service) getMarketUpdate(event *perpsMarketGoerli.PerpsMarketGoerliMarketUpdated, blockN uint64) (*models.MarketUpdate, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-getMarketUpdate").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetMarketUpdateFromEvent(event, block.Time), nil
}
