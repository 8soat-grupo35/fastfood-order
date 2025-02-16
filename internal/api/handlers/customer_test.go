package handlers

import (
	"errors"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	mockControllers "github.com/8soat-grupo35/fastfood-order/internal/interfaces/controllers/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type CustomerHandlerSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	controller *mockControllers.MockCustomerController
	handler    *CustomerHandler
	e          *echo.Echo
}

func (suite *CustomerHandlerSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.controller = mockControllers.NewMockCustomerController(suite.ctrl)
	suite.handler = &CustomerHandler{customerController: suite.controller}
	suite.e = echo.New()
}

func (suite *CustomerHandlerSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *CustomerHandlerSuite) TestGetAll() {
	expectedCustomers := []entities.Customer{
		{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"},
	}

	suite.controller.EXPECT().GetAll().Return(expectedCustomers, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/customer", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	err := suite.handler.GetAll(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func (suite *CustomerHandlerSuite) TestGetAllReturnsErrorOnFailure() {
	suite.controller.EXPECT().GetAll().Return(nil, errors.New("query error"))

	req := httptest.NewRequest(http.MethodGet, "/v1/customer", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	suite.handler.GetAll(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "query error")
}

func (suite *CustomerHandlerSuite) TestCreate() {
	newCustomer := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.controller.EXPECT().Create(gomock.Any()).Return(newCustomer, nil)

	req := httptest.NewRequest(http.MethodPost, "/v1/customer", strings.NewReader(`{"name":"John Doe","cpf":"12345678901","email":"test@email.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	err := suite.handler.Create(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func (suite *CustomerHandlerSuite) TestCreateReturnsErrorOnInvalidInput() {
	req := httptest.NewRequest(http.MethodPost, "/v1/customer", strings.NewReader(`{"name":"","cpf":12345678901,"email":"test@email.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	suite.handler.Create(c)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
}

func (suite *CustomerHandlerSuite) TestCreateReturnsErrorOnFailure() {
	suite.controller.EXPECT().Create(gomock.Any()).Return(nil, errors.New("insert error"))

	req := httptest.NewRequest(http.MethodPost, "/v1/customer", strings.NewReader(`{"name":"John Doe","cpf":"12345678901","email":"test@email.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	suite.handler.Create(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "insert error")
}

func (suite *CustomerHandlerSuite) TestGetByCpf() {
	expectedCustomer := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.controller.EXPECT().GetByCpf(gomock.Any()).Return(expectedCustomer, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/customer/cpf/12345678901", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)
	c.SetParamNames("cpf")
	c.SetParamValues("12345678901")

	err := suite.handler.GetByCpf(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func (suite *CustomerHandlerSuite) TestGetByCpfReturnsErrorOnFailure() {
	suite.controller.EXPECT().GetByCpf(gomock.Any()).Return(nil, errors.New("query error"))

	req := httptest.NewRequest(http.MethodGet, "/v1/customer/cpf/12345678901", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)
	c.SetParamNames("cpf")
	c.SetParamValues("12345678901")

	suite.handler.GetByCpf(c)
	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "query error")
}

func (suite *CustomerHandlerSuite) TestUpdate() {
	customerToUpdate := &entities.Customer{ID: 1, Name: "John Doe", CPF: "12345678901", Email: "test@email.com"}

	suite.controller.EXPECT().Update(uint32(1), gomock.Any()).Return(customerToUpdate, nil)

	req := httptest.NewRequest(http.MethodPut, "/v1/customer/1", strings.NewReader(`{"name":"John Doe","cpf":"12345678901","email":"test@email.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := suite.handler.Update(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func (suite *CustomerHandlerSuite) TestUpdateReturnsErrorOnInvalidInput() {
	req := httptest.NewRequest(http.MethodPut, "/v1/customer/1", strings.NewReader(`{"name":"","cpf":12345678901,"email":"test@email.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	suite.handler.Update(c)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
}

func (suite *CustomerHandlerSuite) TestUpdateReturnsErrorOnFailure() {
	suite.controller.EXPECT().Update(uint32(1), gomock.Any()).Return(nil, errors.New("update error"))

	req := httptest.NewRequest(http.MethodPut, "/v1/customer/1", strings.NewReader(`{"name":"John Doe","cpf":"12345678901","email":"test@email.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	suite.handler.Update(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "update error")
}

func (suite *CustomerHandlerSuite) TestDelete() {
	suite.controller.EXPECT().Delete(uint32(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/v1/customer/1", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := suite.handler.Delete(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func (suite *CustomerHandlerSuite) TestDeleteReturnsErrorOnFailure() {
	suite.controller.EXPECT().Delete(uint32(1)).Return(errors.New("delete error"))

	req := httptest.NewRequest(http.MethodDelete, "/v1/customer/1", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	suite.handler.Delete(c)
	assert.Equal(suite.T(), http.StatusNotFound, rec.Code)
	assert.Contains(suite.T(), rec.Body.String(), "delete error")
}

func TestCustomerHandlerSuite(t *testing.T) {
	suite.Run(t, new(CustomerHandlerSuite))
}
