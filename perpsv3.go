package perpsv3_Go

import (
	"math/big"

	"github.com/gateway-fm/perpsv3-Go/contracts/Account"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gateway-fm/perpsv3-Go/config"
	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/errors"
	"github.com/gateway-fm/perpsv3-Go/events"
	"github.com/gateway-fm/perpsv3-Go/models"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
	"github.com/gateway-fm/perpsv3-Go/services"
)

// IPerpsv3 is an interface for perpsv3 lib
type IPerpsv3 interface {
	// RetrieveTrades is used to get logs from the "OrderSettled" event perps market contract within given block range
	//   - use 0 for fromBlock to use default value of a first contract block
	//   - use nil for toBlock to use default value of a last blockchain block
	RetrieveTrades(fromBlock uint64, toBLock *uint64) ([]*models.Trade, error)

	// RetrieveTradesLimit is used to get all "OrderSettled" events and their additional data from the contract
	// with given block search limit. If given limit is 0 function will set default value to 20 000 blocks
	RetrieveTradesLimit(limit uint64) ([]*models.Trade, error)

	// RetrieveOrders is used to get logs from the "OrderCommitted" event perps market contract within given block range
	//   - use 0 for fromBlock to use default value of a first contract block
	//   - use nil for toBlock to use default value of a last blockchain block
	RetrieveOrders(fromBlock uint64, toBLock *uint64) ([]*models.Order, error)

	// RetrieveOrdersLimit is used to get all "OrderCommitted" events and their additional data from the contract
	// with given block search limit. If given limit is 0 function will set default value to 20 000 blocks
	RetrieveOrdersLimit(limit uint64) ([]*models.Order, error)

	// RetrieveMarketUpdates is used to get logs from the "MarketUpdated" event perps market contract within given block
	// range
	//   - use 0 for fromBlock to use default value of a first contract block
	//   - use nil for toBlock to use default value of a last blockchain block
	RetrieveMarketUpdates(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdate, error)

	// RetrieveMarketUpdatesBig is used to get logs from the "MarketUpdated" event perps market contract within given block
	// range
	//   - use 0 for fromBlock to use default value of a first contract block
	//   - use nil for toBlock to use default value of a last blockchain block
	// It will return a MarketUpdateBig model with big.Int values
	RetrieveMarketUpdatesBig(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdateBig, error)

	// RetrieveMarketUpdatesLimit is used to get all "MarketUpdated" events and their additional data from the contract
	// with given block search limit. If given limit is 0 function will set default value to 20 000 blocks
	RetrieveMarketUpdatesLimit(limit uint64) ([]*models.MarketUpdate, error)

	// RetrieveMarketUpdatesBigLimit is used to get all "MarketUpdated" events and their additional data from the contract
	// with given block search limit. If given limit is 0 function will set default value to 20 000 blocks
	// It will return a MarketUpdateBig model with big.Int values
	RetrieveMarketUpdatesBigLimit(limit uint64) ([]*models.MarketUpdateBig, error)

	// RetrieveLiquidations is used to get logs from the "PositionLiquidated" event perps market contract within given block
	// range
	//   - use 0 for fromBlock to use default value of a first contract block
	//   - use nil for toBlock to use default value of a last blockchain block
	RetrieveLiquidations(fromBlock uint64, toBLock *uint64) ([]*models.Liquidation, error)

	// RetrieveLiquidationsLimit is used to get all "PositionLiquidated" events and their additional data from the contract
	// with given block search limit. For most public RPC providers the value for limit is 20 000 blocks
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

	// RetrieveCollateralWithdrawnLimit is used to get all `Withdrawn` events from the Core contract with given start block, end block and block search
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
	RetrieveRewardClaimed(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.RewardClaimed, error)

	// RetrieveRewardDistributedLimit is used to get all `RewardDistributed` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveRewardDistributedLimit(limit uint64) ([]*models.RewardDistributed, error)
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

	// RetrievePoolCreatedLimit is used to get all `PoolCreated` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrievePoolCreatedLimit(limit uint64) ([]*models.PoolCreated, error)
	RetrievePoolCreated(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.PoolCreated, error)

	// RetrieveLiquidationsCoreLimit is used to get all `Liquidation` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveLiquidationsCoreLimit(limit uint64) ([]*models.CoreLiquidation, error)
	RetrieveLiquidationsCore(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CoreLiquidation, error)

	// RetrieveVaultLiquidationsCoreLimit is used to get all `VaultLiquidation` events from the Core contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	RetrieveVaultLiquidationsCoreLimit(limit uint64) ([]*models.CoreVaultLiquidation, error)
	RetrieveVaultLiquidationsCore(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CoreVaultLiquidation, error)

	// ListenTrades is used to subscribe on the contract "OrderSettled" event. The goroutine will return events on the
	// TradesChan chanel and errors on the ErrChan chanel.
	// To close the subscription use events.TradeSubscription `Close` function
	ListenTrades() (*events.TradeSubscription, error)

	// ListenOrders is used to subscribe on the contract "OrderCommitted" event. The goroutine will return events on the
	// OrdersChan chanel and errors on the ErrChan chanel.
	// To close the subscription use events.OrderSubscription `Close` function
	ListenOrders() (*events.OrderSubscription, error)

	// ListenMarketUpdates is used to subscribe on the contract "MarketUpdated" event. The goroutine will return events
	// on the MarketUpdateChan chanel and errors on the ErrChan chanel.
	// To close the subscription use events.MarketUpdateSubscription `Close` function
	ListenMarketUpdates() (*events.MarketUpdateSubscription, error)

	// ListenMarketUpdatesBig is used to subscribe on the contract "MarketUpdated" event. The goroutine will return events
	// on the MarketUpdateChan chanel and errors on the ErrChan chanel.
	// To close the subscription use events.MarketUpdateSubscription `Close` function
	ListenMarketUpdatesBig() (*events.MarketUpdateSubscriptionBig, error)

	// ListenLiquidations is used to subscribe on the contract "PositionLiquidated" event. The goroutine will return events
	// on the LiquidationsChan chanel and errors on the ErrChan chanel.
	// To close the subscription use events.LiquidationSubscription `Close` function
	ListenLiquidations() (*events.LiquidationSubscription, error)

	// ListenAccountCreated is used to listen to all 'AccountCreated' contract events and return them as models.Account
	// struct and return errors on ErrChan chanel
	ListenAccountCreated() (*events.AccountCreatedSubscription, error)

	// ListenAccountCreatedCore is used to listen to all 'AccountCreated' core contract events and return them as models.Account
	// struct and return errors on ErrChan chanel
	ListenAccountCreatedCore() (*events.AccountCreatedCoreSubscription, error)

	// ListenAccountLiquidated is used to listen to all 'AccountLiquidated' contract events and return them as models.AccountLiquidated
	// struct and return errors on ErrChan chanel
	ListenAccountLiquidated() (*events.AccountLiquidatedSubscription, error)

	// ListenAccountPermissionRevoked is used to listen to all 'PermissionRevoked' contract events and return them as models.PermissionChanged
	// struct and return errors on ErrChan chanel
	ListenAccountPermissionRevoked() (*events.AccountPermissionRevokedSubscription, error)

	// ListenAccountPermissionGranted is used to listen to all 'PermissionGranted' contract events and return them as models.PermissionChanged
	// struct and return errors on ErrChan chanel
	ListenAccountPermissionGranted() (*events.AccountPermissionGrantedSubscription, error)

	// ListenUSDMinted is used to listen to all 'USDMinted' Core contract events and return them as models.USDMinted
	// struct and return errors on ErrChan chanel
	ListenUSDMinted() (*events.USDMintedSubscription, error)

	// ListenUSDBurned is used to listen to all 'USDBurned' Core contract events and return them as models.USDBurned
	// struct and return errors on ErrChan chanel
	ListenUSDBurned() (*events.USDBurnedSubscription, error)

	// ListenDelegationUpdated is used to listen to all 'DelegationUpdated' Core contract events and return them as models.DelegationUpdated
	// struct and return errors on ErrChan chanel
	ListenDelegationUpdated() (*events.DelegationUpdatedSubscription, error)

	// ListenCollateralWithdrawn is used to listen to all 'Withdrawn' Core contract events and return them as models.CollateralWithdrawn
	// struct and return errors on ErrChan chanel
	ListenCollateralWithdrawn() (*events.CollateralWithdrawnSubscription, error)

	// ListenCollateralDeposited is used to listen to all 'Deposited' Core contract events and return them as models.CollateralDeposited
	// struct and return errors on ErrChan chanel
	ListenCollateralDeposited() (*events.CollateralDepositedSubscription, error)

	// ListenRewardDistributed is used to listen to all 'RewardDistributed' Core contract events and return them as models.RewardDistributed
	// struct and return errors on ErrChan chanel
	ListenRewardDistributed() (*events.RewardDistributedSubscription, error)

	// ListenRewardClaimed is used to listen to all 'RewardClaimed' Core contract events and return them as models.RewardClaimed
	// struct and return errors on ErrChan chanel
	ListenRewardClaimed() (*events.RewardClaimedSubscription, error)

	// ListenMarketUSDWithdrawn is used to listen to all 'MarketUSDWithdrawn' Core contract events and return them as models.MarketUSDWithdrawn
	// struct and return errors on ErrChan chanel
	ListenMarketUSDWithdrawn() (*events.MarketUSDWithdrawnSubscription, error)

	// ListenMarketUSDDeposited is used to listen to all 'MarketUSDDeposited' Core contract events and return them as models.MarketUSDDeposited
	// struct and return errors on ErrChan chanel
	ListenMarketUSDDeposited() (*events.MarketUSDDepositedSubscription, error)

	// ListenMarketRegistered is used to listen to all 'MarketRegistered' Core contract events and return them as models.MarketRegistered
	// struct and return errors on ErrChan chanel
	ListenMarketRegistered() (*events.MarketRegisteredSubscription, error)

	// ListenPoolCreated is used to listen to all 'PoolCreated' Core contract events and return them as models.PoolCreated
	// struct and return errors on ErrChan chanel
	ListenPoolCreated() (*events.PoolCreatedSubscription, error)

	// ListenVaultLiquidationsCore is used to listen to all 'VaultLiquidations' Core contract events and return them as models.CoreVaultLiquidation
	// struct and return errors on ErrChan chanel
	ListenVaultLiquidationsCore() (*events.VaultLiquidationsCoreSubscription, error)

	// ListenLiquidationsCore is used to listen to all 'Liquidations' Core contract events and return them as models.CoreLiquidation
	// struct and return errors on ErrChan chanel
	ListenLiquidationsCore() (*events.LiquidationsCoreSubscription, error)

	// ListenAccountTransfer is used to listen to all 'Transfer' Account contract events and return them as models.AccountTransfer
	// struct and return errors on ErrChan chanel
	ListenAccountTransfer() (*events.AccountTransferSubscription, error)

	// ListenAccountCorePermissionRevoked is used to listen to all 'PermissionRevoked' Core contract events and return them as models.PermissionChanged
	// struct and return errors on ErrChan chanel
	ListenAccountCorePermissionRevoked() (*events.AccountCorePermissionRevokedSubscription, error)

	// ListenAccountCorePermissionGranted is used to listen to all 'PermissionGranted' Core contract events and return them as models.PermissionChanged
	// struct and return errors on ErrChan chanel
	ListenAccountCorePermissionGranted() (*events.AccountCorePermissionGrantedSubscription, error)

	// GetPosition is used to get position data struct from latest block with given params
	// Function can return contract error if market ID is invalid
	GetPosition(accountID *big.Int, marketID *big.Int) (*models.Position, error)

	// GetMarketMetadata is used to get market metadata by given market ID. Given market id cannot be nil and should exist
	// in the smart contract
	GetMarketMetadata(marketID *big.Int) (*models.MarketMetadata, error)

	// GetMarketSummary is used to get market summary by given market ID. Given market id cannot be nil
	GetMarketSummary(marketID *big.Int) (*models.MarketSummary, error)

	// GetMarketIDs is used to get market IDs from the smart contract
	GetMarketIDs() ([]*big.Int, error)

	// GetFoundingRate is used to get current market founding rate by given market ID
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

	// Config is used to get current lib config
	Config() *config.PerpsvConfig

	// Close used to stop the lib work
	Close()
}

// Perpsv3 is a main perpsv3 lib object, it is implementing IPerpsv3 interface
type Perpsv3 struct {
	config    *config.PerpsvConfig
	service   services.IService
	events    events.IEvents
	rpcClient *ethclient.Client
}

// Create used to get Perpsv3 instance with given configuration settings
func Create(conf *config.PerpsvConfig) (IPerpsv3, error) {
	lib := &Perpsv3{
		config: conf,
	}

	if err := lib.init(); err != nil {
		return nil, err
	}

	return lib, nil
}

func GetOptimismGoerliDefaultConfig(rpcURL string) *config.PerpsvConfig {
	return config.GetOptimismGoerliDefaultConfig(rpcURL)
}

func GetBaseAndromedaDefaultConfig(rpcURL string) *config.PerpsvConfig {
	return config.GetBaseAndromedaDefaultConfig(rpcURL)
}

func GetBaseSepoliaDefaultConfig(rpcURL string) *config.PerpsvConfig {
	return config.GetBaseSepoliaDefaultConfig(rpcURL)
}

func GetBaseMainnetDefaultConfig(rpcURL string) *config.PerpsvConfig {
	return config.GetBaseMainnetDefaultConfig(rpcURL)
}

func (p *Perpsv3) RetrieveTrades(fromBlock uint64, toBLock *uint64) ([]*models.Trade, error) {
	return p.service.RetrieveTrades(fromBlock, toBLock)
}

func (p *Perpsv3) RetrieveTradesLimit(limit uint64) ([]*models.Trade, error) {
	return p.service.RetrieveTradesLimit(limit)
}

func (p *Perpsv3) RetrieveOrders(fromBlock uint64, toBLock *uint64) ([]*models.Order, error) {
	return p.service.RetrieveOrders(fromBlock, toBLock)
}

func (p *Perpsv3) RetrieveOrdersLimit(limit uint64) ([]*models.Order, error) {
	return p.service.RetrieveOrdersLimit(limit)
}

func (p *Perpsv3) RetrieveMarketUpdates(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdate, error) {
	return p.service.RetrieveMarketUpdates(fromBlock, toBLock)
}

func (p *Perpsv3) RetrieveMarketUpdatesBig(fromBlock uint64, toBLock *uint64) ([]*models.MarketUpdateBig, error) {
	return p.service.RetrieveMarketUpdatesBig(fromBlock, toBLock)
}

func (p *Perpsv3) RetrieveMarketUpdatesLimit(limit uint64) ([]*models.MarketUpdate, error) {
	return p.service.RetrieveMarketUpdatesLimit(limit)
}

func (p *Perpsv3) RetrieveMarketUpdatesBigLimit(limit uint64) ([]*models.MarketUpdateBig, error) {
	return p.service.RetrieveMarketUpdatesBigLimit(limit)
}

func (p *Perpsv3) RetrieveLiquidations(fromBlock uint64, toBLock *uint64) ([]*models.Liquidation, error) {
	return p.service.RetrieveLiquidations(fromBlock, toBLock)
}

func (p *Perpsv3) RetrieveLiquidationsLimit(limit uint64) ([]*models.Liquidation, error) {
	return p.service.RetrieveLiquidationsLimit(limit)
}

func (p *Perpsv3) RetrieveAccountLiquidationsLimit(limit uint64) ([]*models.AccountLiquidated, error) {
	return p.service.RetrieveAccountLiquidationsLimit(limit)
}

func (p *Perpsv3) RetrieveUSDMintedLimit(limit uint64) ([]*models.USDMinted, error) {
	return p.service.RetrieveUSDMintedLimit(limit)
}

func (p *Perpsv3) RetrieveUSDMinted(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.USDMinted, error) {
	return p.service.RetrieveUSDMinted(fromBlock, toBlock, limit)
}

func (p *Perpsv3) RetrieveUSDBurnedLimit(limit uint64) ([]*models.USDBurned, error) {
	return p.service.RetrieveUSDBurnedLimit(limit)
}

func (p *Perpsv3) RetrieveUSDBurned(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.USDBurned, error) {
	return p.service.RetrieveUSDBurned(fromBlock, toBlock, limit)
}

func (p *Perpsv3) RetrieveDelegationUpdatedLimit(limit uint64) ([]*models.DelegationUpdated, error) {
	return p.service.RetrieveDelegationUpdatedLimit(limit)
}

func (p *Perpsv3) RetrieveDelegationUpdated(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.DelegationUpdated, error) {
	return p.service.RetrieveDelegationUpdated(fromBlock, toBlock, limit)
}

func (p *Perpsv3) RetrieveCollateralWithdrawnLimit(limit uint64) ([]*models.CollateralWithdrawn, error) {
	return p.service.RetrieveCollateralWithdrawnLimit(limit)
}

func (p *Perpsv3) RetrieveCollateralWithdrawn(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CollateralWithdrawn, error) {
	return p.service.RetrieveCollateralWithdrawn(fromBlock, toBlock, limit)
}

func (p *Perpsv3) RetrieveCollateralDepositedLimit(limit uint64) ([]*models.CollateralDeposited, error) {
	return p.service.RetrieveCollateralDepositedLimit(limit)
}

func (p *Perpsv3) RetrieveCollateralDeposited(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CollateralDeposited, error) {
	return p.service.RetrieveCollateralDeposited(fromBlock, toBlock, limit)
}

func (p *Perpsv3) RetrieveRewardClaimedLimit(limit uint64) ([]*models.RewardClaimed, error) {
	return p.service.RetrieveRewardClaimedLimit(limit)
}

func (p *Perpsv3) RetrieveRewardClaimed(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.RewardClaimed, error) {
	return p.service.RetrieveRewardClaimed(fromBlock, toBlock, limit)
}

func (p *Perpsv3) RetrieveRewardDistributedLimit(limit uint64) ([]*models.RewardDistributed, error) {
	return p.service.RetrieveRewardDistributedLimit(limit)
}

func (p *Perpsv3) RetrieveRewardDistributed(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.RewardDistributed, error) {
	return p.service.RetrieveRewardDistributed(fromBlock, toBlock, limit)
}

func (p *Perpsv3) RetrieveMarketUSDDepositedLimit(limit uint64) ([]*models.MarketUSDDeposited, error) {
	return p.service.RetrieveMarketUSDDepositedLimit(limit)
}

func (p *Perpsv3) RetrieveMarketUSDWithdrawnLimit(limit uint64) ([]*models.MarketUSDWithdrawn, error) {
	return p.service.RetrieveMarketUSDWithdrawnLimit(limit)
}

func (p *Perpsv3) RetrieveMarketRegistered(limit uint64) ([]*models.MarketRegistered, error) {
	return p.service.RetrieveMarketRegistered(limit)
}

func (p *Perpsv3) RetrieveMarketRegisteredOpts(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.MarketRegistered, error) {
	return p.service.RetrieveMarketRegisteredOpts(fromBlock, toBlock, limit)
}

func (p *Perpsv3) RetrievePoolCreated(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.PoolCreated, error) {
	return p.service.RetrievePoolCreated(fromBlock, toBlock, limit)
}

func (p *Perpsv3) RetrievePoolCreatedLimit(limit uint64) ([]*models.PoolCreated, error) {
	return p.service.RetrievePoolCreatedLimit(limit)
}

func (p *Perpsv3) RetrieveLiquidationsCoreLimit(limit uint64) ([]*models.CoreLiquidation, error) {
	return p.service.RetrieveLiquidationsCoreLimit(limit)
}

func (p *Perpsv3) RetrieveVaultLiquidationsCoreLimit(limit uint64) ([]*models.CoreVaultLiquidation, error) {
	return p.service.RetrieveVaultLiquidationsCoreLimit(limit)
}

func (p *Perpsv3) RetrieveLiquidationsCore(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CoreLiquidation, error) {
	return p.service.RetrieveLiquidationsCore(fromBlock, toBlock, limit)
}

func (p *Perpsv3) RetrieveVaultLiquidationsCore(fromBlock uint64, toBlock uint64, limit uint64) ([]*models.CoreVaultLiquidation, error) {
	return p.service.RetrieveVaultLiquidationsCore(fromBlock, toBlock, limit)
}
func (p *Perpsv3) ListenTrades() (*events.TradeSubscription, error) {
	return p.events.ListenTrades()
}

func (p *Perpsv3) ListenOrders() (*events.OrderSubscription, error) {
	return p.events.ListenOrders()
}

func (p *Perpsv3) ListenMarketUpdates() (*events.MarketUpdateSubscription, error) {
	return p.events.ListenMarketUpdates()
}

func (p *Perpsv3) ListenMarketUpdatesBig() (*events.MarketUpdateSubscriptionBig, error) {
	return p.events.ListenMarketUpdatesBig()
}

func (p *Perpsv3) ListenLiquidations() (*events.LiquidationSubscription, error) {
	return p.events.ListenLiquidations()
}

func (p *Perpsv3) ListenAccountCreated() (*events.AccountCreatedSubscription, error) {
	return p.events.ListenAccountCreated()
}

func (p *Perpsv3) ListenAccountCreatedCore() (*events.AccountCreatedCoreSubscription, error) {
	return p.events.ListenAccountCreatedCore()
}

func (p *Perpsv3) ListenAccountLiquidated() (*events.AccountLiquidatedSubscription, error) {
	return p.events.ListenAccountLiquidated()
}

func (p *Perpsv3) ListenAccountPermissionRevoked() (*events.AccountPermissionRevokedSubscription, error) {
	return p.events.ListenAccountPermissionRevoked()
}

func (p *Perpsv3) ListenAccountPermissionGranted() (*events.AccountPermissionGrantedSubscription, error) {
	return p.events.ListenAccountPermissionGranted()
}

func (p *Perpsv3) ListenUSDMinted() (*events.USDMintedSubscription, error) {
	return p.events.ListenUSDMinted()
}

func (p *Perpsv3) ListenUSDBurned() (*events.USDBurnedSubscription, error) {
	return p.events.ListenUSDBurned()
}

func (p *Perpsv3) ListenDelegationUpdated() (*events.DelegationUpdatedSubscription, error) {
	return p.events.ListenDelegationUpdated()
}

func (p *Perpsv3) ListenCollateralWithdrawn() (*events.CollateralWithdrawnSubscription, error) {
	return p.events.ListenCollateralWithdrawn()
}

func (p *Perpsv3) ListenCollateralDeposited() (*events.CollateralDepositedSubscription, error) {
	return p.events.ListenCollateralDeposited()
}

func (p *Perpsv3) ListenRewardDistributed() (*events.RewardDistributedSubscription, error) {
	return p.events.ListenRewardDistributed()
}

func (p *Perpsv3) ListenRewardClaimed() (*events.RewardClaimedSubscription, error) {
	return p.events.ListenRewardClaimed()
}

func (p *Perpsv3) ListenMarketUSDWithdrawn() (*events.MarketUSDWithdrawnSubscription, error) {
	return p.events.ListenMarketUSDWithdrawn()
}

func (p *Perpsv3) ListenMarketUSDDeposited() (*events.MarketUSDDepositedSubscription, error) {
	return p.events.ListenMarketUSDDeposited()
}

func (p *Perpsv3) ListenMarketRegistered() (*events.MarketRegisteredSubscription, error) {
	return p.events.ListenMarketRegistered()
}

func (p *Perpsv3) ListenPoolCreated() (*events.PoolCreatedSubscription, error) {
	return p.events.ListenPoolCreated()
}

func (p *Perpsv3) ListenVaultLiquidationsCore() (*events.VaultLiquidationsCoreSubscription, error) {
	return p.events.ListenVaultLiquidationsCore()
}

func (p *Perpsv3) ListenLiquidationsCore() (*events.LiquidationsCoreSubscription, error) {
	return p.events.ListenLiquidationsCore()
}

func (p *Perpsv3) ListenAccountTransfer() (*events.AccountTransferSubscription, error) {
	return p.events.ListenAccountTransfer()
}

func (p *Perpsv3) ListenAccountCorePermissionRevoked() (*events.AccountCorePermissionRevokedSubscription, error) {
	return p.events.ListenAccountCorePermissionRevoked()
}

func (p *Perpsv3) ListenAccountCorePermissionGranted() (*events.AccountCorePermissionGrantedSubscription, error) {
	return p.events.ListenAccountCorePermissionGranted()
}

func (p *Perpsv3) GetPosition(accountID *big.Int, marketID *big.Int) (*models.Position, error) {
	return p.service.GetPosition(accountID, marketID)
}

func (p *Perpsv3) GetMarketMetadata(marketID *big.Int) (*models.MarketMetadata, error) {
	return p.service.GetMarketMetadata(marketID)
}

func (p *Perpsv3) GetMarketSummary(marketID *big.Int) (*models.MarketSummary, error) {
	return p.service.GetMarketSummary(marketID)
}

func (p *Perpsv3) GetMarketIDs() ([]*big.Int, error) {
	return p.service.GetMarketIDs()
}

func (p *Perpsv3) GetFoundingRate(marketId *big.Int) (*big.Int, error) {
	return p.service.GetFoundingRate(marketId)
}

func (p *Perpsv3) GetAvailableMargin(accountId *big.Int) (*big.Int, error) {
	return p.service.GetAvailableMargin(accountId)
}

func (p *Perpsv3) GetLiquidationParameters(marketId *big.Int) (*models.LiquidationParameters, error) {
	return p.service.GetLiquidationParameters(marketId)
}

func (p *Perpsv3) GetFundingParameters(marketId *big.Int) (*models.FundingParameters, error) {
	return p.service.GetFundingParameters(marketId)
}

func (p *Perpsv3) GetAccountLastInteraction(accountId *big.Int) (*big.Int, error) {
	return p.service.GetAccountLastInteraction(accountId)
}

func (p *Perpsv3) GetAccountOwner(accountId *big.Int) (string, error) {
	return p.service.GetAccountOwner(accountId)
}

func (p *Perpsv3) GetCollateralAmount(accountId *big.Int, marketId *big.Int) (*big.Int, error) {
	return p.service.GetCollateralAmount(accountId, marketId)
}

func (p *Perpsv3) GetRequiredMaintenanceMargin(accountId *big.Int) (*big.Int, error) {
	return p.service.GetRequiredMaintenanceMargin(accountId)
}

func (p *Perpsv3) GetCollateralPrice(blockNumber *big.Int, collateralType common.Address) (*models.CollateralPrice, error) {
	return p.service.GetCollateralPrice(blockNumber, collateralType)
}

func (p *Perpsv3) GetVaultDebt(poolID *big.Int, collateralType common.Address) (*big.Int, error) {
	return p.service.GetVaultDebt(poolID, collateralType)
}

func (p *Perpsv3) GetVaultDebtHistorical(poolID *big.Int, collateralType common.Address, blockNumber *big.Int) (*big.Int, error) {
	return p.service.GetVaultDebtHistorical(poolID, collateralType, blockNumber)
}

func (p *Perpsv3) GetVaultCollateral(poolID *big.Int, collateralType common.Address) (amount *big.Int, value *big.Int, err error) {
	return p.service.GetVaultCollateral(poolID, collateralType)
}

func (p *Perpsv3) GetVaultCollateralHistorical(poolID *big.Int, collateralType common.Address, blockNumber *big.Int) (amount *big.Int, value *big.Int, err error) {
	return p.service.GetVaultCollateralHistorical(poolID, collateralType, blockNumber)
}

func (p *Perpsv3) GetPoolConfiguration(poolID *big.Int) (*models.PoolConfiguration, error) {
	return p.service.GetPoolConfiguration(poolID)
}

func (p *Perpsv3) GetPoolName(poolID *big.Int) (string, error) {
	return p.service.GetPoolName(poolID)
}

func (p *Perpsv3) GetAccountCollateralCore(accountId *big.Int, collateralType common.Address) (*models.AccountCollateral, error) {
	return p.service.GetAccountCollateralCore(accountId, collateralType)
}

func (p *Perpsv3) GetAccountAvailableCollateral(accountId *big.Int, collateralType common.Address) (*big.Int, error) {
	return p.service.GetAccountAvailableCollateral(accountId, collateralType)
}

func (p *Perpsv3) GetCollateralConfigurations(hideDisabled bool) ([]*models.CollateralConfiguration, error) {
	return p.service.GetCollateralConfigurations(hideDisabled)
}

func (p *Perpsv3) FormatAccounts() ([]*models.Account, error) {
	return p.service.FormatAccounts()
}

func (p *Perpsv3) FormatAccount(id *big.Int) (*models.Account, error) {
	return p.service.FormatAccount(id)
}

func (p *Perpsv3) FormatAccountsLimit(limit uint64) ([]*models.Account, error) {
	return p.service.FormatAccountsLimit(limit)
}

func (p *Perpsv3) FormatAccountsCoreLimit(limit uint64) ([]*models.Account, error) {
	return p.service.FormatAccountsCoreLimit(limit)
}

func (p *Perpsv3) FormatAccountCore(id *big.Int) (*models.Account, error) {
	return p.service.FormatAccountCore(id)
}

func (p *Perpsv3) FormatAccountsCore(fromBlock, toBlock, limit uint64) ([]*models.Account, error) {
	return p.service.FormatAccountsCore(fromBlock, toBlock, limit)
}

func (p *Perpsv3) Config() *config.PerpsvConfig {
	return p.config
}

func (p *Perpsv3) Close() {
	p.rpcClient.Close()
}

// init used to initialize all lib dependencies
func (p *Perpsv3) init() error {
	if p.config.RPC == "" {
		return errors.BlankRPCURLErr
	}

	rpcClient, err := ethclient.Dial(p.config.RPC)
	if err != nil {
		logger.Log().WithField("layer", "Init").Errorf("error dial rpc: %v", err.Error())
		return errors.GetDialRPCErr(err)
	}

	p.rpcClient = rpcClient

	coreContact, err := p.getCoreContract()
	if err != nil {
		return err
	}

	perpsMarketContract, err := p.getPerpsMarket()
	if err != nil {
		return err
	}

	accountContract, err := p.getAccountContract()
	if err != nil {
		return err
	}

	srv, err := services.NewService(rpcClient, p.config, coreContact, perpsMarketContract)
	if err != nil {
		return err
	}

	p.service = srv
	p.events = events.NewEvents(rpcClient, coreContact, perpsMarketContract, accountContract)

	return nil
}

// getGoerliCoreContract is used to get core contract instance deployed on goerli test net
func (p *Perpsv3) getCoreContract() (*core.Core, error) {
	if p.config.ContractAddresses.Core != "" {
		addr, err := getAddr(p.config.ContractAddresses.Core, "core")
		if err != nil {
			return nil, err
		}

		contract, err := core.NewCore(addr, p.rpcClient)
		if err != nil {
			logger.Log().WithField("layer", "Init").Errorf("error getting core contract: %v", err.Error())
			return nil, errors.GetInitContractErr(err)
		}

		return contract, nil
	} else {
		logger.Log().WithField("layer", "Init").Errorf("no core contract address")
		return nil, errors.BlankContractAddrErr
	}
}

func (p *Perpsv3) getAccountContract() (*Account.Account, error) {
	if p.config.ContractAddresses.Account != "" {
		addr, err := getAddr(p.config.ContractAddresses.Account, "account")
		if err != nil {
			return nil, err
		}

		contract, err := Account.NewAccount(addr, p.rpcClient)
		if err != nil {
			logger.Log().WithField("layer", "Init").Errorf("error getting account contract: %v", err.Error())
			return nil, errors.GetInitContractErr(err)
		}

		return contract, nil
	} else {
		logger.Log().WithField("layer", "Init").Errorf("no account contract address")
		return nil, errors.BlankContractAddrErr
	}
}

// getGoerliPerpsMarket is used to get perps market contract instance
func (p *Perpsv3) getPerpsMarket() (*perpsMarket.PerpsMarket, error) {
	if p.config.ContractAddresses.PerpsMarket != "" {
		addr, err := getAddr(p.config.ContractAddresses.PerpsMarket, "perps")
		if err != nil {
			return nil, err
		}

		contract, err := perpsMarket.NewPerpsMarket(addr, p.rpcClient)
		if err != nil {
			logger.Log().WithField("layer", "Init").Errorf("error getting perps market contract: %v", err.Error())
			return nil, errors.GetInitContractErr(err)
		}

		return contract, nil
	} else {
		logger.Log().WithField("layer", "Init").Errorf("no perps market contract address")
		return nil, errors.BlankContractAddrErr
	}
}

// getAddr is used to get contract address as common.Address type
func getAddr(addr string, name string) (common.Address, error) {
	isAddr := common.IsHexAddress(addr)
	if !isAddr {
		logger.Log().WithField("layer", "Init").Errorf("invalid %v contract address", name)
		return common.Address{}, errors.InvalidContractAddrErr
	}

	return common.HexToAddress(addr), nil
}

// TODO: fix test and unmute
// createTest used for testing
//func createTest(conf *config.PerpsvConfig) (*Perpsv3, error) {
//	lib := &Perpsv3{
//		config: conf,
//	}
//
//	if err := lib.init(); err != nil {
//		return nil, err
//	}
//
//	return lib, nil
//}
//	}
//
//	return lib, nil
//}
