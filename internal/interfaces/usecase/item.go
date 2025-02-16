package usecase

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
)

//go:generate mockgen -source=item.go -destination=mock/item.go
type ItemUseCase interface {
	GetAll(category string) ([]entities.Item, error)
	Create(item dto.ItemDto) (*entities.Item, error)
	Update(itemId uint32, item dto.ItemDto) (*entities.Item, error)
	Delete(itemId uint32) error
}
