package services

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

func (s *Service) FormatAccount(id *big.Int) (*models.Account, error) {
	return s.formatAccount(id)
}

func (s *Service) FormatAccountsLimit(limit uint64) ([]*models.Account, error) {
	iterations, last, err := s.getIterationsForLimitQuery(limit)
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
	iterations, last, err := s.getIterationsForLimitQuery(limit)
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
	margin, err := s.perpsMarket.GetAvailableMargin(nil, accountId)
	if err != nil {
		logger.Log().WithField("layer", "").Errorf("get avaliable margin error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetAvailableMargin")
	}

	return margin, nil
}

func (s *Service) GetRequiredMaintenanceMargin(accountId *big.Int) (*big.Int, error) {
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
