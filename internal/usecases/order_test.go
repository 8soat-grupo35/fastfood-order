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

type OrderUseCaseSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	repo    *mockRepository.MockOrderRepository
	useCase usecase.OrderUseCase
}

func (suite *OrderUseCaseSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mockRepository.NewMockOrderRepository(suite.ctrl)
	suite.useCase = NewOrderUseCase(suite.repo)
}

func (suite *OrderUseCaseSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *OrderUseCaseSuite) TestGetAll() {
	expectedOrders := []entities.Order{
		{ID: 1, Status: "Pending"},
	}

	suite.repo.EXPECT().GetAll().Return(expectedOrders, nil)

	orders, err := suite.useCase.GetAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedOrders, orders)
}

func (suite *OrderUseCaseSuite) TestGetAllReturnsErrorOnRepositoryFailure() {
	suite.repo.EXPECT().GetAll().Return(nil, errors.New("query error"))

	orders, err := suite.useCase.GetAll()
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), orders)
	assert.Equal(suite.T(), "get order from repository has failed", err.Error())
}

func (suite *OrderUseCaseSuite) TestCreate() {
	itemsDto := []dto.OrderItemDto{
		{Id: 1, Quantity: 2},
	}
	orderDto := dto.OrderDto{Status: "Pending", CustomerID: 1, Items: itemsDto}
	newOrder := &entities.Order{ID: 1, Status: "Pending"}

	suite.repo.EXPECT().Create(gomock.Any()).Return(newOrder, nil)

	createdOrder, err := suite.useCase.Create(orderDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), newOrder, createdOrder)
}

func (suite *OrderUseCaseSuite) TestCreateReturnsErrorOnInvalidOrder() {
	orderDto := dto.OrderDto{Status: "Pending", CustomerID: 1}

	createdOrder, err := suite.useCase.Create(orderDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), createdOrder)
	assert.Contains(suite.T(), err.Error(), "items: cannot be blank")
}

func (suite *OrderUseCaseSuite) TestCreateReturnsErrorOnRepositoryFailure() {
	itemsDto := []dto.OrderItemDto{
		{Id: 1, Quantity: 2},
	}
	orderDto := dto.OrderDto{Status: "Pending", CustomerID: 1, Items: itemsDto}

	suite.repo.EXPECT().Create(gomock.Any()).Return(nil, errors.New("insert error"))

	createdOrder, err := suite.useCase.Create(orderDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), createdOrder)
	assert.Equal(suite.T(), "create order on repository has failed", err.Error())
}

func (suite *OrderUseCaseSuite) TestUpdateStatus() {
	items := []entities.OrderItem{
		{ID: 1, Quantity: 2},
	}
	orderToUpdate := &entities.Order{ID: 1, Status: entities.IN_PREPARATION_STATUS, CustomerID: 1, Items: items}
	orderAfterUpdate := &entities.Order{ID: 1, Status: entities.DONE_STATUS, CustomerID: 1, Items: items}

	suite.repo.EXPECT().GetById(uint32(1)).Return(orderToUpdate, nil)
	suite.repo.EXPECT().Update(uint32(1), gomock.Any()).Return(orderAfterUpdate, nil)

	updatedOrder, err := suite.useCase.UpdateStatus(1, entities.DONE_STATUS)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), orderAfterUpdate, updatedOrder)
}

func (suite *OrderUseCaseSuite) TestUpdateStatusReturnsErrorOnOrderNotFound() {
	suite.repo.EXPECT().GetById(uint32(1)).Return(nil, errors.New("order not found"))

	updatedOrder, err := suite.useCase.UpdateStatus(1, entities.DONE_STATUS)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), updatedOrder)
	assert.Equal(suite.T(), "order not found", err.Error())
}

func TestOrderUseCaseSuite(t *testing.T) {
	suite.Run(t, new(OrderUseCaseSuite))
}
