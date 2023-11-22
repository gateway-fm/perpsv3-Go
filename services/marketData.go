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

func (s *Service) RetrieveMarketUpdatesLimit(limit uint64) ([]*models.MarketUpdate, error) {
	iterations, last, err := s.getIterationsForLimitQuery(limit)
	if err != nil {
		return nil, err
	}

	var marketUpdates []*models.MarketUpdate

	logger.Log().WithField("layer", "Service-RetrieveMarketUpdatesLimit").Infof(
		"fetching market updates with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.perpsMarketFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveMarketUpdatesLimit").Infof("-- iteration %v", i)
		}
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

	logger.Log().WithField("layer", "Service-RetrieveMarketUpdatesLimit").Infof("task completed successfully")

	return marketUpdates, nil
}

func (s *Service) RetrieveMarketUpdatesBigLimit(limit uint64) ([]*models.MarketUpdateBig, error) {
	iterations, last, err := s.getIterationsForLimitQuery(limit)
	if err != nil {
		return nil, err
	}

	var marketUpdates []*models.MarketUpdateBig

	logger.Log().WithField("layer", "Service-RetrieveMarketUpdatesBigLimit").Infof(
		"fetching market updates with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.perpsMarketFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveMarketUpdatesBigLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsPerpsMarket(fromBlock, &toBlock)

		res, err := s.retrieveMarketUpdatesBig(opts)
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

	logger.Log().WithField("layer", "Service-RetrieveMarketUpdatesBigLimit").Infof("task completed successfully")

	return marketUpdates, nil
}

func (s *Service) RetrieveMarketUpdates(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdate, error) {
	opts := s.getFilterOptsPerpsMarket(fromBlock, toBLock)
	return s.retrieveMarketUpdates(opts)
}

func (s *Service) RetrieveMarketUpdatesBig(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdateBig, error) {
	opts := s.getFilterOptsPerpsMarket(fromBlock, toBLock)
	return s.retrieveMarketUpdatesBig(opts)
}

func (s *Service) GetMarketMetadata(marketID *big.Int) (*models.MarketMetadata, error) {
	if marketID == nil {
		logger.Log().WithField("layer", "Service-GetMarketMetadata").Errorf("received nil market id")
		return nil, errors.GetInvalidArgumentErr("market id cannot be nil")
	}
	res, err := s.perpsMarket.Metadata(nil, marketID)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetMarketMetadata").Errorf("error from the contract: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perpsMarket", "metadata")
	}

	if res.Name == "" || res.Symbol == "" {
		logger.Log().WithField("layer", "Service-GetMarketMetadata").Errorf("received blank name and symbol from the contract")
		return nil, errors.GetInvalidArgumentErr("market id does not exist")
	}

	return models.GetMarketMetadataFromContractResponse(marketID, res.Name, res.Symbol), nil
}

func (s *Service) GetMarketSummary(marketID *big.Int) (*models.MarketSummary, error) {
	if marketID == nil {
		logger.Log().WithField("layer", "Service-GetMarketSummary").Errorf("received nil market id")
		return nil, errors.GetInvalidArgumentErr("market id cannot be nil")
	}

	res, err := s.perpsMarket.GetMarketSummary(nil, marketID)
	if err != nil {
		if err.Error() == "execution reverted" {
			logger.Log().WithField("layer", "Service-GetMarketSummary").Errorf("contract error, market does not exist")
			return nil, errors.GetInvalidArgumentErr("market does not exist")
		} else {
			logger.Log().WithField("layer", "Service-GetMarketSummary").Errorf("error from the contract: %v", err.Error())
			return nil, errors.GetReadContractErr(err, "perpsMarket", "getMarketSummary")
		}
	}

	block, err := s.rpcClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetMarketSummary").Errorf(
			"get latest block error: %v", err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetMarketSummaryFromContractModel(res, marketID, block.Time), nil
}

func (s *Service) GetMarketIDs() ([]*big.Int, error) {
	res, err := s.perpsMarket.GetMarkets(nil)
	if err != nil {
		return nil, errors.GetReadContractErr(err, "perpsMarket", "getMarkets")
	}

	return res, nil
}

func (s *Service) GetLiquidationParameters(marketId *big.Int) (*models.LiquidationParameters, error) {
	resp, err := s.perpsMarket.GetLiquidationParameters(nil, marketId)
	if err != nil {
		logger.Log().WithField("layer", "").Errorf("")
		return nil, errors.GetReadContractErr(err, "", "")
	}

	return models.GetLiquidationParameters(resp), nil
}

func (s *Service) GetFundingParameters(marketId *big.Int) (*models.FundingParameters, error) {
	resp, err := s.perpsMarket.GetFundingParameters(nil, marketId)
	if err != nil {
		logger.Log().WithField("layer", "").Errorf("")
		return nil, errors.GetReadContractErr(err, "", "")
	}

	return models.GetFundingParameters(resp), nil
}

func (s *Service) GetFoundingRate(marketId *big.Int) (*big.Int, error) {
	rate, err := s.perpsMarket.CurrentFundingRate(nil, marketId)
	if err != nil {
		return nil, errors.GetReadContractErr(err, "perpsMarket", "currentFoundingRate")
	}

	return rate, nil
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

// retrieveMarketUpdates is used to get retrieve market updates with given filter options
func (s *Service) retrieveMarketUpdatesBig(opts *bind.FilterOpts) ([]*models.MarketUpdateBig, error) {
	iterator, err := s.perpsMarket.FilterMarketUpdated(opts)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveMarketUpdates").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "perps market")
	}

	var marketUpdates []*models.MarketUpdateBig

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveMarketUpdates").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "perps market")
		}

		marketUpdate, err := s.getMarketUpdateBig(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		marketUpdates = append(marketUpdates, marketUpdate)
	}

	return marketUpdates, nil
}

// getMarketUpdate is used to get models.MarketUpdate from given event and block number
func (s *Service) getMarketUpdate(event *perpsMarket.PerpsMarketMarketUpdated, blockN uint64) (*models.MarketUpdate, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-getMarketUpdate").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetMarketUpdateFromEvent(event, block.Time), nil
}

// getMarketUpdate is used to get models.MarketUpdate from given event and block number
func (s *Service) getMarketUpdateBig(event *perpsMarket.PerpsMarketMarketUpdated, blockN uint64) (*models.MarketUpdateBig, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-getMarketUpdate").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetMarketUpdateBigFromEvent(event, block.Time), nil
}
