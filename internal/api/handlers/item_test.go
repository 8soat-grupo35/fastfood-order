package handlers

import (
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

type ItemHandlerSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	controller *mockControllers.MockItemController
	handler    *ItemHandler
	e          *echo.Echo
}

func (suite *ItemHandlerSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.controller = mockControllers.NewMockItemController(suite.ctrl)
	suite.handler = &ItemHandler{itemController: suite.controller}
	suite.e = echo.New()
}

func (suite *ItemHandlerSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *ItemHandlerSuite) TestGetAll() {
	expectedItems := []entities.Item{
		{ID: 1, Name: "Burger", Category: "LANCHE"},
	}

	suite.controller.EXPECT().GetAllByCategory(gomock.Any()).Return(expectedItems, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/item", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	err := suite.handler.GetAll(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Equal(suite.T(), `[{"ID":1,"Name":"Burger","Category":"LANCHE","Price":0,"ImageUrl":"","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null}]`+"\n", rec.Body.String())
}

func (suite *ItemHandlerSuite) TestCreate() {
	newItem := &entities.Item{ID: 1, Name: "Burger", Category: "LANCHE", Price: 10, ImageUrl: "http://image.com"}

	suite.controller.EXPECT().Create(gomock.Any()).Return(newItem, nil)

	req := httptest.NewRequest(http.MethodPost, "/v1/item", strings.NewReader(`{"name":"Burger","category":"LANCHE","price":10,"imageUrl":"http://image.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)

	err := suite.handler.Create(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Equal(suite.T(), `{"ID":1,"Name":"Burger","Category":"LANCHE","Price":10,"ImageUrl":"http://image.com","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null}`+"\n", rec.Body.String())
}

func (suite *ItemHandlerSuite) TestUpdate() {
	itemAfterUpdate := &entities.Item{ID: 1, Name: "Burger", Category: "LANCHE", Price: 10, ImageUrl: "http://image.com"}

	suite.controller.EXPECT().Update(1, gomock.Any()).Return(itemAfterUpdate, nil)

	req := httptest.NewRequest(http.MethodPut, "/v1/item/1", strings.NewReader(`{"name":"Burger","category":"LANCHE","price":10,"imageUrl":"http://image.com"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := suite.handler.Update(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Equal(suite.T(), `{"ID":1,"Name":"Burger","Category":"LANCHE","Price":10,"ImageUrl":"http://image.com","CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null}`+"\n", rec.Body.String())
}

func (suite *ItemHandlerSuite) TestDelete() {
	suite.controller.EXPECT().Delete(1).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/v1/item/1", nil)
	rec := httptest.NewRecorder()
	c := suite.e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := suite.handler.Delete(c)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
}

func TestItemHandlerSuite(t *testing.T) {
	suite.Run(t, new(ItemHandlerSuite))
}
