// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/andrei3131/hero-wars/game (interfaces: StrikeEngineInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	player "github.com/andrei3131/hero-wars/player"
	gomock "github.com/golang/mock/gomock"
	rand "math/rand"
	reflect "reflect"
)

// MockStrikeEngineInterface is a mock of StrikeEngineInterface interface
type MockStrikeEngineInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStrikeEngineInterfaceMockRecorder
}

// MockStrikeEngineInterfaceMockRecorder is the mock recorder for MockStrikeEngineInterface
type MockStrikeEngineInterfaceMockRecorder struct {
	mock *MockStrikeEngineInterface
}

// NewMockStrikeEngineInterface creates a new mock instance
func NewMockStrikeEngineInterface(ctrl *gomock.Controller) *MockStrikeEngineInterface {
	mock := &MockStrikeEngineInterface{ctrl: ctrl}
	mock.recorder = &MockStrikeEngineInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStrikeEngineInterface) EXPECT() *MockStrikeEngineInterfaceMockRecorder {
	return m.recorder
}

// Attack mocks base method
func (m *MockStrikeEngineInterface) Attack(arg0, arg1 string, arg2 *rand.Rand) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Attack", arg0, arg1, arg2)
}

// Attack indicates an expected call of Attack
func (mr *MockStrikeEngineInterfaceMockRecorder) Attack(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attack", reflect.TypeOf((*MockStrikeEngineInterface)(nil).Attack), arg0, arg1, arg2)
}

// IsFirstTurnPlayed mocks base method
func (m *MockStrikeEngineInterface) IsFirstTurnPlayed() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFirstTurnPlayed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFirstTurnPlayed indicates an expected call of IsFirstTurnPlayed
func (mr *MockStrikeEngineInterfaceMockRecorder) IsFirstTurnPlayed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFirstTurnPlayed", reflect.TypeOf((*MockStrikeEngineInterface)(nil).IsFirstTurnPlayed))
}

// SetFirstTurnPlayed mocks base method
func (m *MockStrikeEngineInterface) SetFirstTurnPlayed() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFirstTurnPlayed")
}

// SetFirstTurnPlayed indicates an expected call of SetFirstTurnPlayed
func (mr *MockStrikeEngineInterfaceMockRecorder) SetFirstTurnPlayed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFirstTurnPlayed", reflect.TypeOf((*MockStrikeEngineInterface)(nil).SetFirstTurnPlayed))
}

// UpdateStrikeEngine mocks base method
func (m *MockStrikeEngineInterface) UpdateStrikeEngine(arg0 *player.Player, arg1 *player.Special, arg2 *player.Player, arg3 *player.Special) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateStrikeEngine", arg0, arg1, arg2, arg3)
}

// UpdateStrikeEngine indicates an expected call of UpdateStrikeEngine
func (mr *MockStrikeEngineInterfaceMockRecorder) UpdateStrikeEngine(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStrikeEngine", reflect.TypeOf((*MockStrikeEngineInterface)(nil).UpdateStrikeEngine), arg0, arg1, arg2, arg3)
}
