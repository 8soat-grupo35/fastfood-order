package controllers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	mockUsecase "github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type ItemControllerSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	useCase    *mockUsecase.MockItemUseCase
	controller *ItemController
}

func (suite *ItemControllerSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.useCase = mockUsecase.NewMockItemUseCase(suite.ctrl)
	suite.controller = &ItemController{UseCase: suite.useCase}
}

func (suite *ItemControllerSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *ItemControllerSuite) TestGetAll() {
	expectedItems := []entities.Item{
		{ID: 1, Name: "Burger", Category: "LANCHE"},
	}

	suite.useCase.EXPECT().GetAll(gomock.Any()).Return(expectedItems, nil)

	items, err := suite.controller.GetAll()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedItems, items)
}

func (suite *ItemControllerSuite) TestGetAllByCategory() {
	expectedItems := []entities.Item{
		{ID: 1, Name: "Burger", Category: "LANCHE"},
	}

	suite.useCase.EXPECT().GetAll(gomock.Any()).Return(expectedItems, nil)

	items, err := suite.controller.GetAllByCategory("LANCHE")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedItems, items)
}

func (suite *ItemControllerSuite) TestCreate() {
	itemDto := dto.ItemDto{Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}
	newItem := &entities.Item{ID: 1, Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}

	suite.useCase.EXPECT().Create(gomock.Any()).Return(newItem, nil)

	createdItem, err := suite.controller.Create(itemDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), newItem, createdItem)
}

func (suite *ItemControllerSuite) TestUpdate() {
	itemDto := dto.ItemDto{Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}
	itemAfterUpdate := &entities.Item{ID: 1, Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}

	suite.useCase.EXPECT().Update(uint32(1), gomock.Any()).Return(itemAfterUpdate, nil)

	updatedItem, err := suite.controller.Update(1, itemDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), itemAfterUpdate, updatedItem)
}

func (suite *ItemControllerSuite) TestDelete() {
	suite.useCase.EXPECT().Delete(uint32(1)).Return(nil)

	err := suite.controller.Delete(1)
	assert.NoError(suite.T(), err)
}

func TestItemControllerSuite(t *testing.T) {
	suite.Run(t, new(ItemControllerSuite))
}
