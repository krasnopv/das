// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/referencedal/proficiency.go

// Package mock_reference is a generated GoMock package.
package mock_reference

import (
	reference "github.com/DancesportSoftware/das/businesslogic/reference"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIProficiencyRepository is a mock of IProficiencyRepository interface
type MockIProficiencyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIProficiencyRepositoryMockRecorder
}

// MockIProficiencyRepositoryMockRecorder is the mock recorder for MockIProficiencyRepository
type MockIProficiencyRepositoryMockRecorder struct {
	mock *MockIProficiencyRepository
}

// NewMockIProficiencyRepository creates a new mock instance
func NewMockIProficiencyRepository(ctrl *gomock.Controller) *MockIProficiencyRepository {
	mock := &MockIProficiencyRepository{ctrl: ctrl}
	mock.recorder = &MockIProficiencyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIProficiencyRepository) EXPECT() *MockIProficiencyRepositoryMockRecorder {
	return m.recorder
}

// SearchProficiency mocks base method
func (m *MockIProficiencyRepository) SearchProficiency(criteria reference.SearchProficiencyCriteria) ([]reference.Proficiency, error) {
	ret := m.ctrl.Call(m, "SearchProficiency", criteria)
	ret0, _ := ret[0].([]reference.Proficiency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProficiency indicates an expected call of SearchProficiency
func (mr *MockIProficiencyRepositoryMockRecorder) SearchProficiency(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProficiency", reflect.TypeOf((*MockIProficiencyRepository)(nil).SearchProficiency), criteria)
}

// CreateProficiency mocks base method
func (m *MockIProficiencyRepository) CreateProficiency(proficiency *reference.Proficiency) error {
	ret := m.ctrl.Call(m, "CreateProficiency", proficiency)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProficiency indicates an expected call of CreateProficiency
func (mr *MockIProficiencyRepositoryMockRecorder) CreateProficiency(proficiency interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProficiency", reflect.TypeOf((*MockIProficiencyRepository)(nil).CreateProficiency), proficiency)
}

// UpdateProficiency mocks base method
func (m *MockIProficiencyRepository) UpdateProficiency(proficiency reference.Proficiency) error {
	ret := m.ctrl.Call(m, "UpdateProficiency", proficiency)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProficiency indicates an expected call of UpdateProficiency
func (mr *MockIProficiencyRepositoryMockRecorder) UpdateProficiency(proficiency interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProficiency", reflect.TypeOf((*MockIProficiencyRepository)(nil).UpdateProficiency), proficiency)
}

// DeleteProficiency mocks base method
func (m *MockIProficiencyRepository) DeleteProficiency(proficiency reference.Proficiency) error {
	ret := m.ctrl.Call(m, "DeleteProficiency", proficiency)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProficiency indicates an expected call of DeleteProficiency
func (mr *MockIProficiencyRepositoryMockRecorder) DeleteProficiency(proficiency interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProficiency", reflect.TypeOf((*MockIProficiencyRepository)(nil).DeleteProficiency), proficiency)
}
