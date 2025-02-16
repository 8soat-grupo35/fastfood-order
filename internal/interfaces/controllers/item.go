package controllers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
)

//go:generate mockgen -source=item.go -destination=mock/item.go
type ItemController interface {
	GetAllByCategory(category string) ([]entities.Item, error)
	Create(itemDto dto.ItemDto) (*entities.Item, error)
	Update(itemId int, itemDto dto.ItemDto) (*entities.Item, error)
	Delete(itemId int) error
}
