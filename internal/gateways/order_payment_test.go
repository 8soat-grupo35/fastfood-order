package gateways

import (
	"errors"
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	mockHttp "github.com/8soat-grupo35/fastfood-order/internal/interfaces/http/mock"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type OrderPaymentRepositorySuite struct {
	suite.Suite
	ctrl   *gomock.Controller
	client *mockHttp.MockClient
	repo   repository.OrderPaymentRepository
}

func (suite *OrderPaymentRepositorySuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.client = mockHttp.NewMockClient(suite.ctrl)
	suite.repo = NewOrderPaymentGateway(suite.client)
}

func (suite *OrderPaymentRepositorySuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *OrderPaymentRepositorySuite) TestCreate() {
	orderPaymentDto := dto.OrderPaymentDto{OrderID: 1, PaymentStatusID: dto.PAYMENT_STATUS_WAITING}
	suite.client.EXPECT().Post("/order-payment", gomock.Any()).Return(nil, nil)

	err := suite.repo.Create(orderPaymentDto)
	assert.NoError(suite.T(), err)
}

func (suite *OrderPaymentRepositorySuite) TestCreateReturnsErrorOnClientFailure() {
	orderPaymentDto := dto.OrderPaymentDto{OrderID: 1}
	suite.client.EXPECT().Post("/order-payment", gomock.Any()).Return(nil, errors.New("client error"))

	err := suite.repo.Create(orderPaymentDto)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "client error", err.Error())
}

func TestOrderPaymentRepositorySuite(t *testing.T) {
	suite.Run(t, new(OrderPaymentRepositorySuite))
}
