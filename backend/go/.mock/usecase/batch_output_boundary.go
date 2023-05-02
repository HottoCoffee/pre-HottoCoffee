// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/batch_output_boundary.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"

	entity "github.com/HottoCoffee/HottoCoffee/core/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockBatchOutputBoundary is a mock of BatchOutputBoundary interface.
type MockBatchOutputBoundary struct {
	ctrl     *gomock.Controller
	recorder *MockBatchOutputBoundaryMockRecorder
}

// MockBatchOutputBoundaryMockRecorder is the mock recorder for MockBatchOutputBoundary.
type MockBatchOutputBoundaryMockRecorder struct {
	mock *MockBatchOutputBoundary
}

// NewMockBatchOutputBoundary creates a new mock instance.
func NewMockBatchOutputBoundary(ctrl *gomock.Controller) *MockBatchOutputBoundary {
	mock := &MockBatchOutputBoundary{ctrl: ctrl}
	mock.recorder = &MockBatchOutputBoundaryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBatchOutputBoundary) EXPECT() *MockBatchOutputBoundaryMockRecorder {
	return m.recorder
}

// SendBatchListResponse mocks base method.
func (m *MockBatchOutputBoundary) SendBatchListResponse(b []entity.Batch) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendBatchListResponse", b)
}

// SendBatchListResponse indicates an expected call of SendBatchListResponse.
func (mr *MockBatchOutputBoundaryMockRecorder) SendBatchListResponse(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBatchListResponse", reflect.TypeOf((*MockBatchOutputBoundary)(nil).SendBatchListResponse), b)
}

// SendBatchResponse mocks base method.
func (m *MockBatchOutputBoundary) SendBatchResponse(b entity.Batch) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendBatchResponse", b)
}

// SendBatchResponse indicates an expected call of SendBatchResponse.
func (mr *MockBatchOutputBoundaryMockRecorder) SendBatchResponse(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBatchResponse", reflect.TypeOf((*MockBatchOutputBoundary)(nil).SendBatchResponse), b)
}

// SendInternalServerErrorResponse mocks base method.
func (m *MockBatchOutputBoundary) SendInternalServerErrorResponse() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendInternalServerErrorResponse")
}

// SendInternalServerErrorResponse indicates an expected call of SendInternalServerErrorResponse.
func (mr *MockBatchOutputBoundaryMockRecorder) SendInternalServerErrorResponse() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendInternalServerErrorResponse", reflect.TypeOf((*MockBatchOutputBoundary)(nil).SendInternalServerErrorResponse))
}

// SendNotFoundResponse mocks base method.
func (m *MockBatchOutputBoundary) SendNotFoundResponse() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendNotFoundResponse")
}

// SendNotFoundResponse indicates an expected call of SendNotFoundResponse.
func (mr *MockBatchOutputBoundaryMockRecorder) SendNotFoundResponse() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNotFoundResponse", reflect.TypeOf((*MockBatchOutputBoundary)(nil).SendNotFoundResponse))
}
