// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/accountrole.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIAccountRoleRepository is a mock of IAccountRoleRepository interface
type MockIAccountRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountRoleRepositoryMockRecorder
}

// MockIAccountRoleRepositoryMockRecorder is the mock recorder for MockIAccountRoleRepository
type MockIAccountRoleRepositoryMockRecorder struct {
	mock *MockIAccountRoleRepository
}

// NewMockIAccountRoleRepository creates a new mock instance
func NewMockIAccountRoleRepository(ctrl *gomock.Controller) *MockIAccountRoleRepository {
	mock := &MockIAccountRoleRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountRoleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAccountRoleRepository) EXPECT() *MockIAccountRoleRepositoryMockRecorder {
	return m.recorder
}

// CreateAccountRole mocks base method
func (m *MockIAccountRoleRepository) CreateAccountRole(role *businesslogic.AccountRole) error {
	ret := m.ctrl.Call(m, "CreateAccountRole", role)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAccountRole indicates an expected call of CreateAccountRole
func (mr *MockIAccountRoleRepositoryMockRecorder) CreateAccountRole(role interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccountRole", reflect.TypeOf((*MockIAccountRoleRepository)(nil).CreateAccountRole), role)
}

// SearchAccountRole mocks base method
func (m *MockIAccountRoleRepository) SearchAccountRole(criteria businesslogic.SearchAccountRoleCriteria) ([]businesslogic.AccountRole, error) {
	ret := m.ctrl.Call(m, "SearchAccountRole", criteria)
	ret0, _ := ret[0].([]businesslogic.AccountRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAccountRole indicates an expected call of SearchAccountRole
func (mr *MockIAccountRoleRepositoryMockRecorder) SearchAccountRole(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAccountRole", reflect.TypeOf((*MockIAccountRoleRepository)(nil).SearchAccountRole), criteria)
}
