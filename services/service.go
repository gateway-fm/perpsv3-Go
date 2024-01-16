package services

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gateway-fm/perpsv3-Go/config"
	"github.com/gateway-fm/perpsv3-Go/rawContracts"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// IService is a service layer interface
type IService interface {
	// RetrieveTrades is used to get logs from the "OrderSettled" event preps market contract within given block range
	RetrieveTrades(fromBlock uint64, toBLock *uint64) ([]*models.Trade, error)

	// RetrieveTradesLimit is used to get all trades and their additional data from the contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveTradesLimit(limit uint64) ([]*models.Trade, error)

	// RetrieveOrders is used to get logs from the "OrderCommitted" event preps market contract within given block range
	RetrieveOrders(fromBlock uint64, toBLock *uint64) ([]*models.Order, error)

	// RetrieveOrdersLimit is used to get all orders and their additional data from the contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveOrdersLimit(limit uint64) ([]*models.Order, error)

	// RetrieveMarketUpdates is used to get logs from the "MarketUpdated" event preps market contract within given block
	// range
	RetrieveMarketUpdates(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdate, error)

	// RetrieveMarketUpdatesBig is used to get logs from the "MarketUpdated" event preps market contract within given block
	// range and return model with big.Int values
	RetrieveMarketUpdatesBig(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdateBig, error)

	// RetrieveMarketUpdatesLimit is used to get all market updates and their additional data from the contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveMarketUpdatesLimit(limit uint64) ([]*models.MarketUpdate, error)

	// RetrieveMarketUpdatesBigLimit is used to get logs from the "MarketUpdated" event preps market contract within given block
	// range and return the model with big.Int values
	RetrieveMarketUpdatesBigLimit(limit uint64) ([]*models.MarketUpdateBig, error)

	// RetrieveLiquidations is used to get logs from the "PositionLiquidated" event preps market contract within given block
	// range
	RetrieveLiquidations(fromBlock uint64, toBLock *uint64) ([]*models.Liquidation, error)

	// RetrieveLiquidationsLimit is used to get all liquidations and their additional data from the contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveLiquidationsLimit(limit uint64) ([]*models.Liquidation, error)

	// RetrieveAccountLiquidationsLimit is used to get all account liquidated events from the contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveAccountLiquidationsLimit(limit uint64) ([]*models.AccountLiquidated, error)

	// RetrieveUSDMintedLimit is used to get all `usdMinted` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveUSDMintedLimit(limit uint64) ([]*models.USDMinted, error)

	// GetPosition is used to get "Position" data struct from the latest block from the perps market with given data
	GetPosition(accountID *big.Int, marketID *big.Int) (*models.Position, error)

	// GetMarketMetadata is used to get market metadata by given market ID
	GetMarketMetadata(marketID *big.Int) (*models.MarketMetadata, error)

	// GetMarketSummary is used to get market summary by given market ID
	GetMarketSummary(marketID *big.Int) (*models.MarketSummary, error)

	// GetMarketIDs is used to get market IDs from the smart contract
	GetMarketIDs() ([]*big.Int, error)

	// GetFoundingRate is used to get current founding rate by given market ID
	GetFoundingRate(marketId *big.Int) (*big.Int, error)

	// GetAvailableMargin is used to get available margin for given account ID
	GetAvailableMargin(accountId *big.Int) (*big.Int, error)

	// GetLiquidationParameters is used to get liquidation params for given market ID
	GetLiquidationParameters(marketId *big.Int) (*models.LiquidationParameters, error)

	// GetFundingParameters is used to get funding params for given market ID
	GetFundingParameters(marketId *big.Int) (*models.FundingParameters, error)

	// GetAccountLastInteraction is used to get accounts last interaction for given account ID
	GetAccountLastInteraction(accountId *big.Int) (*big.Int, error)

	// GetAccountOwner is used to get accounts owner address for given account ID
	GetAccountOwner(accountId *big.Int) (string, error)

	// GetCollateralAmount is used to get accounts collateral amount for given market ID
	GetCollateralAmount(accountId *big.Int, marketId *big.Int) (*big.Int, error)

	// GetRequiredMaintenanceMargin is used to get required maintenance margin for given account ID
	GetRequiredMaintenanceMargin(accountId *big.Int) (*big.Int, error)

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
	chainID          config.ChainID
	rpcClient        *ethclient.Client
	multicallRetries int
	multicallWait    time.Duration

	core           *core.Core
	coreFirstBlock uint64

	perpsMarket           *perpsMarket.PerpsMarket
	rawPerpsContract      rawContracts.IRawPerpsContract
	perpsMarketFirstBlock uint64

	rawERC7412   rawContracts.IRawERC7412Contract
	rawForwarder rawContracts.IRawForwarderContract
}

// NewService is used to get instance of Service
func NewService(
	rpc *ethclient.Client,
	conf *config.PerpsvConfig,
	core *core.Core,
	perps *perpsMarket.PerpsMarket,
) (IService, error) {
	s := &Service{
		chainID:          conf.ChainID,
		rpcClient:        rpc,
		multicallRetries: conf.Multicall.Retries,
		multicallWait:    conf.Multicall.Wait,

		core:           core,
		coreFirstBlock: conf.FirstContractBlocks.Core,

		perpsMarket:           perps,
		perpsMarketFirstBlock: conf.FirstContractBlocks.PerpsMarket,
	}

	rawPerpsContract, err := rawContracts.NewPerps(common.HexToAddress(conf.ContractAddresses.PerpsMarket), rpc)
	if err != nil {
		return nil, err
	}

	s.rawPerpsContract = rawPerpsContract

	if conf.ChainID == config.BaseMainnet || conf.ChainID == config.BaseAndromeda {
		rawERC7412, err := rawContracts.NewERC7412(common.HexToAddress(conf.ContractAddresses.ERC7412), rpc)
		if err != nil {
			return nil, err
		}

		s.rawERC7412 = rawERC7412

		rawForwarder, err := rawContracts.NewForwarder(common.HexToAddress(conf.ContractAddresses.Forwarder), rpc)
		if err != nil {
			return nil, err
		}

		s.rawForwarder = rawForwarder
	}

	return s, nil
}

// getIterationsForLimitQuery is used to get iterations of querying data from the contract with given rpc limit for blocks
// and latest block number. Limit by default (if given limit is 0) is set to 20 000 blocks
func (s *Service) getIterationsForLimitQuery(limit uint64) (iterations uint64, lastBlock uint64, err error) {
	lastBlock, err = s.rpcClient.BlockNumber(context.Background())
	if err != nil {
		logger.Log().WithField("layer", "Service-getIterationsForLimitQuery").Errorf("get latest block rpc error: %v", err.Error())
		return 0, 0, errors.GetRPCProviderErr(err, "BlockNumber")
	}

	if limit == 0 {
		limit = 20000
	}

	iterations = (lastBlock-s.perpsMarketFirstBlock)/limit + 1

	return iterations, lastBlock, nil
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

// getFilterOptsPerpsMarket is used to get options for event filtering on perps market contract
func (s *Service) getFilterOptsCore(fromBlock uint64, toBLock *uint64) *bind.FilterOpts {
	if fromBlock == 0 {
		fromBlock = s.coreFirstBlock
	}

	return &bind.FilterOpts{
		Start:   fromBlock,
		End:     toBLock,
		Context: context.Background(),
	}
}
