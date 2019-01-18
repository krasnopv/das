// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/city.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICityRepository is a mock of ICityRepository interface
type MockICityRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICityRepositoryMockRecorder
}

// MockICityRepositoryMockRecorder is the mock recorder for MockICityRepository
type MockICityRepositoryMockRecorder struct {
	mock *MockICityRepository
}

// NewMockICityRepository creates a new mock instance
func NewMockICityRepository(ctrl *gomock.Controller) *MockICityRepository {
	mock := &MockICityRepository{ctrl: ctrl}
	mock.recorder = &MockICityRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICityRepository) EXPECT() *MockICityRepositoryMockRecorder {
	return m.recorder
}

// CreateCity mocks base method
func (m *MockICityRepository) CreateCity(city *businesslogic.City) error {
	ret := m.ctrl.Call(m, "CreateCity", city)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCity indicates an expected call of CreateCity
func (mr *MockICityRepositoryMockRecorder) CreateCity(city interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCity", reflect.TypeOf((*MockICityRepository)(nil).CreateCity), city)
}

// SearchCity mocks base method
func (m *MockICityRepository) SearchCity(criteria businesslogic.SearchCityCriteria) ([]businesslogic.City, error) {
	ret := m.ctrl.Call(m, "SearchCity", criteria)
	ret0, _ := ret[0].([]businesslogic.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCity indicates an expected call of SearchCity
func (mr *MockICityRepositoryMockRecorder) SearchCity(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCity", reflect.TypeOf((*MockICityRepository)(nil).SearchCity), criteria)
}

// UpdateCity mocks base method
func (m *MockICityRepository) UpdateCity(city businesslogic.City) error {
	ret := m.ctrl.Call(m, "UpdateCity", city)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCity indicates an expected call of UpdateCity
func (mr *MockICityRepositoryMockRecorder) UpdateCity(city interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCity", reflect.TypeOf((*MockICityRepository)(nil).UpdateCity), city)
}

// DeleteCity mocks base method
func (m *MockICityRepository) DeleteCity(city businesslogic.City) error {
	ret := m.ctrl.Call(m, "DeleteCity", city)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCity indicates an expected call of DeleteCity
func (mr *MockICityRepositoryMockRecorder) DeleteCity(city interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCity", reflect.TypeOf((*MockICityRepository)(nil).DeleteCity), city)
}