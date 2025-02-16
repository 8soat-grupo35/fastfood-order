// Code generated by MockGen. DO NOT EDIT.
// Source: customer.go
//
// Generated by this command:
//
//	mockgen -source=customer.go -destination=mock/customer.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	entities "github.com/8soat-grupo35/fastfood-order/internal/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockCustomerRepository is a mock of CustomerRepository interface.
type MockCustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepositoryMockRecorder
	isgomock struct{}
}

// MockCustomerRepositoryMockRecorder is the mock recorder for MockCustomerRepository.
type MockCustomerRepositoryMockRecorder struct {
	mock *MockCustomerRepository
}

// NewMockCustomerRepository creates a new mock instance.
func NewMockCustomerRepository(ctrl *gomock.Controller) *MockCustomerRepository {
	mock := &MockCustomerRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerRepository) EXPECT() *MockCustomerRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCustomerRepository) Create(customer entities.Customer) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", customer)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCustomerRepositoryMockRecorder) Create(customer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomerRepository)(nil).Create), customer)
}

// Delete mocks base method.
func (m *MockCustomerRepository) Delete(customerId uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", customerId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomerRepositoryMockRecorder) Delete(customerId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomerRepository)(nil).Delete), customerId)
}

// GetAll mocks base method.
func (m *MockCustomerRepository) GetAll() ([]entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCustomerRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCustomerRepository)(nil).GetAll))
}

// GetOne mocks base method.
func (m *MockCustomerRepository) GetOne(arg0 entities.Customer) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOne", arg0)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne.
func (mr *MockCustomerRepositoryMockRecorder) GetOne(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockCustomerRepository)(nil).GetOne), arg0)
}

// Update mocks base method.
func (m *MockCustomerRepository) Update(customerId uint32, customer entities.Customer) (*entities.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", customerId, customer)
	ret0, _ := ret[0].(*entities.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCustomerRepositoryMockRecorder) Update(customerId, customer any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomerRepository)(nil).Update), customerId, customer)
}
