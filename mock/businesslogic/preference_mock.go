// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/preference.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIUserPreferenceRepository is a mock of IUserPreferenceRepository interface
type MockIUserPreferenceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserPreferenceRepositoryMockRecorder
}

// MockIUserPreferenceRepositoryMockRecorder is the mock recorder for MockIUserPreferenceRepository
type MockIUserPreferenceRepositoryMockRecorder struct {
	mock *MockIUserPreferenceRepository
}

// NewMockIUserPreferenceRepository creates a new mock instance
func NewMockIUserPreferenceRepository(ctrl *gomock.Controller) *MockIUserPreferenceRepository {
	mock := &MockIUserPreferenceRepository{ctrl: ctrl}
	mock.recorder = &MockIUserPreferenceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUserPreferenceRepository) EXPECT() *MockIUserPreferenceRepositoryMockRecorder {
	return m.recorder
}

// CreatePreference mocks base method
func (m *MockIUserPreferenceRepository) CreatePreference(preference *businesslogic.UserPreference) error {
	ret := m.ctrl.Call(m, "CreatePreference", preference)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePreference indicates an expected call of CreatePreference
func (mr *MockIUserPreferenceRepositoryMockRecorder) CreatePreference(preference interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePreference", reflect.TypeOf((*MockIUserPreferenceRepository)(nil).CreatePreference), preference)
}

// SearchPreference mocks base method
func (m *MockIUserPreferenceRepository) SearchPreference(criteria businesslogic.SearchUserPreferenceCriteria) ([]businesslogic.UserPreference, error) {
	ret := m.ctrl.Call(m, "SearchPreference", criteria)
	ret0, _ := ret[0].([]businesslogic.UserPreference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPreference indicates an expected call of SearchPreference
func (mr *MockIUserPreferenceRepositoryMockRecorder) SearchPreference(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPreference", reflect.TypeOf((*MockIUserPreferenceRepository)(nil).SearchPreference), criteria)
}

// UpdatePreference mocks base method
func (m *MockIUserPreferenceRepository) UpdatePreference(preference businesslogic.UserPreference) error {
	ret := m.ctrl.Call(m, "UpdatePreference", preference)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePreference indicates an expected call of UpdatePreference
func (mr *MockIUserPreferenceRepositoryMockRecorder) UpdatePreference(preference interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePreference", reflect.TypeOf((*MockIUserPreferenceRepository)(nil).UpdatePreference), preference)
}
