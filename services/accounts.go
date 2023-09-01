package services

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
	"math/big"
)

func (s *Service) GetAccount(id *big.Int) (*models.Account, error) {
	return s.getAccount(id)
}

func (s *Service) GetAccountsLimit(limit uint64) ([]*models.Account, error) {
	last, err := s.rpcClient.BlockNumber(context.Background())
	if err != nil {
		logger.Log().WithField("layer", "Service-GetAccountsLimit").Errorf("get latest block rpc error: %v", err.Error())
		return nil, errors.GetRPCProviderErr(err, "BlockNumber")
	}

	iterations := (last-s.perpsMarketFirstBlock)/limit + 1

	var accounts []*models.Account

	fromBlock := s.perpsMarketFirstBlock
	toBlock := fromBlock + limit
	for i := uint64(1); i <= iterations; i++ {
		opts := s.getFilterOptsPerpsMarket(fromBlock, &toBlock)

		res, err := s.getAccounts(opts)
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

	return accounts, nil
}

func (s *Service) GetAccounts() ([]*models.Account, error) {
	opts := s.getFilterOptsPerpsMarket(0, nil)

	return s.getAccounts(opts)
}

// getAccounts is used to get accounts from the contract using event filter function for 'AccountCreated' event
// and given filter options
func (s *Service) getAccounts(opts *bind.FilterOpts) ([]*models.Account, error) {
	iterator, err := s.perpsMarket.FilterAccountCreated(opts, nil, nil)
	if err != nil {
		logger.Log().WithField("layer", "Service-getAccounts").Errorf("error get iterator: %v", err.Error())
		return nil, errors.GetFilterErr(err, "perps market")
	}

	var accounts []*models.Account

	for iterator.Next() {
		if iterator.Error() != nil {
			logger.Log().WithField("layer", "Service-getAccounts").Errorf("iterator error: %v", err.Error())
			return nil, errors.GetFilterErr(iterator.Error(), "perps market")
		}

		account, err := s.getAccount(iterator.Event.AccountId)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

// getAccount is used to get models.Account data from given account id
func (s *Service) getAccount(id *big.Int) (*models.Account, error) {
	owner, err := s.perpsMarket.GetAccountOwner(nil, id)
	if err != nil {
		logger.Log().WithField("layer", "Service-getAccount").Errorf("get account owner error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetAccountOwner")
	}

	time, err := s.perpsMarket.GetAccountLastInteraction(nil, id)
	if err != nil {
		logger.Log().WithField("layer", "Service-getAccount").Errorf("get account last interaction error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetAccountLastInteraction")
	}

	permissions, err := s.perpsMarket.GetAccountPermissions(nil, id)
	if err != nil {
		logger.Log().WithField("layer", "Service-getAccount").Errorf("get account permissions error: %v", err.Error())
		return nil, errors.GetReadContractErr(err, "perps market", "GetAccountPermissions")
	}

	return models.GetAccount(id, owner, time.Uint64(), permissions), nil
}
