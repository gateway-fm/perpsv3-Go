package services

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gateway-fm/perpsv3-Go/config"
	"github.com/gateway-fm/perpsv3-Go/contracts/forwarder"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) FormatAccount(id *big.Int) (*models.Account, error) {
	return s.formatAccount(id)
}

func (s *Service) FormatAccountsLimit(limit uint64) ([]*models.Account, error) {
	iterations, last, err := s.getIterationsForLimitQueryPerpsMarket(limit)
	if err != nil {
		return nil, err
	}

	var accounts []*models.Account

	logger.Log().WithField("layer", "Service-FormatAccountsLimit").Infof(
		"fetching accounts with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.perpsMarketFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-FormatAccountsLimit").Infof("-- iteration %v", i)
		}

		opts := s.getFilterOptsPerpsMarket(fromBlock, &toBlock)

		res, err := s.formatAccounts(opts)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-FormatAccountsLimit").Infof("task completed successfully")

	return accounts, nil
}

func (s *Service) RetrieveAccountLiquidationsLimit(limit uint64) ([]*models.AccountLiquidated, error) {
	iterations, last, err := s.getIterationsForLimitQueryPerpsMarket(limit)
	if err != nil {
		return nil, err
	}

	var accountLiquidations []*models.AccountLiquidated

	logger.Log().WithField("layer", "Service-QueryAccountLiquidatedLimit").Infof(
		"fetching accounts with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.perpsMarketFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-QueryAccountLiquidatedLimit").Infof("-- iteration %v", i)
		}

		opts := s.getFilterOptsPerpsMarket(fromBlock, &toBlock)

		iterator, err := s.perpsMarket.FilterAccountLiquidationAttempt(opts, nil)
		if err != nil {
			logger.Log().WithField("layer", "Service-QueryAccountLiquidatedLimit").Errorf("error get iterator: %v", err.Error())
			return nil, errors.GetFilterErr(err, "perps market")
		}

		for iterator.Next() {
			if iterator.Error() != nil {
				logger.Log().WithField("layer", "Service-QueryAccountLiquidatedLimit").Errorf("iterator error: %v", err.Error())
				return nil, errors.GetFilterErr(iterator.Error(), "perps market")
			}

			accountLiquidations = append(accountLiquidations, &models.AccountLiquidated{
				ID:             iterator.Event.AccountId,
				Reward:         iterator.Event.Reward,
				FullLiquidated: iterator.Event.FullLiquidation,
			})
		}

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-QueryAccountLiquidatedLimit").Infof("task completed successfully")

	return accountLiquidations, nil
}

func (s *Service) FormatAccounts() ([]*models.Account, error) {
	opts := s.getFilterOptsPerpsMarket(0, nil)

	return s.formatAccounts(opts)
}

func (s *Service) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	return s.getAvailableMarginMulticallRetries(accountId, 0)
}

func (s *Service) getAvailableMarginMulticallRetries(accountId *big.Int, fails int) (res *big.Int, err error) {
	switch {
	case s.chainID == config.BaseAndromeda || s.chainID == config.BaseSepolia:
		res, err = s.getAvailableMarginMulticallNoPyth(accountId, true)
		if err != nil && fails <= s.multicallRetries {
			logger.Log().WithField("layer", "GetAvailableMargin").Debugf(
				"failed multicall no pyth, retry in %v total fails %v ", s.multicallWait, fails,
			)
			time.Sleep(s.multicallWait)
			return s.getAvailableMarginMulticallRetries(accountId, fails+1)
		}
	case s.chainID == config.BaseMainnet:
		res, err = s.getAvailableMarginMulticall(accountId, true)
		if err != nil && fails <= s.multicallRetries {
			logger.Log().WithField("layer", "GetAvailableMargin").Debugf(
				"failed multicall with pyth, retry in %v total fails %v ", s.multicallWait, fails,
			)
			time.Sleep(s.multicallWait)
			return s.getAvailableMarginMulticallRetries(accountId, fails+1)
		}
	default:
		res, err = s.getAvailableMargin(accountId)
	}

	return res, err
}

func (s *Service) getAvailableMarginMulticallNoPyth(accountId *big.Int, retry bool) (res *big.Int, err error) {
	getMarginCallData, err := s.rawPerpsContract.GetCallDataAvailableMargin(accountId)
	if err != nil {
		return res, err
	}

	callMargins := forwarder.TrustedMulticallForwarderCall3Value{
		Target:         s.rawPerpsContract.Address(),
		RequireSuccess: true,
		Value:          big.NewInt(0),
		CallData:       getMarginCallData,
	}

	call, err := s.rawForwarder.Aggregate3Value(0, []forwarder.TrustedMulticallForwarderCall3Value{callMargins})
	if err != nil {
		if retry {
			logger.Log().WithField("layer", "GetAvailableMargin").Debugf(
				"failed multicall with no pyth, retry with pyth",
			)
			return s.getAvailableMarginMulticall(accountId, false)
		}
		return res, err
	}

	if len(call) != 1 {
		logger.Log().WithField("layer", "getAvailableMarginMulticallNoPyth").Errorf("received %v from rawForwarder contract, expected 1", len(call))
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call"), "rawForwarder", "Aggregate3Value")
	}

	if !call[0].Success {
		logger.Log().WithField("layer", "getAvailableMarginMulticallNoPyth").Error("call to perps unsuccessful")
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to peprs"), "rawForwarder", "Aggregate3Value")
	}

	unpackedMargin, err := s.rawPerpsContract.UnpackAvailableMargin(call[0].ReturnData)
	if err != nil {
		return res, err
	}

	return unpackedMargin, nil
}

func (s *Service) getAvailableMarginMulticall(accountId *big.Int, retry bool) (res *big.Int, err error) {
	getMarginCallData, err := s.rawPerpsContract.GetCallDataAvailableMargin(accountId)
	if err != nil {
		return res, err
	}

	callMargins := forwarder.TrustedMulticallForwarderCall3Value{
		Target:         s.rawPerpsContract.Address(),
		RequireSuccess: true,
		Value:          big.NewInt(0),
		CallData:       getMarginCallData,
	}

	marketIDs, err := s.GetMarketIDs()
	if err != nil {
		return nil, err
	}

	feedIDs := []string{}
	for _, m := range marketIDs {
		feedID := models.GetPriceFeedIDFromMarketID(m)
		if feedID == models.UNKNOWN {
			logger.Log().WithField("layer", "getAvailableMarginMulticall").Errorf(
				"market ud: %v not supported on andromeda net", m.String(),
			)
			return res, errors.GetReadContractErr(fmt.Errorf("market %v not supported", m.String()), "rawForwarder", "Aggregate3Value")
		}

		feedIDs = append(feedIDs, feedID.String())
	}

	fulfillOracleQueryCallData, err := s.rawERC7412.GetCallFulfillOracleQueryAll(feedIDs)
	if err != nil {
		logger.Log().WithField("layer", "getAvailableMarginMulticall").Errorf(
			"err GetCallFulfillOracleQueryAll",
		)
		return res, err
	}

	callFulfill := forwarder.TrustedMulticallForwarderCall3Value{
		Target:         s.rawERC7412.Address(),
		RequireSuccess: true,
		Value:          big.NewInt(2),
		CallData:       fulfillOracleQueryCallData,
	}

	call, err := s.rawForwarder.Aggregate3Value(2, []forwarder.TrustedMulticallForwarderCall3Value{callFulfill, callMargins})
	if err != nil {
		if retry {
			logger.Log().WithField("layer", "GetAvailableMargin").Debugf(
				"failed multicall with pyth, retry with no pyth",
			)
			return s.getAvailableMarginMulticallNoPyth(accountId, false)
		}
		return res, err
	}

	if len(call) != 2 {
		logger.Log().WithField("layer", "getAvailableMarginMulticall").Errorf("received %v from rawForwarder contract, expected 2", len(call))
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call"), "rawForwarder", "Aggregate3Value")
	}

	if !call[0].Success {
		logger.Log().WithField("layer", "getAvailableMarginMulticall").Error("call to erc7412 unsuccessful")
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to erc7412"), "rawForwarder", "Aggregate3Value")
	}

	if !call[1].Success {
		logger.Log().WithField("layer", "getAvailableMarginMulticall").Error("call to perps unsuccessful")
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to perps market"), "rawForwarder", "Aggregate3Value")
	}

	unpackedMargin, err := s.rawPerpsContract.UnpackAvailableMargin(call[1].ReturnData)
	if err != nil {
		return res, err
	}

	return unpackedMargin, nil
}

func (s *Service) getAvailableMargin(accountId *big.Int) (*big.Int, error) {
	margin, err := s.perpsMarket.GetAvailableMargin(nil, accountId)
	if err != nil {
		logger.Log().WithField("layer", "").Errorf("get avaliable margin error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetAvailableMargin")
	}

	return margin, nil
}

func (s *Service) GetRequiredMaintenanceMargin(accountId *big.Int) (*big.Int, error) {
	return s.getRequiredMaintenanceMarginRetries(accountId, 0)
}

func (s *Service) getRequiredMaintenanceMarginRetries(accountId *big.Int, fails int) (res *big.Int, err error) {
	switch {
	case s.chainID == config.BaseAndromeda || s.chainID == config.BaseSepolia:
		res, err = s.getRequiredMaintenanceMarginMulticallNoPyth(accountId, true)
		if err != nil && fails <= s.multicallRetries {
			logger.Log().WithField("layer", "GetRequiredMaintenanceMargin").Debugf(
				"failed multicall no pyth, retry in %v total fails %v ", s.multicallWait, fails,
			)
			time.Sleep(s.multicallWait)
			return s.getRequiredMaintenanceMarginRetries(accountId, fails+1)
		}
	case s.chainID == config.BaseMainnet:
		res, err = s.getRequiredMaintenanceMarginMulticall(accountId, true)
		if err != nil && fails <= s.multicallRetries {
			logger.Log().WithField("layer", "GetRequiredMaintenanceMargin").Debugf(
				"failed multicall with pyth, retry in %v total fails %v ", s.multicallWait, fails,
			)
			time.Sleep(s.multicallWait)
			return s.getRequiredMaintenanceMarginRetries(accountId, fails+1)
		}
	default:
		res, err = s.getRequiredMaintenanceMargin(accountId)
	}

	return res, err
}

func (s *Service) getRequiredMaintenanceMarginMulticallNoPyth(accountId *big.Int, retries bool) (res *big.Int, err error) {
	getMarginsCallData, err := s.rawPerpsContract.GetCallDataRequiredMargins(accountId)
	if err != nil {
		return res, err
	}

	callMargins := forwarder.TrustedMulticallForwarderCall3Value{
		Target:         s.rawPerpsContract.Address(),
		RequireSuccess: true,
		Value:          big.NewInt(0),
		CallData:       getMarginsCallData,
	}

	call, err := s.rawForwarder.Aggregate3Value(0, []forwarder.TrustedMulticallForwarderCall3Value{callMargins})
	if err != nil {
		if retries {
			logger.Log().WithField("layer", "GetRequiredMaintenanceMargin").Debugf(
				"failed multicall with no pyth, retry with pyth",
			)
			return s.getRequiredMaintenanceMarginMulticall(accountId, false)
		}
		return res, err
	}

	if len(call) != 1 {
		logger.Log().WithField("layer", "getRequiredMaintenanceMarginMulticallNoPyth").Errorf("received %v from rawForwarder contract, expected 2", len(call))
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call"), "rawForwarder", "Aggregate3Value")
	}

	if !call[0].Success {
		logger.Log().WithField("layer", "getRequiredMaintenanceMarginMulticallNoPyth").Error("call to erc7412 unsuccessful")
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to perps"), "rawForwarder", "Aggregate3Value")
	}

	unpackedMargins, err := s.rawPerpsContract.UnpackRequiredMargins(call[0].ReturnData)
	if err != nil {
		return res, err
	}

	return unpackedMargins.RequiredMaintenanceMargin, nil
}

func (s *Service) getRequiredMaintenanceMarginMulticall(accountId *big.Int, retries bool) (res *big.Int, err error) {
	getMarginsCallData, err := s.rawPerpsContract.GetCallDataRequiredMargins(accountId)
	if err != nil {
		return res, err
	}

	callMargins := forwarder.TrustedMulticallForwarderCall3Value{
		Target:         s.rawPerpsContract.Address(),
		RequireSuccess: true,
		Value:          big.NewInt(0),
		CallData:       getMarginsCallData,
	}

	marketIDs, err := s.GetMarketIDs()
	if err != nil {
		return nil, err
	}

	feedIDs := []string{}
	for _, m := range marketIDs {
		feedID := models.GetPriceFeedIDFromMarketID(m)
		if feedID == models.UNKNOWN {
			logger.Log().WithField("layer", "getRequiredMaintenanceMarginMulticall").Errorf(
				"market ud: %v not supported on andromeda net", m.String(),
			)
			return res, errors.GetReadContractErr(fmt.Errorf("market %v not supported", m.String()), "rawForwarder", "Aggregate3Value")
		}

		feedIDs = append(feedIDs, feedID.String())
	}

	fulfillOracleQueryCallData, err := s.rawERC7412.GetCallFulfillOracleQueryAll(feedIDs)
	if err != nil {
		logger.Log().WithField("layer", "getRequiredMaintenanceMarginMulticall").Errorf(
			"err GetCallFulfillOracleQueryAll",
		)
		return res, err
	}

	callFulfill := forwarder.TrustedMulticallForwarderCall3Value{
		Target:         s.rawERC7412.Address(),
		RequireSuccess: true,
		Value:          big.NewInt(2),
		CallData:       fulfillOracleQueryCallData,
	}

	call, err := s.rawForwarder.Aggregate3Value(2, []forwarder.TrustedMulticallForwarderCall3Value{callFulfill, callMargins})
	if err != nil {
		if retries {
			logger.Log().WithField("layer", "GetRequiredMaintenanceMargin").Debugf(
				"failed multicall with pyth, retry with no pyth",
			)
			return s.getRequiredMaintenanceMarginMulticallNoPyth(accountId, false)
		}
		return res, err
	}

	if len(call) != 2 {
		logger.Log().WithField("layer", "getRequiredMaintenanceMarginMulticall").Errorf("received %v from rawForwarder contract, expected 2", len(call))
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call"), "rawForwarder", "Aggregate3Value")
	}

	if !call[0].Success {
		logger.Log().WithField("layer", "getRequiredMaintenanceMarginMulticall").Error("call to erc7412 unsuccessful")
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to erc7412"), "rawForwarder", "Aggregate3Value")
	}

	if !call[1].Success {
		logger.Log().WithField("layer", "getRequiredMaintenanceMarginMulticall").Error("call to perps unsuccessful")
		return res, errors.GetReadContractErr(fmt.Errorf("invalid call to perps market"), "rawForwarder", "Aggregate3Value")
	}

	unpackedMargins, err := s.rawPerpsContract.UnpackRequiredMargins(call[1].ReturnData)
	if err != nil {
		return res, err
	}

	return unpackedMargins.RequiredMaintenanceMargin, nil
}

func (s *Service) getRequiredMaintenanceMargin(accountId *big.Int) (*big.Int, error) {
	requiredMargins, err := s.perpsMarket.GetRequiredMargins(nil, accountId)
	if err != nil {
		logger.Log().WithField("layer", "").Errorf("get required margins error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetRequiredMargins")
	} else if requiredMargins.RequiredMaintenanceMargin == nil {
		logger.Log().WithField("layer", "").Errorf("get required margins error: MaintenanceMargin = nil")
		return nil, errors.GetReadContractErr(fmt.Errorf("required margins error: MaintenanceMargin = nil"),
			"perps market", "GetRequiredMargins")
	}

	return requiredMargins.RequiredMaintenanceMargin, nil
}

func (s *Service) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	time, err := s.perpsMarket.GetAccountLastInteraction(nil, accountId)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetAccountLastInteraction").Errorf("get account last interaction error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetAccountLastInteraction")
	}

	return time, nil
}

func (s *Service) GetAccountOwner(accountId *big.Int) (string, error) {
	owner, err := s.perpsMarket.GetAccountOwner(nil, accountId)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetAccountOwner").Errorf("get account owner error: %v", err.Error())
		return "", errors.GetReadContractErr(err, "perps market", "GetAccountOwner")
	}

	return owner.Hex(), nil
}

func (s *Service) GetCollateralAmount(accountId *big.Int, marketId *big.Int) (*big.Int, error) {
	amount, err := s.perpsMarket.GetCollateralAmount(nil, accountId, marketId)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetCollateralAmount").Errorf("get colleteral amount error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetCollateralAmount")
	}

	return amount, nil
}

// formatAccounts is used to get accounts from the contract using event filter function for 'AccountCreated' event
// and given filter options
func (s *Service) formatAccounts(opts *bind.FilterOpts) ([]*models.Account, error) {
	iterator, err := s.perpsMarket.FilterAccountCreated(opts, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-formatAccounts").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "perps market")
	}

	var accounts []*models.Account

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-formatAccounts").Errorf("iterator error: %v", err.Error())
			return nil, errors.GetFilterErr(iterator.Error(), "perps market")
		}

		account, err := s.formatAccount(iterator.Event.AccountId)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *Service) formatAccountsCore(opts *bind.FilterOpts) ([]*models.Account, error) {
	iterator, err := s.core.FilterAccountCreated(opts, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-formatAccountsCore").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "core")
	}

	var accounts []*models.Account

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-formatAccountsCore").Errorf("iterator error: %v", err.Error())
			return nil, errors.GetFilterErr(iterator.Error(), "core")
		}

		account, err := s.formatAccountCore(iterator.Event.AccountId)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

// formatAccount is used to get models.Account data from given account id
func (s *Service) formatAccount(id *big.Int) (*models.Account, error) {
	owner, err := s.perpsMarket.GetAccountOwner(nil, id)
	if err != nil {
		logger.Log().WithField("layer", "Service-formatAccount").Errorf("get account owner error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetAccountOwner")
	}

	time, err := s.perpsMarket.GetAccountLastInteraction(nil, id)
	if err != nil {
		logger.Log().WithField("layer", "Service-formatAccount").Errorf("get account last interaction error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetAccountLastInteraction")
	}

	permissions, err := s.perpsMarket.GetAccountPermissions(nil, id)
	if err != nil {
		logger.Log().WithField("layer", "Service-formatAccount").Errorf("get account permissions error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetAccountPermissions")
	}

	return models.FormatAccount(id, owner, time.Uint64(), permissions), nil
}

func (s *Service) FormatAccountsCoreLimit(limit uint64) ([]*models.Account, error) {
	iterations, last, err := s.getIterationsForLimitQueryCore(limit)
	if err != nil {
		return nil, err
	}

	var accounts []*models.Account

	logger.Log().WithField("layer", "Service-FormatAccountsCoreLimit").Infof(
		"fetching accounts with limit: %v to block: %v total iterations: %v...",
		limit, last, iterations,
	)

	fromBlock := s.coreFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		if i%10 == 0 || i == iterations {
			logger.Log().WithField("layer", "Service-FormatAccountsCoreLimit").Infof("-- iteration %v", i)
		}

		opts := s.getFilterOptsCore(fromBlock, &toBlock)

		res, err := s.formatAccountsCore(opts)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, res...)

		fromBlock = toBlock + 1

		if i == iterations-1 {
			toBlock = last
		} else {
			toBlock = fromBlock + limit
		}
	}

	logger.Log().WithField("layer", "Service-FormatAccountsLimit").Infof("task completed successfully")

	return accounts, nil
}

func (s *Service) FormatAccountCore(id *big.Int) (*models.Account, error) {
	return s.formatAccountCore(id)
}

func (s *Service) formatAccountCore(id *big.Int) (*models.Account, error) {
	owner, err := s.core.GetAccountOwner(nil, id)
	if err != nil {
		logger.Log().WithField("layer", "Service-formatAccount").Errorf("get account owner error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "core", "GetAccountOwner")
	}

	time, err := s.core.GetAccountLastInteraction(nil, id)
	if err != nil {
		logger.Log().WithField("layer", "Service-formatAccount").Errorf("get account last interaction error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "core", "GetAccountLastInteraction")
	}

	permissions, err := s.core.GetAccountPermissions(nil, id)
	if err != nil {
		logger.Log().WithField("layer", "Service-formatAccount").Errorf("get account permissions error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "core", "GetAccountPermissions")
	}

	return models.FormatAccountCore(id, owner, time.Uint64(), permissions), nil
}

func (s *Service) GetAccountCollateralCore(accountId *big.Int, collateralType common.Address) (*models.AccountCollateral, error) {
	res, err := s.core.GetAccountCollateral(nil, accountId, collateralType)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetAccountCollateralCore").Errorf("get account collateral error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "core", "GetAccountCollateral")
	}

	return models.GetAccountCollateralFromContract(res), nil
}

func (s *Service) GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	res, err := s.core.GetAccountAvailableCollateral(nil, accountId, collateralType)
	if err != nil {
		logger.Log().WithField("layer", "Service-GetAccountAvailableCollateral").Errorf("get account available collateral error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "core", "GetAccountAvailableCollateral")
	}

	return res, nil
}
