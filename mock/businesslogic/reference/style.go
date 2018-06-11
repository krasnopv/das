// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/referencedal/style.go

// Package mock_reference is a generated GoMock package.
package mock_reference

import (
	reference "github.com/DancesportSoftware/das/businesslogic/reference"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIStyleRepository is a mock of IStyleRepository interface
type MockIStyleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIStyleRepositoryMockRecorder
}

// MockIStyleRepositoryMockRecorder is the mock recorder for MockIStyleRepository
type MockIStyleRepositoryMockRecorder struct {
	mock *MockIStyleRepository
}

// NewMockIStyleRepository creates a new mock instance
func NewMockIStyleRepository(ctrl *gomock.Controller) *MockIStyleRepository {
	mock := &MockIStyleRepository{ctrl: ctrl}
	mock.recorder = &MockIStyleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIStyleRepository) EXPECT() *MockIStyleRepositoryMockRecorder {
	return m.recorder
}

// CreateStyle mocks base method
func (m *MockIStyleRepository) CreateStyle(style *reference.Style) error {
	ret := m.ctrl.Call(m, "CreateStyle", style)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStyle indicates an expected call of CreateStyle
func (mr *MockIStyleRepositoryMockRecorder) CreateStyle(style interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStyle", reflect.TypeOf((*MockIStyleRepository)(nil).CreateStyle), style)
}

// SearchStyle mocks base method
func (m *MockIStyleRepository) SearchStyle(criteria reference.SearchStyleCriteria) ([]reference.Style, error) {
	ret := m.ctrl.Call(m, "SearchStyle", criteria)
	ret0, _ := ret[0].([]reference.Style)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchStyle indicates an expected call of SearchStyle
func (mr *MockIStyleRepositoryMockRecorder) SearchStyle(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchStyle", reflect.TypeOf((*MockIStyleRepository)(nil).SearchStyle), criteria)
}

// UpdateStyle mocks base method
func (m *MockIStyleRepository) UpdateStyle(style reference.Style) error {
	ret := m.ctrl.Call(m, "UpdateStyle", style)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStyle indicates an expected call of UpdateStyle
func (mr *MockIStyleRepositoryMockRecorder) UpdateStyle(style interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStyle", reflect.TypeOf((*MockIStyleRepository)(nil).UpdateStyle), style)
}

// DeleteStyle mocks base method
func (m *MockIStyleRepository) DeleteStyle(style reference.Style) error {
	ret := m.ctrl.Call(m, "DeleteStyle", style)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStyle indicates an expected call of DeleteStyle
func (mr *MockIStyleRepositoryMockRecorder) DeleteStyle(style interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStyle", reflect.TypeOf((*MockIStyleRepository)(nil).DeleteStyle), style)
}
