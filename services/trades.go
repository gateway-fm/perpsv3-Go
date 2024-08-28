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

func (s *Service) RetrieveTrades(fromBlock uint64, toBLock *uint64) ([]*models.Trade, error) {
	opts := s.getFilterOptsPerpsMarket(fromBlock, toBLock)
	return s.retrieveTrades(opts)
}

func (s *Service) RetrieveTradesLimit(limit uint64) ([]*models.Trade, error) {
	iterations, last, err := s.getIterationsForLimitQueryPerpsMarket(limit)
	if err != nil {
		return nil, err
	}

	var trades []*models.Trade

	logger.Log().WithField("layer", "Service-RetrieveTradesLimit").Infof(
		"fetching orders updates with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.perpsMarketFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveTradesLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsPerpsMarket(fromBlock, &toBlock)

		res, err := s.retrieveTrades(opts)
		if err != nil {
			return nil, err
		}

		trades = append(trades, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveTradesLimit").Infof("task completed successfully")

	return trades, nil
}

// retrieveTrades is used to retrieve trades with given filter options
func (s *Service) retrieveTrades(opts *bind.FilterOpts) ([]*models.Trade, error) {
	iterator, err := s.perpsMarket.FilterOrderSettled(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveTrades").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "perps market")
	}

	var trades []*models.Trade

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveTrades").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "perps market")
		}
		blockNumber := iterator.Event.Raw.BlockNumber

		trade, err := s.getTrade(iterator.Event, blockNumber)
		if err != nil {
			return nil, err
		}

		trades = append(trades, trade)
	}

	return trades, nil
}

// getTrade is used to get models.Trade from given event and block number
func (s *Service) getTrade(event *perpsMarket.PerpsMarketOrderSettled, blockN uint64) (*models.Trade, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveTrades").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetTradeFromEvent(event, block.Time), nil
}
