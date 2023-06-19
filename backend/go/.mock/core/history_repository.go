// Code generated by MockGen. DO NOT EDIT.
// Source: core/history_repository.go

// Package mock_core is a generated GoMock package.
package mock_core

import (
	reflect "reflect"
	time "time"

	entity "github.com/HottoCoffee/HottoCoffee/core/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockHistoryRepository is a mock of HistoryRepository interface.
type MockHistoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockHistoryRepositoryMockRecorder
}

// MockHistoryRepositoryMockRecorder is the mock recorder for MockHistoryRepository.
type MockHistoryRepositoryMockRecorder struct {
	mock *MockHistoryRepository
}

// NewMockHistoryRepository creates a new mock instance.
func NewMockHistoryRepository(ctrl *gomock.Controller) *MockHistoryRepository {
	mock := &MockHistoryRepository{ctrl: ctrl}
	mock.recorder = &MockHistoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHistoryRepository) EXPECT() *MockHistoryRepositoryMockRecorder {
	return m.recorder
}

// FindAllDuring mocks base method.
func (m *MockHistoryRepository) FindAllDuring(startDate, endDate time.Time) ([]entity.BatchExecutionHistories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllDuring", startDate, endDate)
	ret0, _ := ret[0].([]entity.BatchExecutionHistories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllDuring indicates an expected call of FindAllDuring.
func (mr *MockHistoryRepositoryMockRecorder) FindAllDuring(startDate, endDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllDuring", reflect.TypeOf((*MockHistoryRepository)(nil).FindAllDuring), startDate, endDate)
}

// FindByBatchId mocks base method.
func (m *MockHistoryRepository) FindByBatchId(batchId int) (*entity.BatchExecutionHistories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByBatchId", batchId)
	ret0, _ := ret[0].(*entity.BatchExecutionHistories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByBatchId indicates an expected call of FindByBatchId.
func (mr *MockHistoryRepositoryMockRecorder) FindByBatchId(batchId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByBatchId", reflect.TypeOf((*MockHistoryRepository)(nil).FindByBatchId), batchId)
}

// FindByHistoryIdAndBatchId mocks base method.
func (m *MockHistoryRepository) FindByHistoryIdAndBatchId(historyId, batchId int) (*entity.BatchExecutionHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByHistoryIdAndBatchId", historyId, batchId)
	ret0, _ := ret[0].(*entity.BatchExecutionHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByHistoryIdAndBatchId indicates an expected call of FindByHistoryIdAndBatchId.
func (mr *MockHistoryRepositoryMockRecorder) FindByHistoryIdAndBatchId(historyId, batchId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByHistoryIdAndBatchId", reflect.TypeOf((*MockHistoryRepository)(nil).FindByHistoryIdAndBatchId), historyId, batchId)
}
