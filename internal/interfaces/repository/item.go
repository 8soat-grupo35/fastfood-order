package repository

import "github.com/8soat-grupo35/fastfood-order/internal/entities"

//go:generate mockgen -source=item.go -destination=mock/item.go
type ItemRepository interface {
	GetAll(entities.Item) ([]entities.Item, error)
	GetOne(entities.Item) (*entities.Item, error)
	Create(item entities.Item) (*entities.Item, error)
	Update(itemId uint32, item entities.Item) (*entities.Item, error)
	Delete(itemId uint32) error
}
