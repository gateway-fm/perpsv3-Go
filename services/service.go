package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/models"
)

// IService is a service layer interface
type IService interface {
	// RetrieveTrades is used to get logs from the "OrderSettled" event preps market contract within given block range
	RetrieveTrades(fromBlock uint64, toBLock *uint64) ([]*models.Trade, error)

	// RetrieveOrders is used to get logs from the "OrderCommitted" event preps market contract within given block range
	RetrieveOrders(fromBlock uint64, toBLock *uint64) ([]*models.Order, error)

	// RetrieveMarketUpdates is used to get logs from the "MarketUpdated" event preps market contract within given block
	// range
	RetrieveMarketUpdates(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdate, error)

	// GetPosition is used to get "Position" data struct from the latest block from the perps market with given data
	GetPosition(accountID *big.Int, marketID *big.Int) (*models.Position, error)

	// FormatAccount is used to get account, and it's additional data from the contract by given account id
	FormatAccount(id *big.Int) (*models.Account, error)

	// FormatAccounts is used to get all accounts and their additional data from the contract
	FormatAccounts() ([]*models.Account, error)

	// FormatAccountsLimit is used to get all accounts and their additional data from the contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	FormatAccountsLimit(limit uint64) ([]*models.Account, error)
}

// Service is an implementation of IService interface
type Service struct {
	rpcClient             *ethclient.Client
	core                  *coreGoerli.CoreGoerli
	coreFirstBlock        uint64
	spotMarket            *spotMarketGoerli.SpotMarketGoerli
	spotMarketFirstBlock  uint64
	perpsMarket           *perpsMarketGoerli.PerpsMarketGoerli
	perpsMarketFirstBlock uint64
}

// NewService is used to get instance of Service
func NewService(
	rpc *ethclient.Client,
	core *coreGoerli.CoreGoerli,
	coreFirstBlock uint64,
	spotMarket *spotMarketGoerli.SpotMarketGoerli,
	spotMarketFirstBlock uint64,
	perpsMarket *perpsMarketGoerli.PerpsMarketGoerli,
	perpsMarketFirstBlock uint64,
) IService {
	return &Service{
		rpcClient:             rpc,
		core:                  core,
		coreFirstBlock:        coreFirstBlock,
		spotMarket:            spotMarket,
		spotMarketFirstBlock:  spotMarketFirstBlock,
		perpsMarket:           perpsMarket,
		perpsMarketFirstBlock: perpsMarketFirstBlock,
	}
}

// getFilterOptsPerpsMarket is used to get options for event filtering on perps market contract
func (s *Service) getFilterOptsPerpsMarket(fromBlock uint64, toBLock *uint64) *bind.FilterOpts {
	if fromBlock == 0 {
		fromBlock = s.perpsMarketFirstBlock
	}

	return &bind.FilterOpts{
		Start:   fromBlock,
		End:     toBLock,
		Context: context.Background(),
	}
}
