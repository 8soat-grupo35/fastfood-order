package usecases

import (
	"errors"
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	mockRepository "github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository/mock"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type CustomerUseCaseSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	repo    *mockRepository.MockCustomerRepository
	useCase usecase.CustomerUseCase
}

func (suite *CustomerUseCaseSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mockRepository.NewMockCustomerRepository(suite.ctrl)
	suite.useCase = NewCustomerUseCase(suite.repo)
}

func (suite *CustomerUseCaseSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *CustomerUseCaseSuite) TestGetAll() {
	expectedCustomers := []entities.Customer{
		{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"},
	}

	suite.repo.EXPECT().GetAll().Return(expectedCustomers, nil)

	customers, err := suite.useCase.GetAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCustomers, customers)
}

func (suite *CustomerUseCaseSuite) TestGetAllReturnsErrorOnFailure() {
	suite.repo.EXPECT().GetAll().Return(nil, errors.New("query error"))

	customers, err := suite.useCase.GetAll()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), customers)
	assert.Equal(suite.T(), "get customer from repository has failed", err.Error())
}

func (suite *CustomerUseCaseSuite) TestCreate() {
	customerDto := dto.CustomerDto{Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}
	newCustomer := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.repo.EXPECT().Create(gomock.Any()).Return(newCustomer, nil)

	createdCustomer, err := suite.useCase.Create(customerDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), newCustomer, createdCustomer)
}

func (suite *CustomerUseCaseSuite) TestCreateReturnsErrorOnInvalidCustomer() {
	customerDto := dto.CustomerDto{Name: "John", CPF: "12345678901", Email: "test@email.com"}

	createdCustomer, err := suite.useCase.Create(customerDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), createdCustomer)
	assert.Contains(suite.T(), err.Error(), "Name: the length must be between")
}

func (suite *CustomerUseCaseSuite) TestCreateReturnsErrorOnRepositoryFailure() {
	customerDto := dto.CustomerDto{Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.repo.EXPECT().Create(gomock.Any()).Return(nil, errors.New("insert error"))

	createdCustomer, err := suite.useCase.Create(customerDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), createdCustomer)
	assert.Equal(suite.T(), "create customer on repository has failed", err.Error())
}

func (suite *CustomerUseCaseSuite) TestGetByCpf() {
	expectedCustomer := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(expectedCustomer, nil)

	customer, err := suite.useCase.GetByCpf("12345678901")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedCustomer, customer)
}

func (suite *CustomerUseCaseSuite) TestGetByCpfReturnsErrorOnFailure() {
	suite.repo.EXPECT().GetOne(gomock.Any()).Return(nil, errors.New("query error"))

	customer, err := suite.useCase.GetByCpf("12345678901")
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), customer)
	assert.Equal(suite.T(), "error on obtain customer by CPF in repository", err.Error())
}

func (suite *CustomerUseCaseSuite) TestUpdate() {
	customerDto := dto.CustomerDto{Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}
	customerToUpdate := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(customerToUpdate, nil)
	suite.repo.EXPECT().Update(uint32(1), gomock.Any()).Return(customerToUpdate, nil)

	updatedCustomer, err := suite.useCase.Update(1, customerDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), customerToUpdate, updatedCustomer)
}

func (suite *CustomerUseCaseSuite) TestUpdateReturnsErrorOnInvalidCustomer() {
	customerDto := dto.CustomerDto{Name: "John", CPF: "12345678901", Email: "test@email.com"}

	updatedCustomer, err := suite.useCase.Update(1, customerDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), updatedCustomer)
	assert.Contains(suite.T(), err.Error(), "Name: the length must be between")
}

func (suite *CustomerUseCaseSuite) TestUpdateReturnsErrorOnCustomerNotFound() {
	customerDto := dto.CustomerDto{Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(nil, nil)

	updatedCustomer, err := suite.useCase.Update(1, customerDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), updatedCustomer)
	assert.Equal(suite.T(), "customer not found to update", err.Error())
}

func (suite *CustomerUseCaseSuite) TestUpdateReturnsErrorOnRepositoryFailure() {
	customerDto := dto.CustomerDto{Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}
	customerToUpdate := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(customerToUpdate, nil)
	suite.repo.EXPECT().Update(uint32(1), gomock.Any()).Return(nil, errors.New("update error"))

	updatedCustomer, err := suite.useCase.Update(1, customerDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), updatedCustomer)
	assert.Equal(suite.T(), "updated customer on repository has failed", err.Error())
}

func (suite *CustomerUseCaseSuite) TestDelete() {
	customerToDelete := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(customerToDelete, nil)
	suite.repo.EXPECT().Delete(uint32(1)).Return(nil)

	err := suite.useCase.Delete(1)
	assert.NoError(suite.T(), err)
}

func (suite *CustomerUseCaseSuite) TestDeleteReturnsErrorOnCustomerNotFound() {
	suite.repo.EXPECT().GetOne(gomock.Any()).Return(nil, nil)

	err := suite.useCase.Delete(1)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "customer not found to delete", err.Error())
}

func (suite *CustomerUseCaseSuite) TestDeleteReturnsErrorOnRepositoryFailure() {
	customerToDelete := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(customerToDelete, nil)
	suite.repo.EXPECT().Delete(uint32(1)).Return(errors.New("delete error"))

	err := suite.useCase.Delete(1)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error on delete in repository", err.Error())
}

func TestCustomerUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CustomerUseCaseSuite))
}
