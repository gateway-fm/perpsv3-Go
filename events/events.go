package events

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"

	"github.com/gateway-fm/perpsv3-Go/contracts/core"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarket"
	"github.com/gateway-fm/perpsv3-Go/pkg/logger"
)

// IEvents is an interface that is used to work with contract event listeners
type IEvents interface {
	// ListenTrades is used to listen to all 'OrderSettled' contract events and return them as models.Trade struct and
	// return errors on ErrChan chanel
	ListenTrades() (*TradeSubscription, error)

	// ListenOrders is used to listen to all 'OrderCommitted' contract events and return them as models.Order struct and
	// return errors on ErrChan chanel
	ListenOrders() (*OrderSubscription, error)

	// ListenMarketUpdates is used to listen to all 'MarketUpdated' contract events and return them as models.MarketUpdate
	// struct and return errors on ErrChan chanel
	ListenMarketUpdates() (*MarketUpdateSubscription, error)

	// ListenMarketUpdatesBig is used to listen to all 'MarketUpdated' contract events and return them as models.MarketUpdateBig
	// struct and return errors on ErrChan chanel
	ListenMarketUpdatesBig() (*MarketUpdateSubscriptionBig, error)

	// ListenLiquidations is used to listen to all 'PositionLiquidated' contract events and return them as models.Liquidation
	// struct and return errors on ErrChan chanel
	ListenLiquidations() (*LiquidationSubscription, error)

	// ListenAccountCreated is used to listen to all 'AccountCreated' contract events and return them as models.Account
	// struct and return errors on ErrChan chanel
	ListenAccountCreated() (*AccountCreatedSubscription, error)

	// ListenAccountLiquidated is used to listen to all 'AccountLiquidated' contract events and return them as models.AccountLiquidated
	// struct and return errors on ErrChan chanel
	ListenAccountLiquidated() (*AccountLiquidatedSubscription, error)

	// ListenAccountPermissionRevoked is used to listen to all 'PermissionRevoked' contract events and return them as models.PermissionChanged
	// struct and return errors on ErrChan chanel
	ListenAccountPermissionRevoked() (*AccountPermissionRevokedSubscription, error)

	// ListenAccountPermissionGranted is used to listen to all 'PermissionGranted' contract events and return them as models.PermissionChanged
	// struct and return errors on ErrChan chanel
	ListenAccountPermissionGranted() (*AccountPermissionGrantedSubscription, error)

	// ListenUSDMinted is used to listen to all 'USDMinted' Core contract events and return them as models.USDMinted
	// struct and return errors on ErrChan chanel
	ListenUSDMinted() (*USDMintedSubscription, error)
}

// Events implements IEvents interface
type Events struct {
	rpcClient   *ethclient.Client
	core        *core.Core
	perpsMarket *perpsMarket.PerpsMarket
}

// NewEvents is used to create new Events instance that implements IEvents interface
func NewEvents(
	client *ethclient.Client,
	core *core.Core,
	perpsMarket *perpsMarket.PerpsMarket,
) IEvents {
	return &Events{
		rpcClient:   client,
		core:        core,
		perpsMarket: perpsMarket,
	}
}

// basicSubscription is an event subscription struct with two channels:
//   - stop is a `stop` signal chanel
//   - eventSub is an event subscription chanel to handle errors
type basicSubscription struct {
	stop     chan struct{}
	ErrChan  chan error
	eventSub event.Subscription
}

// newBasicSubscription returns new basicSubscription instance
func newBasicSubscription(eventSub event.Subscription) *basicSubscription {
	return &basicSubscription{
		stop:     make(chan struct{}),
		ErrChan:  make(chan error),
		eventSub: eventSub,
	}
}

// Close is used to stop the event chanel and send the `stop` signal
func (s *basicSubscription) Close() {
	close(s.stop)
	close(s.ErrChan)
	s.eventSub.Unsubscribe()
}

// getAccountLastInteraction
func getAccountLastInteraction(id *big.Int, perpsMarket *perpsMarket.PerpsMarket) *big.Int {
	lastInteraction, err := perpsMarket.GetAccountLastInteraction(nil, id)
	if err != nil {
		logger.Log().WithField("layer", "Events-Accounts").Errorf(
			"error query account last interaction: %v, last interaction set to 0", err.Error(),
		)
		lastInteraction = big.NewInt(0)
	}

	return lastInteraction
}
