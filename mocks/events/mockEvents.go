// Code generated by MockGen. DO NOT EDIT.
// Source: events/events.go

// Package mock_events is a generated GoMock package.
package mock_events

import (
	reflect "reflect"

	events "github.com/gateway-fm/perpsv3-Go/events"
	gomock "github.com/golang/mock/gomock"
)

// MockIEvents is a mock of IEvents interface.
type MockIEvents struct {
	ctrl     *gomock.Controller
	recorder *MockIEventsMockRecorder
}

// MockIEventsMockRecorder is the mock recorder for MockIEvents.
type MockIEventsMockRecorder struct {
	mock *MockIEvents
}

// NewMockIEvents creates a new mock instance.
func NewMockIEvents(ctrl *gomock.Controller) *MockIEvents {
	mock := &MockIEvents{ctrl: ctrl}
	mock.recorder = &MockIEventsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEvents) EXPECT() *MockIEventsMockRecorder {
	return m.recorder
}

// ListenAccountCorePermissionGranted mocks base method.
func (m *MockIEvents) ListenAccountCorePermissionGranted() (*events.AccountCorePermissionGrantedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAccountCorePermissionGranted")
	ret0, _ := ret[0].(*events.AccountCorePermissionGrantedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenAccountCorePermissionGranted indicates an expected call of ListenAccountCorePermissionGranted.
func (mr *MockIEventsMockRecorder) ListenAccountCorePermissionGranted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAccountCorePermissionGranted", reflect.TypeOf((*MockIEvents)(nil).ListenAccountCorePermissionGranted))
}

// ListenAccountCorePermissionRevoked mocks base method.
func (m *MockIEvents) ListenAccountCorePermissionRevoked() (*events.AccountCorePermissionRevokedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAccountCorePermissionRevoked")
	ret0, _ := ret[0].(*events.AccountCorePermissionRevokedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenAccountCorePermissionRevoked indicates an expected call of ListenAccountCorePermissionRevoked.
func (mr *MockIEventsMockRecorder) ListenAccountCorePermissionRevoked() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAccountCorePermissionRevoked", reflect.TypeOf((*MockIEvents)(nil).ListenAccountCorePermissionRevoked))
}

// ListenAccountCreated mocks base method.
func (m *MockIEvents) ListenAccountCreated() (*events.AccountCreatedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAccountCreated")
	ret0, _ := ret[0].(*events.AccountCreatedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenAccountCreated indicates an expected call of ListenAccountCreated.
func (mr *MockIEventsMockRecorder) ListenAccountCreated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAccountCreated", reflect.TypeOf((*MockIEvents)(nil).ListenAccountCreated))
}

// ListenAccountCreatedCore mocks base method.
func (m *MockIEvents) ListenAccountCreatedCore() (*events.AccountCreatedCoreSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAccountCreatedCore")
	ret0, _ := ret[0].(*events.AccountCreatedCoreSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenAccountCreatedCore indicates an expected call of ListenAccountCreatedCore.
func (mr *MockIEventsMockRecorder) ListenAccountCreatedCore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAccountCreatedCore", reflect.TypeOf((*MockIEvents)(nil).ListenAccountCreatedCore))
}

// ListenAccountLiquidated mocks base method.
func (m *MockIEvents) ListenAccountLiquidated() (*events.AccountLiquidatedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAccountLiquidated")
	ret0, _ := ret[0].(*events.AccountLiquidatedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenAccountLiquidated indicates an expected call of ListenAccountLiquidated.
func (mr *MockIEventsMockRecorder) ListenAccountLiquidated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAccountLiquidated", reflect.TypeOf((*MockIEvents)(nil).ListenAccountLiquidated))
}

// ListenAccountPermissionGranted mocks base method.
func (m *MockIEvents) ListenAccountPermissionGranted() (*events.AccountPermissionGrantedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAccountPermissionGranted")
	ret0, _ := ret[0].(*events.AccountPermissionGrantedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenAccountPermissionGranted indicates an expected call of ListenAccountPermissionGranted.
func (mr *MockIEventsMockRecorder) ListenAccountPermissionGranted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAccountPermissionGranted", reflect.TypeOf((*MockIEvents)(nil).ListenAccountPermissionGranted))
}

// ListenAccountPermissionRevoked mocks base method.
func (m *MockIEvents) ListenAccountPermissionRevoked() (*events.AccountPermissionRevokedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAccountPermissionRevoked")
	ret0, _ := ret[0].(*events.AccountPermissionRevokedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenAccountPermissionRevoked indicates an expected call of ListenAccountPermissionRevoked.
func (mr *MockIEventsMockRecorder) ListenAccountPermissionRevoked() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAccountPermissionRevoked", reflect.TypeOf((*MockIEvents)(nil).ListenAccountPermissionRevoked))
}

// ListenAccountTransfer mocks base method.
func (m *MockIEvents) ListenAccountTransfer() (*events.AccountTransferSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAccountTransfer")
	ret0, _ := ret[0].(*events.AccountTransferSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenAccountTransfer indicates an expected call of ListenAccountTransfer.
func (mr *MockIEventsMockRecorder) ListenAccountTransfer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAccountTransfer", reflect.TypeOf((*MockIEvents)(nil).ListenAccountTransfer))
}

// ListenCollateralDeposited mocks base method.
func (m *MockIEvents) ListenCollateralDeposited() (*events.CollateralDepositedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenCollateralDeposited")
	ret0, _ := ret[0].(*events.CollateralDepositedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenCollateralDeposited indicates an expected call of ListenCollateralDeposited.
func (mr *MockIEventsMockRecorder) ListenCollateralDeposited() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenCollateralDeposited", reflect.TypeOf((*MockIEvents)(nil).ListenCollateralDeposited))
}

// ListenCollateralWithdrawn mocks base method.
func (m *MockIEvents) ListenCollateralWithdrawn() (*events.CollateralWithdrawnSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenCollateralWithdrawn")
	ret0, _ := ret[0].(*events.CollateralWithdrawnSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenCollateralWithdrawn indicates an expected call of ListenCollateralWithdrawn.
func (mr *MockIEventsMockRecorder) ListenCollateralWithdrawn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenCollateralWithdrawn", reflect.TypeOf((*MockIEvents)(nil).ListenCollateralWithdrawn))
}

// ListenDelegationUpdated mocks base method.
func (m *MockIEvents) ListenDelegationUpdated() (*events.DelegationUpdatedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenDelegationUpdated")
	ret0, _ := ret[0].(*events.DelegationUpdatedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenDelegationUpdated indicates an expected call of ListenDelegationUpdated.
func (mr *MockIEventsMockRecorder) ListenDelegationUpdated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenDelegationUpdated", reflect.TypeOf((*MockIEvents)(nil).ListenDelegationUpdated))
}

// ListenLiquidations mocks base method.
func (m *MockIEvents) ListenLiquidations() (*events.LiquidationSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenLiquidations")
	ret0, _ := ret[0].(*events.LiquidationSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenLiquidations indicates an expected call of ListenLiquidations.
func (mr *MockIEventsMockRecorder) ListenLiquidations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenLiquidations", reflect.TypeOf((*MockIEvents)(nil).ListenLiquidations))
}

// ListenLiquidationsCore mocks base method.
func (m *MockIEvents) ListenLiquidationsCore() (*events.LiquidationsCoreSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenLiquidationsCore")
	ret0, _ := ret[0].(*events.LiquidationsCoreSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenLiquidationsCore indicates an expected call of ListenLiquidationsCore.
func (mr *MockIEventsMockRecorder) ListenLiquidationsCore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenLiquidationsCore", reflect.TypeOf((*MockIEvents)(nil).ListenLiquidationsCore))
}

// ListenMarketRegistered mocks base method.
func (m *MockIEvents) ListenMarketRegistered() (*events.MarketRegisteredSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenMarketRegistered")
	ret0, _ := ret[0].(*events.MarketRegisteredSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenMarketRegistered indicates an expected call of ListenMarketRegistered.
func (mr *MockIEventsMockRecorder) ListenMarketRegistered() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenMarketRegistered", reflect.TypeOf((*MockIEvents)(nil).ListenMarketRegistered))
}

// ListenMarketUSDDeposited mocks base method.
func (m *MockIEvents) ListenMarketUSDDeposited() (*events.MarketUSDDepositedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenMarketUSDDeposited")
	ret0, _ := ret[0].(*events.MarketUSDDepositedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenMarketUSDDeposited indicates an expected call of ListenMarketUSDDeposited.
func (mr *MockIEventsMockRecorder) ListenMarketUSDDeposited() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenMarketUSDDeposited", reflect.TypeOf((*MockIEvents)(nil).ListenMarketUSDDeposited))
}

// ListenMarketUSDWithdrawn mocks base method.
func (m *MockIEvents) ListenMarketUSDWithdrawn() (*events.MarketUSDWithdrawnSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenMarketUSDWithdrawn")
	ret0, _ := ret[0].(*events.MarketUSDWithdrawnSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenMarketUSDWithdrawn indicates an expected call of ListenMarketUSDWithdrawn.
func (mr *MockIEventsMockRecorder) ListenMarketUSDWithdrawn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenMarketUSDWithdrawn", reflect.TypeOf((*MockIEvents)(nil).ListenMarketUSDWithdrawn))
}

// ListenMarketUpdates mocks base method.
func (m *MockIEvents) ListenMarketUpdates() (*events.MarketUpdateSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenMarketUpdates")
	ret0, _ := ret[0].(*events.MarketUpdateSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenMarketUpdates indicates an expected call of ListenMarketUpdates.
func (mr *MockIEventsMockRecorder) ListenMarketUpdates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenMarketUpdates", reflect.TypeOf((*MockIEvents)(nil).ListenMarketUpdates))
}

// ListenMarketUpdatesBig mocks base method.
func (m *MockIEvents) ListenMarketUpdatesBig() (*events.MarketUpdateSubscriptionBig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenMarketUpdatesBig")
	ret0, _ := ret[0].(*events.MarketUpdateSubscriptionBig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenMarketUpdatesBig indicates an expected call of ListenMarketUpdatesBig.
func (mr *MockIEventsMockRecorder) ListenMarketUpdatesBig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenMarketUpdatesBig", reflect.TypeOf((*MockIEvents)(nil).ListenMarketUpdatesBig))
}

// ListenOrders mocks base method.
func (m *MockIEvents) ListenOrders() (*events.OrderSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenOrders")
	ret0, _ := ret[0].(*events.OrderSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenOrders indicates an expected call of ListenOrders.
func (mr *MockIEventsMockRecorder) ListenOrders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenOrders", reflect.TypeOf((*MockIEvents)(nil).ListenOrders))
}

// ListenPoolCreated mocks base method.
func (m *MockIEvents) ListenPoolCreated() (*events.PoolCreatedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenPoolCreated")
	ret0, _ := ret[0].(*events.PoolCreatedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenPoolCreated indicates an expected call of ListenPoolCreated.
func (mr *MockIEventsMockRecorder) ListenPoolCreated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenPoolCreated", reflect.TypeOf((*MockIEvents)(nil).ListenPoolCreated))
}

// ListenRewardClaimed mocks base method.
func (m *MockIEvents) ListenRewardClaimed() (*events.RewardClaimedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenRewardClaimed")
	ret0, _ := ret[0].(*events.RewardClaimedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenRewardClaimed indicates an expected call of ListenRewardClaimed.
func (mr *MockIEventsMockRecorder) ListenRewardClaimed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenRewardClaimed", reflect.TypeOf((*MockIEvents)(nil).ListenRewardClaimed))
}

// ListenRewardDistributed mocks base method.
func (m *MockIEvents) ListenRewardDistributed() (*events.RewardDistributedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenRewardDistributed")
	ret0, _ := ret[0].(*events.RewardDistributedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenRewardDistributed indicates an expected call of ListenRewardDistributed.
func (mr *MockIEventsMockRecorder) ListenRewardDistributed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenRewardDistributed", reflect.TypeOf((*MockIEvents)(nil).ListenRewardDistributed))
}

// ListenTrades mocks base method.
func (m *MockIEvents) ListenTrades() (*events.TradeSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenTrades")
	ret0, _ := ret[0].(*events.TradeSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenTrades indicates an expected call of ListenTrades.
func (mr *MockIEventsMockRecorder) ListenTrades() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenTrades", reflect.TypeOf((*MockIEvents)(nil).ListenTrades))
}

// ListenUSDBurned mocks base method.
func (m *MockIEvents) ListenUSDBurned() (*events.USDBurnedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenUSDBurned")
	ret0, _ := ret[0].(*events.USDBurnedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenUSDBurned indicates an expected call of ListenUSDBurned.
func (mr *MockIEventsMockRecorder) ListenUSDBurned() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenUSDBurned", reflect.TypeOf((*MockIEvents)(nil).ListenUSDBurned))
}

// ListenUSDMinted mocks base method.
func (m *MockIEvents) ListenUSDMinted() (*events.USDMintedSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenUSDMinted")
	ret0, _ := ret[0].(*events.USDMintedSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenUSDMinted indicates an expected call of ListenUSDMinted.
func (mr *MockIEventsMockRecorder) ListenUSDMinted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenUSDMinted", reflect.TypeOf((*MockIEvents)(nil).ListenUSDMinted))
}

// ListenVaultLiquidationsCore mocks base method.
func (m *MockIEvents) ListenVaultLiquidationsCore() (*events.VaultLiquidationsCoreSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenVaultLiquidationsCore")
	ret0, _ := ret[0].(*events.VaultLiquidationsCoreSubscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenVaultLiquidationsCore indicates an expected call of ListenVaultLiquidationsCore.
func (mr *MockIEventsMockRecorder) ListenVaultLiquidationsCore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenVaultLiquidationsCore", reflect.TypeOf((*MockIEvents)(nil).ListenVaultLiquidationsCore))
}
