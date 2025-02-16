package handlers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	mockControllers "github.com/8soat-grupo35/fastfood-order/internal/interfaces/controllers/mock"
	"github.com/8soat-grupo35/fastfood-order/internal/presenters"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type OrderHandlerSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	controller *mockControllers.MockOrderController
	handler    *OrderHandler
	e          *echo.Echo
}

func (suite *OrderHandlerSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.controller = mockControllers.NewMockOrderController(suite.ctrl)
	suite.handler = &OrderHandler{orderController: suite.controller}
	suite.e = echo.New()
}

func (suite *OrderHandlerSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *OrderHandlerSuite) TestGetAll() {
	expectedOrders := []entities.Order{
		{ID: 1, Status: "Pending"},
	}

	suite.controller.EXPECT().GetAll().Return(expectedOrders, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/orders", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	err := suite.handler.GetAll(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func (suite *OrderHandlerSuite) TestCheckout() {
	orderPresenter := &presenters.OrderPresenter{Id: 1}

	suite.controller.EXPECT().Checkout(gomock.Any()).Return(orderPresenter, nil)

	req := httptest.NewRequest(http.MethodPost, "/v1/orders/checkout", strings.NewReader(`{"status":"Pending","customerID":1,"items":[{"id":1,"quantity":2}]}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	err := suite.handler.Checkout(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Equal(suite.T(), `{"id":1}`+"\n", rec.Body.String())
}

func (suite *OrderHandlerSuite) TestUpdateStatus() {
	items := []entities.OrderItem{
		{ID: 1, Quantity: 2},
	}
	orderAfterUpdate := &entities.Order{ID: 1, Status: entities.DONE_STATUS, CustomerID: 1, Items: items}

	suite.controller.EXPECT().UpdateStatus(uint32(1), gomock.Any()).Return(orderAfterUpdate, nil)

	req := httptest.NewRequest(http.MethodPatch, "/v1/orders/1", strings.NewReader(`{"status":"DONE"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := suite.handler.UpdateStatus(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func TestOrderHandlerSuite(t *testing.T) {
	suite.Run(t, new(OrderHandlerSuite))
}
