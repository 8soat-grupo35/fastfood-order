package repository

import "github.com/8soat-grupo35/fastfood-order/internal/entities"

//go:generate mockgen -source=order.go -destination=mock/order.go
type OrderRepository interface {
	GetAll() ([]entities.Order, error)
	GetById(id uint32) (*entities.Order, error)
	Create(order entities.Order) (*entities.Order, error)
	Update(id uint32, order entities.Order) (*entities.Order, error)
}
