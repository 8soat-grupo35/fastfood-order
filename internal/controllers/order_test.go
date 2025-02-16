package controllers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	mockUsecase "github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase/mock"
	"github.com/8soat-grupo35/fastfood-order/internal/presenters"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type OrderControllerSuite struct {
	suite.Suite
	ctrl                *gomock.Controller
	useCase             *mockUsecase.MockOrderUseCase
	orderPaymentUseCase *mockUsecase.MockOrderPaymentUseCase
	controller          *OrderController
}

func (suite *OrderControllerSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.useCase = mockUsecase.NewMockOrderUseCase(suite.ctrl)
	suite.orderPaymentUseCase = mockUsecase.NewMockOrderPaymentUseCase(suite.ctrl)
	suite.controller = &OrderController{UseCase: suite.useCase, OrderPaymentUseCase: suite.orderPaymentUseCase}
}

func (suite *OrderControllerSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *OrderControllerSuite) TestGetAll() {
	expectedOrders := []entities.Order{
		{ID: 1, Status: "Pending"},
	}

	suite.useCase.EXPECT().GetAll().Return(expectedOrders, nil)

	orders, err := suite.controller.GetAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedOrders, orders)
}

func (suite *OrderControllerSuite) TestCheckout() {
	itemsDto := []dto.OrderItemDto{
		{Id: 1, Quantity: 2},
	}
	orderDto := dto.OrderDto{Status: "Pending", CustomerID: 1, Items: itemsDto}
	newOrder := &entities.Order{ID: 1, Status: "Pending"}
	orderCreated := &presenters.OrderPresenter{
		Id: 1,
	}

	suite.useCase.EXPECT().Create(gomock.Any()).Return(newOrder, nil)
	suite.orderPaymentUseCase.EXPECT().Create(*newOrder).Return(nil)

	createdOrder, err := suite.controller.Checkout(orderDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), orderCreated, createdOrder)
}

func (suite *OrderControllerSuite) TestUpdateStatus() {
	items := []entities.OrderItem{
		{ID: 1, Quantity: 2},
	}

	orderAfterUpdate := &entities.Order{ID: 1, Status: entities.DONE_STATUS, CustomerID: 1, Items: items}

	suite.useCase.EXPECT().UpdateStatus(uint32(1), gomock.Any()).Return(orderAfterUpdate, nil)

	updatedOrder, err := suite.controller.UpdateStatus(1, entities.DONE_STATUS)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), orderAfterUpdate, updatedOrder)
}

func TestOrderControllerSuite(t *testing.T) {
	suite.Run(t, new(OrderControllerSuite))
}
