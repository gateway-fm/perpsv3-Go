package services

import (
	"context"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) RetrieveOrders(fromBlock uint64, toBLock *uint64) ([]*models.Order, error) {
	opts := s.getFilterOptsPerpsMarket(fromBlock, toBLock)
	return s.retrieveOrders(opts)
}

func (s *Service) RetrieveOrdersLimit(limit uint64) ([]*models.Order, error) {
	iterations, last, err := s.getIterationsForLimitQueryPerpsMarket(limit)
	if err != nil {
		return nil, err
	}

	var orders []*models.Order

	logger.Log().WithField("layer", "Service-RetrieveOrdersLimit").Infof(
		"fetching orders updates with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.perpsMarketFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveOrdersLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsPerpsMarket(fromBlock, &toBlock)

		res, err := s.retrieveOrders(opts)
		if err != nil {
			return nil, err
		}

		orders = append(orders, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveOrdersLimit").Infof("task completed successfully")

	return orders, nil
}

// retrieveOrders is used to retrieve orders with given filter options
func (s *Service) retrieveOrders(opts *bind.FilterOpts) ([]*models.Order, error) {
	iterator, err := s.perpsMarket.FilterOrderCommitted(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveOrders").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "perps market")
	}

	var orders []*models.Order

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveOrders").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "perps market")
		}

		order, err := s.getOrder(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

// getOrder is used to get models.Order from given event and block number
func (s *Service) getOrder(event *perpsMarket.PerpsMarketOrderCommitted, blockN uint64) (*models.Order, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveTrades").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetOrderFromEvent(event, block.Time), nil
}
