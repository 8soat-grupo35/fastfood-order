package controllers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/gateways"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"
	"github.com/8soat-grupo35/fastfood-order/internal/usecases"
	"gorm.io/gorm"
)

type ItemController struct {
	dbConnection *gorm.DB
	gateway      repository.ItemRepository
	useCase      usecase.ItemUseCase
}

func NewItemController(db *gorm.DB) *ItemController {
	gateway := gateways.NewItemGateway(db)
	return &ItemController{
		dbConnection: db,
		gateway:      gateway,
		useCase:      usecases.NewItemUseCase(gateway),
	}
}

func (i *ItemController) GetAll() ([]entities.Item, error) {
	return i.useCase.GetAll("")
}

func (i *ItemController) GetAllByCategory(category string) ([]entities.Item, error) {
	return i.useCase.GetAll(category)
}

func (i *ItemController) Create(itemDto dto.ItemDto) (*entities.Item, error) {
	return i.useCase.Create(itemDto)
}

func (i *ItemController) Update(itemId int, itemDto dto.ItemDto) (*entities.Item, error) {
	return i.useCase.Update(uint32(itemId), itemDto)
}

func (i *ItemController) Delete(itemId int) error {
	return i.useCase.Delete(uint32(itemId))
}
