// Code generated by MockGen. DO NOT EDIT.
// Source: ./models/models.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPersistence is a mock of Persistence interface
type MockPersistence struct {
	ctrl     *gomock.Controller
	recorder *MockPersistenceMockRecorder
}

// MockPersistenceMockRecorder is the mock recorder for MockPersistence
type MockPersistenceMockRecorder struct {
	mock *MockPersistence
}

// NewMockPersistence creates a new mock instance
func NewMockPersistence(ctrl *gomock.Controller) *MockPersistence {
	mock := &MockPersistence{ctrl: ctrl}
	mock.recorder = &MockPersistenceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPersistence) EXPECT() *MockPersistenceMockRecorder {
	return m.recorder
}

// Put mocks base method
func (m *MockPersistence) Put(key string, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockPersistenceMockRecorder) Put(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockPersistence)(nil).Put), key, value)
}

// Get mocks base method
func (m *MockPersistence) Get(key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPersistenceMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPersistence)(nil).Get), key)
}

// MockCacheInterface is a mock of CacheInterface interface
type MockCacheInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCacheInterfaceMockRecorder
}

// MockCacheInterfaceMockRecorder is the mock recorder for MockCacheInterface
type MockCacheInterfaceMockRecorder struct {
	mock *MockCacheInterface
}

// NewMockCacheInterface creates a new mock instance
func NewMockCacheInterface(ctrl *gomock.Controller) *MockCacheInterface {
	mock := &MockCacheInterface{ctrl: ctrl}
	mock.recorder = &MockCacheInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCacheInterface) EXPECT() *MockCacheInterfaceMockRecorder {
	return m.recorder
}

// Put mocks base method
func (m *MockCacheInterface) Put(key string, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockCacheInterfaceMockRecorder) Put(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockCacheInterface)(nil).Put), key, value)
}

// Get mocks base method
func (m *MockCacheInterface) Get(key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCacheInterfaceMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCacheInterface)(nil).Get), key)
}

// Flush mocks base method
func (m *MockCacheInterface) Flush(key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush", key)
}

// Flush indicates an expected call of Flush
func (mr *MockCacheInterfaceMockRecorder) Flush(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockCacheInterface)(nil).Flush), key)
}
