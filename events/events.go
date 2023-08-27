package events

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/gateway-fm/perpsv3-Go/contracts/coreGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/perpsMarketGoerli"
	"github.com/gateway-fm/perpsv3-Go/contracts/spotMarketGoerli"
)

// IEvents is an interface that is used to work with contract event listeners
type IEvents interface {
	// ListenTrades is used to listen to all 'OrderSettled' contract events and return them as models.Trade struct and
	// return errors on ErrChan chanel
	ListenTrades() (*TradeSubscription, error)
}

// Events implements IEvents interface
type Events struct {
	rpcClient   *ethclient.Client
	core        *coreGoerli.CoreGoerli
	spotMarket  *spotMarketGoerli.SpotMarketGoerli
	perpsMarket *perpsMarketGoerli.PerpsMarketGoerli
}

// NewEvents is used to create new Events instance that implements IEvents interface
func NewEvents(
	client *ethclient.Client,
	core *coreGoerli.CoreGoerli,
	spotMarket *spotMarketGoerli.SpotMarketGoerli,
	perpsMarket *perpsMarketGoerli.PerpsMarketGoerli,
) IEvents {
	return &Events{
		rpcClient:   client,
		core:        core,
		spotMarket:  spotMarket,
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
	s.eventSub.Unsubscribe()
	close(s.stop)
	close(s.ErrChan)
}
