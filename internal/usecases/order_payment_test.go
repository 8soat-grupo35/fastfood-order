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

type OrderPaymentUseCaseSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	repo    *mockRepository.MockOrderPaymentRepository
	useCase usecase.OrderPaymentUseCase
}

func (suite *OrderPaymentUseCaseSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mockRepository.NewMockOrderPaymentRepository(suite.ctrl)
	suite.useCase = NewOrderPaymentUseCase(suite.repo)
}

func (suite *OrderPaymentUseCaseSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *OrderPaymentUseCaseSuite) TestCreate() {
	order := entities.Order{ID: 1}
	orderPayment := dto.OrderPaymentDto{
		OrderID: int(order.ID),
	}
	suite.repo.EXPECT().Create(orderPayment).Return(nil)

	err := suite.useCase.Create(order)
	assert.NoError(suite.T(), err)
}

func (suite *OrderPaymentUseCaseSuite) TestCreateReturnsErrorOnClientFailure() {
	order := entities.Order{ID: 1}
	orderPayment := dto.OrderPaymentDto{
		OrderID: int(order.ID),
	}
	suite.repo.EXPECT().Create(orderPayment).Return(errors.New("client error"))

	err := suite.useCase.Create(order)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "client error", err.Error())
}

func TestOrderPaymentUseCaseSuite(t *testing.T) {
	suite.Run(t, new(OrderPaymentUseCaseSuite))
}
