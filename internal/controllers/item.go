package controllers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/gateways"
	controllersInterface "github.com/8soat-grupo35/fastfood-order/internal/interfaces/controllers"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"
	"github.com/8soat-grupo35/fastfood-order/internal/usecases"
	"gorm.io/gorm"
)

type ItemController struct {
	UseCase usecase.ItemUseCase
}

func NewItemController(db *gorm.DB) controllersInterface.ItemController {
	gateway := gateways.NewItemGateway(db)
	return &ItemController{
		UseCase: usecases.NewItemUseCase(gateway),
	}
}

func (i *ItemController) GetAll() ([]entities.Item, error) {
	return i.UseCase.GetAll("")
}

func (i *ItemController) GetAllByCategory(category string) ([]entities.Item, error) {
	return i.UseCase.GetAll(category)
}

func (i *ItemController) Create(itemDto dto.ItemDto) (*entities.Item, error) {
	return i.UseCase.Create(itemDto)
}

func (i *ItemController) Update(itemId int, itemDto dto.ItemDto) (*entities.Item, error) {
	return i.UseCase.Update(uint32(itemId), itemDto)
}

func (i *ItemController) Delete(itemId int) error {
	return i.UseCase.Delete(uint32(itemId))
}
