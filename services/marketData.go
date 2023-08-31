package services

import (
	"context"
	"math/big"

	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) RetrieveMarketUpdates(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdate, error) {
	opts := s.getFilterOptsPerpsMarket(fromBlock, toBLock)

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
