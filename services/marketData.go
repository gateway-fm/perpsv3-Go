package services

import (
	"context"
	"fmt"
	"github.com/gateway-fm/perpsv3-Go/config"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/contracts/forwarder"
	"math/big"
	"time"

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

func (s *Service) RetrieveMarketRegistered(limit uint64) ([]*models.MarketRegistered, error) {
	iterations, last, err := s.getIterationsForLimitQuery(limit)
	if err != nil {
		return nil, err
	}

	var marketRegistrations []*models.MarketRegistered

	logger.Log().WithField("layer", "Service-RetrieveMarketRegistered").Infof(
		"fetching market registrations with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveMarketRegistered").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveMarketRegistrations(opts)
		if err != nil {
			return nil, err
		}

		marketRegistrations = append(marketRegistrations, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveMarketRegistered").Infof("task completed successfully")

	return marketRegistrations, nil
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

func (s *Service) getMarketSummaryRetries(marketID *big.Int, fails int) (res perpsMarket.IPerpsMarketModuleMarketSummary, err error) {
	switch {
	case s.chainID == config.BaseAndromeda || s.chainID == config.BaseSepolia:
		res, err = s.getMarketSummaryMultiCallNoPyth(marketID, true)
		if err != nil && fails <= s.multicallRetries {
			time.Sleep(s.multicallWait)
			return s.getMarketSummaryRetries(marketID, fails+1)
		}
	case s.chainID == config.BaseMainnet:
		res, err = s.getMarketSummaryMultiCall(marketID, true)
		if err != nil && fails <= s.multicallRetries {
			time.Sleep(s.multicallWait)
			return s.getMarketSummaryRetries(marketID, fails+1)
		}
	default:
		res, err = s.getMarketSummary(marketID)
	}

	return res, err
}

func (s *Service) GetMarketSummary(marketID *big.Int) (*models.MarketSummary, error) {
	if marketID == nil {
		logger.Log().WithField("layer", "Service-GetMarketSummary").Errorf("received nil market id")
		return nil, errors.GetInvalidArgumentErr("market id cannot be nil")
	}

	res, err := s.getMarketSummaryRetries(marketID, s.multicallRetries)
	if err != nil {
		return nil, err
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

// getMarketSummaryMultiCallNoPyth is used to get market summary using Forwarder contract with no erc7412 wrapper
func (s *Service) getMarketSummaryMultiCallNoPyth(marketID *big.Int, retry bool) (res perpsMarket.IPerpsMarketModuleMarketSummary, err error) {
	getMarketSummaryCallData, err := s.rawPerpsContract.GetCallDataMarketSummary(marketID)
	if err != nil {
		return res, err
	}

	callSummary := forwarder.TrustedMulticallForwarderCall3Value{
		Target:         s.rawPerpsContract.Address(),
		RequireSuccess: true,
		Value:          big.NewInt(0),
		CallData:       getMarketSummaryCallData,
	}

	call, err := s.rawForwarder.Aggregate3Value(0, []forwarder.TrustedMulticallForwarderCall3Value{callSummary})
	if err != nil {
		if retry {
			return s.getMarketSummaryMultiCall(marketID, false)
		}

		return res, err
	}

	if len(call) != 1 {
		logger.Log().WithField("layer", "getMarketSummaryMultiCallNoPyth").Errorf("received %v from rawForwarder contract, expected 1", len(call))
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call"), "rawForwarder", "Aggregate3Value")
	}

	if !call[0].Success {
		logger.Log().WithField("layer", "getMarketSummaryMultiCallNoPyth").Error("call to perps unsuccessful")
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to perps"), "rawForwarder", "Aggregate3Value")
	}

	unpackedSummary, err := s.rawPerpsContract.UnpackGetMarketSummary(call[0].ReturnData)
	if err != nil {
		return res, err
	}

	return *unpackedSummary, nil
}

// getMarketSummary is used to get market summary using Forwarder contract
func (s *Service) getMarketSummaryMultiCall(marketID *big.Int, retry bool) (res perpsMarket.IPerpsMarketModuleMarketSummary, err error) {
	getMarketSummaryCallData, err := s.rawPerpsContract.GetCallDataMarketSummary(marketID)
	if err != nil {
		return res, err
	}

	callSummary := forwarder.TrustedMulticallForwarderCall3Value{
		Target:         s.rawPerpsContract.Address(),
		RequireSuccess: true,
		Value:          big.NewInt(0),
		CallData:       getMarketSummaryCallData,
	}

	feedID := models.GetPriceFeedIDFromMarketID(marketID)
	if feedID == models.UNKNOWN {
		logger.Log().WithField("layer", "getMarketSummaryMultiCall").Errorf(
			"market ud: %v not supported on andromeda net", marketID.String(),
		)
		return res, errors.GetReadContractErr(fmt.Errorf("market %v not supported", marketID.String()), "rawForwarder", "Aggregate3Value")
	}

	fulfillOracleQueryCallData, err := s.rawERC7412.GetCallFulfillOracleQuery(feedID.String())
	if err != nil {
		return res, err
	}

	callFulfill := forwarder.TrustedMulticallForwarderCall3Value{
		Target:         s.rawERC7412.Address(),
		RequireSuccess: true,
		Value:          big.NewInt(1),
		CallData:       fulfillOracleQueryCallData,
	}

	call, err := s.rawForwarder.Aggregate3Value(1, []forwarder.TrustedMulticallForwarderCall3Value{callFulfill, callSummary})
	if err != nil {
		if retry {
			return s.getMarketSummaryMultiCallNoPyth(marketID, false)
		}

		return res, err
	}

	if len(call) != 2 {
		logger.Log().WithField("layer", "getMarketSummaryMultiCall").Errorf("received %v from rawForwarder contract, expected 2", len(call))
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call"), "rawForwarder", "Aggregate3Value")
	}

	if !call[0].Success {
		logger.Log().WithField("layer", "getMarketSummaryMultiCall").Error("call to erc7412 unsuccessful")
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to erc7412"), "rawForwarder", "Aggregate3Value")
	}

	if !call[1].Success {
		logger.Log().WithField("layer", "getMarketSummaryMultiCall").Error("call to erc7412 unsuccessful")
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to perps market"), "rawForwarder", "Aggregate3Value")
	}

	unpackedSummary, err := s.rawPerpsContract.UnpackGetMarketSummary(call[1].ReturnData)
	if err != nil {
		return res, err
	}

	return *unpackedSummary, nil
}

// getMarketSummary is used to get market summary straight from the perps contract
func (s *Service) getMarketSummary(marketID *big.Int) (res perpsMarket.IPerpsMarketModuleMarketSummary, err error) {
	res, err = s.perpsMarket.GetMarketSummary(nil, marketID)
	if err != nil {
		if err.Error() == "execution reverted" {
			logger.Log().WithField("layer", "Service-GetMarketSummary").Errorf("contract error, market does not exist")
			err = errors.GetInvalidArgumentErr("market does not exist")
		} else {
			logger.Log().WithField("layer", "Service-GetMarketSummary").Errorf("error from the contract: %v", err.Error())
			err = errors.GetReadContractErr(err, "perpsMarket", "getMarketSummary")
		}
	}

	return res, err
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

func (s *Service) retrieveMarketRegistrations(opts *bind.FilterOpts) ([]*models.MarketRegistered, error) {
	iterator, err := s.core.FilterMarketRegistered(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-retrieveMarketRegistrations").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var marketRegistrations []*models.MarketRegistered

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-retrieveMarketRegistrations").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		marketUpdate, err := s.getMarketRegistration(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		marketRegistrations = append(marketRegistrations, marketUpdate)
	}

	return marketRegistrations, nil
}

// getMarketUpdate is used to get models.MarketUpdate from given event and block number
func (s *Service) getMarketRegistration(event *core.CoreMarketRegistered, blockN uint64) (*models.MarketRegistered, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-getMarketRegistration").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetMarketRegisteredFromEvent(event, block.Time), nil
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
