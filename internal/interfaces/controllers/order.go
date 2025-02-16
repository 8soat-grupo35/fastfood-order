package controllers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/presenters"
)

//go:generate mockgen -source=order.go -destination=mock/order.go
type OrderController interface {
	GetAll() ([]entities.Order, error)
	Checkout(orderDto dto.OrderDto) (*presenters.OrderPresenter, error)
	UpdateStatus(id uint32, status string) (*entities.Order, error)
}
