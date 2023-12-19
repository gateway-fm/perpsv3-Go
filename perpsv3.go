package perpsv3_Go

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

	// ListenAccountLiquidated is used to listen to all 'AccountLiquidated' contract events and return them as models.AccountLiquidated
	// struct and return errors on ErrChan chanel
	ListenAccountLiquidated() (*events.AccountLiquidatedSubscription, error)

	// ListenAccountPermissionRevoked is used to listen to all 'PermissionRevoked' contract events and return them as models.PermissionChanged
	// struct and return errors on ErrChan chanel
	ListenAccountPermissionRevoked() (*events.AccountPermissionRevokedSubscription, error)

	// ListenAccountPermissionGranted is used to listen to all 'PermissionGranted' contract events and return them as models.PermissionChanged
	// struct and return errors on ErrChan chanel
	ListenAccountPermissionGranted() (*events.AccountPermissionGrantedSubscription, error)

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

	// FormatAccount is used to get account, and it's additional data from the contract by given account id
	FormatAccount(id *big.Int) (*models.Account, error)

	// FormatAccounts is used to get all accounts and their additional data from the contract
	FormatAccounts() ([]*models.Account, error)

	// FormatAccountsLimit is used to get all accounts and their additional data from the contract with given block search
	// limit. For most public RPC providers the value for limit is 20 000 blocks
	FormatAccountsLimit(limit uint64) ([]*models.Account, error)

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

func (p *Perpsv3) ListenAccountLiquidated() (*events.AccountLiquidatedSubscription, error) {
	return p.events.ListenAccountLiquidated()
}

func (p *Perpsv3) ListenAccountPermissionRevoked() (*events.AccountPermissionRevokedSubscription, error) {
	return p.events.ListenAccountPermissionRevoked()
}

func (p *Perpsv3) ListenAccountPermissionGranted() (*events.AccountPermissionGrantedSubscription, error) {
	return p.events.ListenAccountPermissionGranted()
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

func (p *Perpsv3) FormatAccounts() ([]*models.Account, error) {
	return p.service.FormatAccounts()
}

func (p *Perpsv3) FormatAccount(id *big.Int) (*models.Account, error) {
	return p.service.FormatAccount(id)
}

func (p *Perpsv3) FormatAccountsLimit(limit uint64) ([]*models.Account, error) {
	return p.service.FormatAccountsLimit(limit)
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

	srv, err := services.NewService(rpcClient, p.config, coreContact, perpsMarketContract)
	if err != nil {
		return err
	}

	p.service = srv
	p.events = events.NewEvents(rpcClient, coreContact, perpsMarketContract)

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

// createTest used for testing
func createTest(conf *config.PerpsvConfig) (*Perpsv3, error) {
	lib := &Perpsv3{
		config: conf,
	}

	if err := lib.init(); err != nil {
		return nil, err
	}

	return lib, nil
}
