// Code generated by MockGen. DO NOT EDIT.
// Source: controllers/interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/mahendrafathan/go-cake/models"
)

// MockCakeRepository is a mock of CakeRepository interface.
type MockCakeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCakeRepositoryMockRecorder
}

// MockCakeRepositoryMockRecorder is the mock recorder for MockCakeRepository.
type MockCakeRepositoryMockRecorder struct {
	mock *MockCakeRepository
}

// NewMockCakeRepository creates a new mock instance.
func NewMockCakeRepository(ctrl *gomock.Controller) *MockCakeRepository {
	mock := &MockCakeRepository{ctrl: ctrl}
	mock.recorder = &MockCakeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCakeRepository) EXPECT() *MockCakeRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCakeRepository) Create(arg0 models.Cake) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCakeRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCakeRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockCakeRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCakeRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCakeRepository)(nil).Delete), id)
}

// FindAll mocks base method.
func (m *MockCakeRepository) FindAll() ([]models.Cake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]models.Cake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCakeRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCakeRepository)(nil).FindAll))
}

// FindOne mocks base method.
func (m *MockCakeRepository) FindOne(id int) (models.Cake, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", id)
	ret0, _ := ret[0].(models.Cake)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockCakeRepositoryMockRecorder) FindOne(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockCakeRepository)(nil).FindOne), id)
}

// Update mocks base method.
func (m *MockCakeRepository) Update(id int, cake models.Cake) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, cake)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCakeRepositoryMockRecorder) Update(id, cake interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCakeRepository)(nil).Update), id, cake)
}
