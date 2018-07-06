// Code generated by MockGen. DO NOT EDIT.
// Source: ./businesslogic/competitionentry.go

// Package mock_businesslogic is a generated GoMock package.
package mock_businesslogic

import (
	businesslogic "github.com/DancesportSoftware/das/businesslogic"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIAthleteCompetitionEntryRepository is a mock of IAthleteCompetitionEntryRepository interface
type MockIAthleteCompetitionEntryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAthleteCompetitionEntryRepositoryMockRecorder
}

// MockIAthleteCompetitionEntryRepositoryMockRecorder is the mock recorder for MockIAthleteCompetitionEntryRepository
type MockIAthleteCompetitionEntryRepositoryMockRecorder struct {
	mock *MockIAthleteCompetitionEntryRepository
}

// NewMockIAthleteCompetitionEntryRepository creates a new mock instance
func NewMockIAthleteCompetitionEntryRepository(ctrl *gomock.Controller) *MockIAthleteCompetitionEntryRepository {
	mock := &MockIAthleteCompetitionEntryRepository{ctrl: ctrl}
	mock.recorder = &MockIAthleteCompetitionEntryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAthleteCompetitionEntryRepository) EXPECT() *MockIAthleteCompetitionEntryRepositoryMockRecorder {
	return m.recorder
}

// CreateAthleteCompetitionEntry mocks base method
func (m *MockIAthleteCompetitionEntryRepository) CreateAthleteCompetitionEntry(entry *businesslogic.AthleteCompetitionEntry) error {
	ret := m.ctrl.Call(m, "CreateAthleteCompetitionEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAthleteCompetitionEntry indicates an expected call of CreateAthleteCompetitionEntry
func (mr *MockIAthleteCompetitionEntryRepositoryMockRecorder) CreateAthleteCompetitionEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAthleteCompetitionEntry", reflect.TypeOf((*MockIAthleteCompetitionEntryRepository)(nil).CreateAthleteCompetitionEntry), entry)
}

// DeleteAthleteCompetitionEntry mocks base method
func (m *MockIAthleteCompetitionEntryRepository) DeleteAthleteCompetitionEntry(entry businesslogic.AthleteCompetitionEntry) error {
	ret := m.ctrl.Call(m, "DeleteAthleteCompetitionEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAthleteCompetitionEntry indicates an expected call of DeleteAthleteCompetitionEntry
func (mr *MockIAthleteCompetitionEntryRepositoryMockRecorder) DeleteAthleteCompetitionEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAthleteCompetitionEntry", reflect.TypeOf((*MockIAthleteCompetitionEntryRepository)(nil).DeleteAthleteCompetitionEntry), entry)
}

// SearchAthleteCompetitionEntry mocks base method
func (m *MockIAthleteCompetitionEntryRepository) SearchAthleteCompetitionEntry(criteria businesslogic.SearchAthleteCompetitionEntryCriteria) ([]businesslogic.AthleteCompetitionEntry, error) {
	ret := m.ctrl.Call(m, "SearchAthleteCompetitionEntry", criteria)
	ret0, _ := ret[0].([]businesslogic.AthleteCompetitionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAthleteCompetitionEntry indicates an expected call of SearchAthleteCompetitionEntry
func (mr *MockIAthleteCompetitionEntryRepositoryMockRecorder) SearchAthleteCompetitionEntry(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAthleteCompetitionEntry", reflect.TypeOf((*MockIAthleteCompetitionEntryRepository)(nil).SearchAthleteCompetitionEntry), criteria)
}

// UpdateAthleteCompetitionEntry mocks base method
func (m *MockIAthleteCompetitionEntryRepository) UpdateAthleteCompetitionEntry(entry businesslogic.AthleteCompetitionEntry) error {
	ret := m.ctrl.Call(m, "UpdateAthleteCompetitionEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAthleteCompetitionEntry indicates an expected call of UpdateAthleteCompetitionEntry
func (mr *MockIAthleteCompetitionEntryRepositoryMockRecorder) UpdateAthleteCompetitionEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAthleteCompetitionEntry", reflect.TypeOf((*MockIAthleteCompetitionEntryRepository)(nil).UpdateAthleteCompetitionEntry), entry)
}

// MockIPartnershipCompetitionEntryRepository is a mock of IPartnershipCompetitionEntryRepository interface
type MockIPartnershipCompetitionEntryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPartnershipCompetitionEntryRepositoryMockRecorder
}

// MockIPartnershipCompetitionEntryRepositoryMockRecorder is the mock recorder for MockIPartnershipCompetitionEntryRepository
type MockIPartnershipCompetitionEntryRepositoryMockRecorder struct {
	mock *MockIPartnershipCompetitionEntryRepository
}

// NewMockIPartnershipCompetitionEntryRepository creates a new mock instance
func NewMockIPartnershipCompetitionEntryRepository(ctrl *gomock.Controller) *MockIPartnershipCompetitionEntryRepository {
	mock := &MockIPartnershipCompetitionEntryRepository{ctrl: ctrl}
	mock.recorder = &MockIPartnershipCompetitionEntryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPartnershipCompetitionEntryRepository) EXPECT() *MockIPartnershipCompetitionEntryRepositoryMockRecorder {
	return m.recorder
}

// CreatePartnershipCompetitionEntry mocks base method
func (m *MockIPartnershipCompetitionEntryRepository) CreatePartnershipCompetitionEntry(entry *businesslogic.PartnershipCompetitionEntry) error {
	ret := m.ctrl.Call(m, "CreatePartnershipCompetitionEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePartnershipCompetitionEntry indicates an expected call of CreatePartnershipCompetitionEntry
func (mr *MockIPartnershipCompetitionEntryRepositoryMockRecorder) CreatePartnershipCompetitionEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePartnershipCompetitionEntry", reflect.TypeOf((*MockIPartnershipCompetitionEntryRepository)(nil).CreatePartnershipCompetitionEntry), entry)
}

// DeletePartnershipCompetitionEntry mocks base method
func (m *MockIPartnershipCompetitionEntryRepository) DeletePartnershipCompetitionEntry(entry businesslogic.PartnershipCompetitionEntry) error {
	ret := m.ctrl.Call(m, "DeletePartnershipCompetitionEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePartnershipCompetitionEntry indicates an expected call of DeletePartnershipCompetitionEntry
func (mr *MockIPartnershipCompetitionEntryRepositoryMockRecorder) DeletePartnershipCompetitionEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePartnershipCompetitionEntry", reflect.TypeOf((*MockIPartnershipCompetitionEntryRepository)(nil).DeletePartnershipCompetitionEntry), entry)
}

// SearchPartnershipCompetitionEntry mocks base method
func (m *MockIPartnershipCompetitionEntryRepository) SearchPartnershipCompetitionEntry(criteria businesslogic.SearchPartnershipCompetitionEntryCriteria) ([]businesslogic.PartnershipCompetitionEntry, error) {
	ret := m.ctrl.Call(m, "SearchPartnershipCompetitionEntry", criteria)
	ret0, _ := ret[0].([]businesslogic.PartnershipCompetitionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPartnershipCompetitionEntry indicates an expected call of SearchPartnershipCompetitionEntry
func (mr *MockIPartnershipCompetitionEntryRepositoryMockRecorder) SearchPartnershipCompetitionEntry(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPartnershipCompetitionEntry", reflect.TypeOf((*MockIPartnershipCompetitionEntryRepository)(nil).SearchPartnershipCompetitionEntry), criteria)
}

// UpdatePartnershipCompetitionEntry mocks base method
func (m *MockIPartnershipCompetitionEntryRepository) UpdatePartnershipCompetitionEntry(entry businesslogic.PartnershipCompetitionEntry) error {
	ret := m.ctrl.Call(m, "UpdatePartnershipCompetitionEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePartnershipCompetitionEntry indicates an expected call of UpdatePartnershipCompetitionEntry
func (mr *MockIPartnershipCompetitionEntryRepositoryMockRecorder) UpdatePartnershipCompetitionEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePartnershipCompetitionEntry", reflect.TypeOf((*MockIPartnershipCompetitionEntryRepository)(nil).UpdatePartnershipCompetitionEntry), entry)
}

// MockIAdjudicatorCompetitionEntryRepository is a mock of IAdjudicatorCompetitionEntryRepository interface
type MockIAdjudicatorCompetitionEntryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAdjudicatorCompetitionEntryRepositoryMockRecorder
}

// MockIAdjudicatorCompetitionEntryRepositoryMockRecorder is the mock recorder for MockIAdjudicatorCompetitionEntryRepository
type MockIAdjudicatorCompetitionEntryRepositoryMockRecorder struct {
	mock *MockIAdjudicatorCompetitionEntryRepository
}

// NewMockIAdjudicatorCompetitionEntryRepository creates a new mock instance
func NewMockIAdjudicatorCompetitionEntryRepository(ctrl *gomock.Controller) *MockIAdjudicatorCompetitionEntryRepository {
	mock := &MockIAdjudicatorCompetitionEntryRepository{ctrl: ctrl}
	mock.recorder = &MockIAdjudicatorCompetitionEntryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIAdjudicatorCompetitionEntryRepository) EXPECT() *MockIAdjudicatorCompetitionEntryRepositoryMockRecorder {
	return m.recorder
}

// CreateAdjudicatorCompetitionEntry mocks base method
func (m *MockIAdjudicatorCompetitionEntryRepository) CreateAdjudicatorCompetitionEntry(entry *businesslogic.AdjudicatorCompetitionEntry) error {
	ret := m.ctrl.Call(m, "CreateAdjudicatorCompetitionEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAdjudicatorCompetitionEntry indicates an expected call of CreateAdjudicatorCompetitionEntry
func (mr *MockIAdjudicatorCompetitionEntryRepositoryMockRecorder) CreateAdjudicatorCompetitionEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdjudicatorCompetitionEntry", reflect.TypeOf((*MockIAdjudicatorCompetitionEntryRepository)(nil).CreateAdjudicatorCompetitionEntry), entry)
}

// DeleteAdjudicatorCompetitionEntry mocks base method
func (m *MockIAdjudicatorCompetitionEntryRepository) DeleteAdjudicatorCompetitionEntry(entry businesslogic.AdjudicatorCompetitionEntry) error {
	ret := m.ctrl.Call(m, "DeleteAdjudicatorCompetitionEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAdjudicatorCompetitionEntry indicates an expected call of DeleteAdjudicatorCompetitionEntry
func (mr *MockIAdjudicatorCompetitionEntryRepositoryMockRecorder) DeleteAdjudicatorCompetitionEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdjudicatorCompetitionEntry", reflect.TypeOf((*MockIAdjudicatorCompetitionEntryRepository)(nil).DeleteAdjudicatorCompetitionEntry), entry)
}

// SearchAdjudicatorCompetitionEntry mocks base method
func (m *MockIAdjudicatorCompetitionEntryRepository) SearchAdjudicatorCompetitionEntry(criteria businesslogic.SearchAdjudicatorCompetitionEntryCriteria) ([]businesslogic.AdjudicatorCompetitionEntry, error) {
	ret := m.ctrl.Call(m, "SearchAdjudicatorCompetitionEntry", criteria)
	ret0, _ := ret[0].([]businesslogic.AdjudicatorCompetitionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchAdjudicatorCompetitionEntry indicates an expected call of SearchAdjudicatorCompetitionEntry
func (mr *MockIAdjudicatorCompetitionEntryRepositoryMockRecorder) SearchAdjudicatorCompetitionEntry(criteria interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchAdjudicatorCompetitionEntry", reflect.TypeOf((*MockIAdjudicatorCompetitionEntryRepository)(nil).SearchAdjudicatorCompetitionEntry), criteria)
}

// UpdateAdjudicatorCompetitionEntry mocks base method
func (m *MockIAdjudicatorCompetitionEntryRepository) UpdateAdjudicatorCompetitionEntry(entry businesslogic.AdjudicatorCompetitionEntry) error {
	ret := m.ctrl.Call(m, "UpdateAdjudicatorCompetitionEntry", entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdjudicatorCompetitionEntry indicates an expected call of UpdateAdjudicatorCompetitionEntry
func (mr *MockIAdjudicatorCompetitionEntryRepositoryMockRecorder) UpdateAdjudicatorCompetitionEntry(entry interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdjudicatorCompetitionEntry", reflect.TypeOf((*MockIAdjudicatorCompetitionEntryRepository)(nil).UpdateAdjudicatorCompetitionEntry), entry)
}
