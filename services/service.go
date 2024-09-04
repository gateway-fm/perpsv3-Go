package services

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gateway-fm/perpsv3-Go/config"
	"github.com/gateway-fm/perpsv3-Go/contracts/Account"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
	"github.com/gateway-fm/perpsv3-Go/rawContracts"
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

	// RetrieveUSDMintedLimit is used to get all `usdMinted` events from the Core contract with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveUSDMinted(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.USDMinted, error)

	// RetrieveUSDBurnedLimit is used to get all `usdBurned` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveUSDBurnedLimit(limit uint64) ([]*models.USDBurned, error)

	// RetrieveUSDBurnedLimit is used to get all `usdBurned` events from the Core contract with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveUSDBurned(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.USDBurned, error)

	// RetrieveDelegationUpdatedLimit is used to get all `DelegationUpdated` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveDelegationUpdatedLimit(limit uint64) ([]*models.DelegationUpdated, error)

	// RetrieveDelegationUpdated is used to get all `DelegationUpdated` events with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveDelegationUpdated(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.DelegationUpdated, error)

	// RetrieveCollateralWithdrawnLimit is used to get all `Withdrawn` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveCollateralWithdrawnLimit(limit uint64) ([]*models.CollateralWithdrawn, error)

	// RetrieveCollateralWithdrawn is used to get all `Withdrawn` events from the Core contract with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveCollateralWithdrawn(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CollateralWithdrawn, error)

	// RetrieveCollateralDepositedLimit is used to get all `Deposited` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveCollateralDepositedLimit(limit uint64) ([]*models.CollateralDeposited, error)

	// RetrieveCollateralDepositedLimit is used to get all `Deposited` events from the Core contract with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveCollateralDeposited(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CollateralDeposited, error)

	// RetrieveRewardClaimedLimit is used to get all `RewardClaimed` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveRewardClaimedLimit(limit uint64) ([]*models.RewardClaimed, error)

	// RetrieveRewardClaimed is used to get all `DelegationUpdated` events with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveRewardClaimed(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.RewardClaimed, error)

	// RetrieveRewardDistributedLimit is used to get all `RewardDistributed` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveRewardDistributedLimit(limit uint64) ([]*models.RewardDistributed, error)

	// RetrieveRewardDistributed is used to get all `DelegationUpdated` events with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveRewardDistributed(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.RewardDistributed, error)

	// RetrieveMarketUSDDepositedLimit is used to get all `MarketUSDDeposited` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveMarketUSDDepositedLimit(limit uint64) ([]*models.MarketUSDDeposited, error)

	// RetrieveMarketUSDWithdrawnLimit is used to get all `MarketUSDWithdrawn` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveMarketUSDWithdrawnLimit(limit uint64) ([]*models.MarketUSDWithdrawn, error)

	// RetrieveMarketRegistered is used to get all `MarketRegistered` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveMarketRegistered(limit uint64) ([]*models.MarketRegistered, error)

	// RetrieveMarketRegisteredOpts is used to get all `MarketRegistered` events with given start block, end block and block search
	//	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveMarketRegisteredOpts(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.MarketRegistered, error)

	// RetrievePoolCreated is used to get all `PoolCreated` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrievePoolCreatedLimit(limit uint64) ([]*models.PoolCreated, error)

	RetrievePoolCreated(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.PoolCreated, error)

	// RetrieveLiquidationsCoreLimit is used to get all `Liquidation` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	//RetrieveLiquidationsCore(limit uint64) ([]*models.CoreLiquidation, error)
	RetrieveLiquidationsCoreLimit(limit uint64) ([]*models.CoreLiquidation, error)

	//RetrieveVaultLiquidationsCoreLimit is used to get all `VaultLiquidation` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	//RetrieveVaultLiquidationsCore(limit uint64) ([]*models.CoreVaultLiquidation, error)
	RetrieveVaultLiquidationsCoreLimit(limit uint64) ([]*models.CoreVaultLiquidation, error)

	// RetrieveLiquidationsCore is used to get all `DelegationUpdated` events with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveLiquidationsCore(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CoreLiquidation, error)

	// RetrieveVaultLiquidationsCore is used to get all `DelegationUpdated` events with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveVaultLiquidationsCore(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CoreVaultLiquidation, error)

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

	// GetCollateralPrice is used to get collateral price for given block number and collateralType
	GetCollateralPrice(blockNumber *big.Int, collateralType common.Address) (*models.CollateralPrice, error)

	// GetVaultDebt is used to get vault debt for given pool ID and collateralType
	GetVaultDebt(poolID *big.Int, collateralType common.Address) (*big.Int, error)

	// GetVaultDebtHistorical is used to get vault debt for given pool ID, collateralType and block number
	GetVaultDebtHistorical(poolID *big.Int, collateralType common.Address, blockNumber *big.Int) (*big.Int, error)

	// GetVaultCollateral is used to get vault collateral for given pool ID and collateralType
	GetVaultCollateral(poolID *big.Int, collateralType common.Address) (amount *big.Int, value *big.Int, err error)

	// GetVaultCollateralHistorical is used to get vault collateral for given pool ID, collateralType and block number
	GetVaultCollateralHistorical(poolID *big.Int, collateralType common.Address, blockNumber *big.Int) (amount *big.Int, value *big.Int, err error)

	// GetPoolConfiguration is used to get MarketConfigurations array
	GetPoolConfiguration(poolID *big.Int) (*models.PoolConfiguration, error)

	// GetPoolName is used to get pool name from given PoolID
	GetPoolName(poolID *big.Int) (string, error)

	// GetAccountCollateralCore is used to get account collateral data for given account ID and collateral type
	GetAccountCollateralCore(accountId *big.Int, collateralType common.Address) (*models.AccountCollateral, error)

	// GetAccountAvailableCollateral is used to get account available collateral data for given account ID and collateral type
	GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error)

	// GetCollateralConfigurations is used to get CollateralConfiguration data
	GetCollateralConfigurations(hideDisabled bool) ([]*models.CollateralConfiguration, error)

	// FormatAccount is used to get account, and it's additional data from the contract by given account id
	FormatAccount(id *big.Int) (*models.Account, error)

	// FormatAccounts is used to get all accounts and their additional data from the contract
	FormatAccounts() ([]*models.Account, error)

	// FormatAccountsLimit is used to get all accounts and their additional data from the contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	FormatAccountsLimit(limit uint64) ([]*models.Account, error)

	// FormatAccountCore is used to get all accounts and their additional data from the core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	FormatAccountCore(id *big.Int) (*models.Account, error)

	// FormatAccountsCoreLimit is used to get all accounts and their additional data from the contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	FormatAccountsCoreLimit(limit uint64) ([]*models.Account, error)

	// FormatAccountsCoreLimit is used to get all accounts and their additional data from the contract with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	FormatAccountsCore(fromBlock, toBlock, limit uint64) ([]*models.Account, error)

	// RetrieveChangeOwner is used to get all owner changes and additional data from the contract with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveChangeOwner(fromBlock, toBlock, limit uint64) ([]*models.AccountTransfer, error)

	// RetrievePermissionRevoked is used to get all the revoked permission and additional data from the contract with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrievePermissionRevoked(fromBlock, toBlock, limit uint64) ([]*models.PermissionChanged, error)

	// RetrievePermissionGranted is used to get all the granted permission and additional data from the contract with given start block, end block and block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrievePermissionGranted(fromBlock, toBlock, limit uint64) ([]*models.PermissionChanged, error)
}

type ContractType int

const (
	ContractPerpsMarket ContractType = iota

	ContractCore
)

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
	rawCore      rawContracts.IRawCoreContract

	accountContract *Account.Account
}

// NewService is used to get instance of Service
func NewService(
	rpc *ethclient.Client,
	conf *config.PerpsvConfig,
	core *core.Core,
	perps *perpsMarket.PerpsMarket,
	accountContract *Account.Account,
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
		accountContract:       accountContract,
	}

	rawPerpsContract, err := rawContracts.NewPerps(common.HexToAddress(conf.ContractAddresses.PerpsMarket), rpc)
	if err != nil {
		return nil, err
	}

	s.rawPerpsContract = rawPerpsContract

	rawCoreContract, err := rawContracts.NewCore(common.HexToAddress(conf.ContractAddresses.Core), rpc)
	if err != nil {
		return nil, err
	}

	s.rawCore = rawCoreContract

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

// getIterationsForLimitQueryPerpsMarket is used to get iterations of querying data from the contract with given rpc limit for blocks
// and latest block number. Limit by default (if given limit is 0) is set to 20 000 blocks
func (s *Service) getIterationsForLimitQueryPerpsMarket(limit uint64) (iterations uint64, lastBlock uint64, err error) {
	return s.getIterationsForQuery(0, 0, limit, ContractPerpsMarket)
}

// getIterationsForLimitQueryCore is used to get iterations of querying data from the contract with given rpc limit for blocks
// and latest block number. Limit by default (if given limit is 0) is set to 20 000 blocks
func (s *Service) getIterationsForLimitQueryCore(limit uint64) (iterations uint64, lastBlock uint64, err error) {
	return s.getIterationsForQuery(0, 0, limit, ContractCore)
}

func (s *Service) getIterationsForQuery(fromBlock uint64, toBlock uint64, limit uint64, contractType ContractType) (iterations uint64, lastBlock uint64, err error) {
	if fromBlock == 0 {
		switch contractType {
		case ContractPerpsMarket:
			fromBlock = s.perpsMarketFirstBlock
		case ContractCore:
			fromBlock = s.coreFirstBlock
		}
	}

	if toBlock == 0 {
		lastBlock, err = s.rpcClient.BlockNumber(context.Background())
		if err != nil {
			logger.Log().WithField("layer", "Service-getIterationsForQuery").Errorf("get latest block rpc error: %v", err.Error())
			return 0, 0, errors.GetRPCProviderErr(err, "BlockNumber")
		}

		toBlock = lastBlock
	}

	if limit == 0 {
		limit = 20000
	}

	if fromBlock > toBlock {
		logger.Log().WithField("layer", "Service-getIterationsForQuery").Error("fromBlock > toBlock")
		return 0, 0, fmt.Errorf("fromBlock > toBlock")
	}

	iterations = (toBlock-fromBlock)/limit + 1

	return iterations, toBlock, nil
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
