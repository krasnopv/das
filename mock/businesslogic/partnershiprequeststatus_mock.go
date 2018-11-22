// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/partnershiprequeststatus.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIPartnershipRequestStatusRepository is a mock of IPartnershipRequestStatusRepository interface
type MockIPartnershipRequestStatusRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPartnershipRequestStatusRepositoryMockRecorder
}

// MockIPartnershipRequestStatusRepositoryMockRecorder is the mock recorder for MockIPartnershipRequestStatusRepository
type MockIPartnershipRequestStatusRepositoryMockRecorder struct {
	mock *MockIPartnershipRequestStatusRepository
}

// NewMockIPartnershipRequestStatusRepository creates a new mock instance
func NewMockIPartnershipRequestStatusRepository(ctrl *gomock.Controller) *MockIPartnershipRequestStatusRepository {
	mock := &MockIPartnershipRequestStatusRepository{ctrl: ctrl}
	mock.recorder = &MockIPartnershipRequestStatusRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPartnershipRequestStatusRepository) EXPECT() *MockIPartnershipRequestStatusRepositoryMockRecorder {
	return m.recorder
}

// GetPartnershipRequestStatus mocks base method
func (m *MockIPartnershipRequestStatusRepository) GetPartnershipRequestStatus() ([]businesslogic.PartnershipRequestStatus, error) {
	ret := m.ctrl.Call(m, "GetPartnershipRequestStatus")
	ret0, _ := ret[0].([]businesslogic.PartnershipRequestStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartnershipRequestStatus indicates an expected call of GetPartnershipRequestStatus
func (mr *MockIPartnershipRequestStatusRepositoryMockRecorder) GetPartnershipRequestStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartnershipRequestStatus", reflect.TypeOf((*MockIPartnershipRequestStatusRepository)(nil).GetPartnershipRequestStatus))
}