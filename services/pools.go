package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/contracts/forwarder"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) RetrieveUSDMintedLimit(limit uint64) ([]*models.USDMinted, error) {
	iterations, last, err := s.getIterationsForLimitQuery(limit)
	if err != nil {
		return nil, err
	}

	var mints []*models.USDMinted

	logger.Log().WithField("layer", "Service-RetrieveUSDMintedLimit").Infof(
		"fetching USDMinted with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveUSDMintedLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveUSDMinted(opts)
		if err != nil {
			return nil, err
		}

		mints = append(mints, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsLimit").Infof("task completed successfully")

	return mints, nil
}

func (s *Service) RetrieveUSDBurnedLimit(limit uint64) ([]*models.USDBurned, error) {
	iterations, last, err := s.getIterationsForLimitQuery(limit)
	if err != nil {
		return nil, err
	}

	var burns []*models.USDBurned

	logger.Log().WithField("layer", "Service-RetrieveUSDBurnedLimit").Infof(
		"fetching USDMBurned with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveUSDMBurnedLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveUSDBurned(opts)
		if err != nil {
			return nil, err
		}

		burns = append(burns, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsLimit").Infof("task completed successfully")

	return burns, nil
}

func (s *Service) RetrieveDelegationUpdatedLimit(limit uint64) ([]*models.DelegationUpdated, error) {
	iterations, last, err := s.getIterationsForLimitQuery(limit)
	if err != nil {
		return nil, err
	}

	var delegations []*models.DelegationUpdated

	logger.Log().WithField("layer", "Service-RetrieveDelegationUpdatedLimit").Infof(
		"fetching DelegationUpdated with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-RetrieveUSDMBurnedLimit").Infof("-- iteration %v", i)
		}
		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.retrieveDelegationUpdated(opts)
		if err != nil {
			return nil, err
		}

		delegations = append(delegations, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-RetrieveLiquidationsLimit").Infof("task completed successfully")

	return delegations, nil
}

func (s *Service) retrieveDelegationUpdated(opts *bind.FilterOpts) ([]*models.DelegationUpdated, error) {
	iterator, err := s.core.FilterDelegationUpdated(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveDelegationUpdatedLimit").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var delegations []*models.DelegationUpdated

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveDelegationUpdatedLimit").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		mint, err := s.getDelegationUpdated(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		delegations = append(delegations, mint)
	}

	return delegations, nil
}

func (s *Service) retrieveUSDBurned(opts *bind.FilterOpts) ([]*models.USDBurned, error) {
	iterator, err := s.core.FilterUsdBurned(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveUSDBurnedLimit").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var mints []*models.USDBurned

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveUSDBurnedLimit").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		mint, err := s.getUSDBurned(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		mints = append(mints, mint)
	}

	return mints, nil
}

func (s *Service) retrieveUSDMinted(opts *bind.FilterOpts) ([]*models.USDMinted, error) {
	iterator, err := s.core.FilterUsdMinted(opts, nil, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveUSDMintedLimit").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var mints []*models.USDMinted

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-RetrieveUSDMintedLimit").Errorf("iterator error: %v", iterator.Error().Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		mint, err := s.getUSDMinted(iterator.Event, iterator.Event.Raw.BlockNumber)
		if err != nil {
			return nil, err
		}

		mints = append(mints, mint)
	}

	return mints, nil
}

// getLiquidation is used to get models.Liquidation from given event and block number
func (s *Service) getUSDMinted(event *core.CoreUsdMinted, blockN uint64) (*models.USDMinted, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveUSDMintedLimit").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetUSDMintedFromEvent(event, block.Time), nil
}

func (s *Service) getUSDBurned(event *core.CoreUsdBurned, blockN uint64) (*models.USDBurned, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveUSDBurnedLimit").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetUSDBurnedFromEvent(event, block.Time), nil
}

func (s *Service) getDelegationUpdated(event *core.CoreDelegationUpdated, blockN uint64) (*models.DelegationUpdated, error) {
	block, err := s.rpcClient.HeaderByNumber(context.Background(), big.NewInt(int64(blockN)))
	if err != nil {
		logger.Log().WithField("layer", "Service-RetrieveDelegationUpdatedLimit").Errorf(
			"get block:%v by number error: %v", blockN, err.Error(),
		)
		return nil, errors.GetRPCProviderErr(err, "HeaderByNumber")
	}

	return models.GetDelegationUpdatedFromEvent(event, block.Time), nil
}

func (s *Service) GetVaultCollateral(poolID *big.Int, collateralType common.Address) (amount *big.Int, value *big.Int, err error) {
	res, err := s.core.GetVaultCollateral(nil, poolID, collateralType)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetVaultCollateral").Errorf("error from the contract: %v", err.Error())
		return nil, nil, errors.GetReadContractErr(err, "core", "getVaultCollateral")
	}

	return res.Amount, res.Value, nil
}

func (s *Service) GetVaultDebt(poolID *big.Int, collateralType common.Address) (*big.Int, error) {
	if poolID == nil {
		logger.Log().WithField("layer", "Service-GetGetVaultDebt").Errorf("received nil pool id")
		return nil, errors.GetInvalidArgumentErr("pool id cannot be nil")
	}
	return s.getVaultDebtRetries(poolID, collateralType, s.multicallRetries)
}

func (s *Service) getVaultDebtRetries(poolID *big.Int, collateralType common.Address, fails int) (res *big.Int, err error) {
	res, err = s.getVaultDebtMultiCallNoPyth(poolID, collateralType)
	if err != nil && fails <= s.multicallRetries {
		time.Sleep(s.multicallWait)
		return s.getVaultDebtRetries(poolID, collateralType, fails+1)
	}

	return res, err
}

func (s *Service) getVaultDebtMultiCallNoPyth(poolID *big.Int, collateralType common.Address) (res *big.Int, err error) {
	getVaultDebtCallData, err := s.rawCore.GetCallDataVaultDebt(poolID, collateralType)
	if err != nil {
		return res, err
	}

	callVaultDebt := forwarder.TrustedMulticallForwarderCall3Value{
		Target:         s.rawCore.Address(),
		RequireSuccess: true,
		Value:          big.NewInt(0),
		CallData:       getVaultDebtCallData,
	}

	call, err := s.rawForwarder.Aggregate3Value(0, []forwarder.TrustedMulticallForwarderCall3Value{callVaultDebt})
	if err != nil {
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

	unpackedDebt, err := s.rawCore.UnpackVaultDebt(call[0].ReturnData)
	if err != nil {
		return res, err
	}

	return unpackedDebt, nil
}
