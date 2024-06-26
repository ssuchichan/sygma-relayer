// Code generated by MockGen. DO NOT EDIT.
// Source: ./chains/evm/listener/eventHandlers/event-handler.go

// Package mock_eventHandlers is a generated GoMock package.
package mock_eventHandlers

import (
	context "context"
	big "math/big"
	reflect "reflect"

	events "github.com/ChainSafe/sygma-relayer/chains/evm/calls/events"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
	message "github.com/sygmaprotocol/sygma-core/relayer/message"
)

// MockEventListener is a mock of EventListener interface.
type MockEventListener struct {
	ctrl     *gomock.Controller
	recorder *MockEventListenerMockRecorder
}

// MockEventListenerMockRecorder is the mock recorder for MockEventListener.
type MockEventListenerMockRecorder struct {
	mock *MockEventListener
}

// NewMockEventListener creates a new mock instance.
func NewMockEventListener(ctrl *gomock.Controller) *MockEventListener {
	mock := &MockEventListener{ctrl: ctrl}
	mock.recorder = &MockEventListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventListener) EXPECT() *MockEventListenerMockRecorder {
	return m.recorder
}

// FetchDeposits mocks base method.
func (m *MockEventListener) FetchDeposits(ctx context.Context, address common.Address, startBlock, endBlock *big.Int) ([]*events.Deposit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchDeposits", ctx, address, startBlock, endBlock)
	ret0, _ := ret[0].([]*events.Deposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchDeposits indicates an expected call of FetchDeposits.
func (mr *MockEventListenerMockRecorder) FetchDeposits(ctx, address, startBlock, endBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchDeposits", reflect.TypeOf((*MockEventListener)(nil).FetchDeposits), ctx, address, startBlock, endBlock)
}

// FetchFrostKeygenEvents mocks base method.
func (m *MockEventListener) FetchFrostKeygenEvents(ctx context.Context, address common.Address, startBlock, endBlock *big.Int) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchFrostKeygenEvents", ctx, address, startBlock, endBlock)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFrostKeygenEvents indicates an expected call of FetchFrostKeygenEvents.
func (mr *MockEventListenerMockRecorder) FetchFrostKeygenEvents(ctx, address, startBlock, endBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchFrostKeygenEvents", reflect.TypeOf((*MockEventListener)(nil).FetchFrostKeygenEvents), ctx, address, startBlock, endBlock)
}

// FetchKeygenEvents mocks base method.
func (m *MockEventListener) FetchKeygenEvents(ctx context.Context, address common.Address, startBlock, endBlock *big.Int) ([]types.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchKeygenEvents", ctx, address, startBlock, endBlock)
	ret0, _ := ret[0].([]types.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchKeygenEvents indicates an expected call of FetchKeygenEvents.
func (mr *MockEventListenerMockRecorder) FetchKeygenEvents(ctx, address, startBlock, endBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchKeygenEvents", reflect.TypeOf((*MockEventListener)(nil).FetchKeygenEvents), ctx, address, startBlock, endBlock)
}

// FetchRefreshEvents mocks base method.
func (m *MockEventListener) FetchRefreshEvents(ctx context.Context, address common.Address, startBlock, endBlock *big.Int) ([]*events.Refresh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRefreshEvents", ctx, address, startBlock, endBlock)
	ret0, _ := ret[0].([]*events.Refresh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRefreshEvents indicates an expected call of FetchRefreshEvents.
func (mr *MockEventListenerMockRecorder) FetchRefreshEvents(ctx, address, startBlock, endBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRefreshEvents", reflect.TypeOf((*MockEventListener)(nil).FetchRefreshEvents), ctx, address, startBlock, endBlock)
}

// FetchRetryDepositEvents mocks base method.
func (m *MockEventListener) FetchRetryDepositEvents(event events.RetryEvent, bridgeAddress common.Address, blockConfirmations *big.Int) ([]events.Deposit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRetryDepositEvents", event, bridgeAddress, blockConfirmations)
	ret0, _ := ret[0].([]events.Deposit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRetryDepositEvents indicates an expected call of FetchRetryDepositEvents.
func (mr *MockEventListenerMockRecorder) FetchRetryDepositEvents(event, bridgeAddress, blockConfirmations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRetryDepositEvents", reflect.TypeOf((*MockEventListener)(nil).FetchRetryDepositEvents), event, bridgeAddress, blockConfirmations)
}

// FetchRetryEvents mocks base method.
func (m *MockEventListener) FetchRetryEvents(ctx context.Context, contractAddress common.Address, startBlock, endBlock *big.Int) ([]events.RetryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRetryEvents", ctx, contractAddress, startBlock, endBlock)
	ret0, _ := ret[0].([]events.RetryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRetryEvents indicates an expected call of FetchRetryEvents.
func (mr *MockEventListenerMockRecorder) FetchRetryEvents(ctx, contractAddress, startBlock, endBlock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRetryEvents", reflect.TypeOf((*MockEventListener)(nil).FetchRetryEvents), ctx, contractAddress, startBlock, endBlock)
}

// MockDepositHandler is a mock of DepositHandler interface.
type MockDepositHandler struct {
	ctrl     *gomock.Controller
	recorder *MockDepositHandlerMockRecorder
}

// MockDepositHandlerMockRecorder is the mock recorder for MockDepositHandler.
type MockDepositHandlerMockRecorder struct {
	mock *MockDepositHandler
}

// NewMockDepositHandler creates a new mock instance.
func NewMockDepositHandler(ctrl *gomock.Controller) *MockDepositHandler {
	mock := &MockDepositHandler{ctrl: ctrl}
	mock.recorder = &MockDepositHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDepositHandler) EXPECT() *MockDepositHandlerMockRecorder {
	return m.recorder
}

// HandleDeposit mocks base method.
func (m *MockDepositHandler) HandleDeposit(sourceID, destID uint8, nonce uint64, resourceID [32]byte, calldata, handlerResponse []byte, messageID string) (*message.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDeposit", sourceID, destID, nonce, resourceID, calldata, handlerResponse, messageID)
	ret0, _ := ret[0].(*message.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleDeposit indicates an expected call of HandleDeposit.
func (mr *MockDepositHandlerMockRecorder) HandleDeposit(sourceID, destID, nonce, resourceID, calldata, handlerResponse, messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDeposit", reflect.TypeOf((*MockDepositHandler)(nil).HandleDeposit), sourceID, destID, nonce, resourceID, calldata, handlerResponse, messageID)
}
