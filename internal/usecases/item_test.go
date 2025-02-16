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

type ItemUseCaseSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	repo    *mockRepository.MockItemRepository
	useCase usecase.ItemUseCase
}

func (suite *ItemUseCaseSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mockRepository.NewMockItemRepository(suite.ctrl)
	suite.useCase = NewItemUseCase(suite.repo)
}

func (suite *ItemUseCaseSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *ItemUseCaseSuite) TestGetAll() {
	expectedItems := []entities.Item{
		{ID: 1, Name: "Burger", Category: "LANCHE"},
	}

	suite.repo.EXPECT().GetAll(gomock.Any()).Return(expectedItems, nil)

	items, err := suite.useCase.GetAll("LANCHE")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedItems, items)
}

func (suite *ItemUseCaseSuite) TestGetAllReturnsErrorOnInvalidCategory() {
	items, err := suite.useCase.GetAll("INVALID_CATEGORY")
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), items)
	assert.Contains(suite.T(), err.Error(), "Category: must be a valid value between")
}

func (suite *ItemUseCaseSuite) TestGetAllReturnsErrorOnRepositoryFailure() {
	suite.repo.EXPECT().GetAll(gomock.Any()).Return(nil, errors.New("query error"))

	items, err := suite.useCase.GetAll("LANCHE")
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), items)
	assert.Equal(suite.T(), "get item from repository has failed", err.Error())
}

func (suite *ItemUseCaseSuite) TestCreate() {
	itemDto := dto.ItemDto{Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}
	newItem := &entities.Item{ID: 1, Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}

	suite.repo.EXPECT().Create(gomock.Any()).Return(newItem, nil)

	createdItem, err := suite.useCase.Create(itemDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), newItem, createdItem)
}

func (suite *ItemUseCaseSuite) TestCreateReturnsErrorOnInvalidItem() {
	itemDto := dto.ItemDto{Name: "", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}

	createdItem, err := suite.useCase.Create(itemDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), createdItem)
	assert.Contains(suite.T(), err.Error(), "Name: cannot be blank")
}

func (suite *ItemUseCaseSuite) TestCreateReturnsErrorOnRepositoryFailure() {
	itemDto := dto.ItemDto{Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}

	suite.repo.EXPECT().Create(gomock.Any()).Return(nil, errors.New("insert error"))

	createdItem, err := suite.useCase.Create(itemDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), createdItem)
	assert.Equal(suite.T(), "create item on repository has failed", err.Error())
}

func (suite *ItemUseCaseSuite) TestUpdate() {
	itemDto := dto.ItemDto{Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}
	itemToUpdate := &entities.Item{ID: 1, Name: "Burger", Category: "SOBREMESA", Price: 5.0, ImageUrl: "http://image.com"}
	itemAfterUpdate := &entities.Item{ID: 1, Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(itemToUpdate, nil)
	suite.repo.EXPECT().Update(uint32(1), gomock.Any()).Return(itemAfterUpdate, nil)

	updatedItem, err := suite.useCase.Update(1, itemDto)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), itemAfterUpdate, updatedItem)
}

func (suite *ItemUseCaseSuite) TestUpdateReturnsErrorOnInvalidItem() {
	itemDto := dto.ItemDto{Name: "", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}

	updatedItem, err := suite.useCase.Update(1, itemDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), updatedItem)
	assert.Contains(suite.T(), err.Error(), "Name: cannot be blank")
}

func (suite *ItemUseCaseSuite) TestUpdateReturnsErrorOnItemNotFound() {
	itemDto := dto.ItemDto{Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(nil, nil)

	updatedItem, err := suite.useCase.Update(1, itemDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), updatedItem)
	assert.Equal(suite.T(), "item not found to update", err.Error())
}

func (suite *ItemUseCaseSuite) TestUpdateReturnsErrorOnRepositoryFailure() {
	itemDto := dto.ItemDto{Name: "Burger", Category: "LANCHE", Price: 10.0, ImageUrl: "http://image.com"}
	itemToUpdate := &entities.Item{ID: 1, Name: "Burger", Category: "SOBREMESA", Price: 5.0, ImageUrl: "http://image.com"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(itemToUpdate, nil)
	suite.repo.EXPECT().Update(uint32(1), gomock.Any()).Return(nil, errors.New("update error"))

	updatedItem, err := suite.useCase.Update(1, itemDto)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), updatedItem)
	assert.Equal(suite.T(), "updated item on repository has failed", err.Error())
}

func (suite *ItemUseCaseSuite) TestDelete() {
	itemToDelete := &entities.Item{ID: 1, Name: "Burger", Category: "Food"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(itemToDelete, nil)
	suite.repo.EXPECT().Delete(uint32(1)).Return(nil)

	err := suite.useCase.Delete(1)
	assert.NoError(suite.T(), err)
}

func (suite *ItemUseCaseSuite) TestDeleteReturnsErrorOnItemNotFound() {
	suite.repo.EXPECT().GetOne(gomock.Any()).Return(nil, nil)

	err := suite.useCase.Delete(1)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "item not found to delete", err.Error())
}

func (suite *ItemUseCaseSuite) TestDeleteReturnsErrorOnRepositoryFailure() {
	itemToDelete := &entities.Item{ID: 1, Name: "Burger", Category: "Food"}

	suite.repo.EXPECT().GetOne(gomock.Any()).Return(itemToDelete, nil)
	suite.repo.EXPECT().Delete(uint32(1)).Return(errors.New("delete error"))

	err := suite.useCase.Delete(1)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "error on delete in repository", err.Error())
}

func TestItemUseCaseSuite(t *testing.T) {
	suite.Run(t, new(ItemUseCaseSuite))
}
