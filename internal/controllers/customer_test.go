package controllers

import (
	"errors"
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	mockUsecase "github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type CustomerControllerSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	useCase    *mockUsecase.MockCustomerUseCase
	controller *CustomerController
}

func (suite *CustomerControllerSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.useCase = mockUsecase.NewMockCustomerUseCase(suite.ctrl)
	suite.controller = &CustomerController{UseCase: suite.useCase}
}

func (suite *CustomerControllerSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *CustomerControllerSuite) TestGetAll() {
	expectedCustomers := []entities.Customer{
		{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"},
	}

	suite.useCase.EXPECT().GetAll().Return(expectedCustomers, nil)

	customers, err := suite.controller.GetAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCustomers, customers)
}

func (suite *CustomerControllerSuite) TestGetAllReturnsErrorOnFailure() {
	suite.useCase.EXPECT().GetAll().Return(nil, errors.New("query error"))

	customers, err := suite.controller.GetAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), customers)
	assert.Equal(suite.T(), "query error", err.Error())
}

func (suite *CustomerControllerSuite) TestCreate() {
	customerDto := dto.CustomerDto{Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}
	newCustomer := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.useCase.EXPECT().Create(gomock.Any()).Return(newCustomer, nil)

	createdCustomer, err := suite.controller.Create(customerDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), newCustomer, createdCustomer)
}

func (suite *CustomerControllerSuite) TestCreateReturnsErrorOnFailure() {
	customerDto := dto.CustomerDto{Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.useCase.EXPECT().Create(gomock.Any()).Return(nil, errors.New("insert error"))

	createdCustomer, err := suite.controller.Create(customerDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), createdCustomer)
	assert.Equal(suite.T(), "insert error", err.Error())
}

func (suite *CustomerControllerSuite) TestGetByCpf() {
	expectedCustomer := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.useCase.EXPECT().GetByCpf(gomock.Any()).Return(expectedCustomer, nil)

	customer, err := suite.controller.GetByCpf("12345678901")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCustomer, customer)
}

func (suite *CustomerControllerSuite) TestGetByCpfReturnsErrorOnFailure() {
	suite.useCase.EXPECT().GetByCpf(gomock.Any()).Return(nil, errors.New("query error"))

	customer, err := suite.controller.GetByCpf("12345678901")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), customer)
	assert.Equal(suite.T(), "query error", err.Error())
}

func (suite *CustomerControllerSuite) TestUpdate() {
	customerDto := dto.CustomerDto{Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}
	customerToUpdate := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.useCase.EXPECT().Update(uint32(1), gomock.Any()).Return(customerToUpdate, nil)

	updatedCustomer, err := suite.controller.Update(1, customerDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), customerToUpdate, updatedCustomer)
}

func (suite *CustomerControllerSuite) TestUpdateReturnsErrorOnFailure() {
	customerDto := dto.CustomerDto{Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.useCase.EXPECT().Update(uint32(1), gomock.Any()).Return(nil, errors.New("update error"))

	updatedCustomer, err := suite.controller.Update(1, customerDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), updatedCustomer)
	assert.Equal(suite.T(), "update error", err.Error())
}

func (suite *CustomerControllerSuite) TestDelete() {
	suite.useCase.EXPECT().Delete(uint32(1)).Return(nil)

	err := suite.controller.Delete(1)
	assert.NoError(suite.T(), err)
}

func (suite *CustomerControllerSuite) TestDeleteReturnsErrorOnFailure() {
	suite.useCase.EXPECT().Delete(uint32(1)).Return(errors.New("delete error"))

	err := suite.controller.Delete(1)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "delete error", err.Error())
}

func TestCustomerControllerSuite(t *testing.T) {
	suite.Run(t, new(CustomerControllerSuite))
}
