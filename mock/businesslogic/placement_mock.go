// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/placement.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIPlacementRepository is a mock of IPlacementRepository interface
type MockIPlacementRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPlacementRepositoryMockRecorder
}

// MockIPlacementRepositoryMockRecorder is the mock recorder for MockIPlacementRepository
type MockIPlacementRepositoryMockRecorder struct {
	mock *MockIPlacementRepository
}

// NewMockIPlacementRepository creates a new mock instance
func NewMockIPlacementRepository(ctrl *gomock.Controller) *MockIPlacementRepository {
	mock := &MockIPlacementRepository{ctrl: ctrl}
	mock.recorder = &MockIPlacementRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPlacementRepository) EXPECT() *MockIPlacementRepositoryMockRecorder {
	return m.recorder
}

// CreatePlacement mocks base method
func (m *MockIPlacementRepository) CreatePlacement(placement *businesslogic.Placement) error {
	ret := m.ctrl.Call(m, "CreatePlacement", placement)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePlacement indicates an expected call of CreatePlacement
func (mr *MockIPlacementRepositoryMockRecorder) CreatePlacement(placement interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlacement", reflect.TypeOf((*MockIPlacementRepository)(nil).CreatePlacement), placement)
}

// DeletePlacement mocks base method
func (m *MockIPlacementRepository) DeletePlacement(placement businesslogic.Placement) error {
	ret := m.ctrl.Call(m, "DeletePlacement", placement)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlacement indicates an expected call of DeletePlacement
func (mr *MockIPlacementRepositoryMockRecorder) DeletePlacement(placement interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlacement", reflect.TypeOf((*MockIPlacementRepository)(nil).DeletePlacement), placement)
}

// SearchPlacement mocks base method
func (m *MockIPlacementRepository) SearchPlacement(criteria businesslogic.SearchPlacementCriteria) ([]businesslogic.Placement, error) {
	ret := m.ctrl.Call(m, "SearchPlacement", criteria)
	ret0, _ := ret[0].([]businesslogic.Placement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPlacement indicates an expected call of SearchPlacement
func (mr *MockIPlacementRepositoryMockRecorder) SearchPlacement(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPlacement", reflect.TypeOf((*MockIPlacementRepository)(nil).SearchPlacement), criteria)
}

// UpdatePlacement mocks base method
func (m *MockIPlacementRepository) UpdatePlacement(placement businesslogic.Placement) error {
	ret := m.ctrl.Call(m, "UpdatePlacement", placement)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePlacement indicates an expected call of UpdatePlacement
func (mr *MockIPlacementRepositoryMockRecorder) UpdatePlacement(placement interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlacement", reflect.TypeOf((*MockIPlacementRepository)(nil).UpdatePlacement), placement)
}
