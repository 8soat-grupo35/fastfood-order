package usecase

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
)

type OrderUseCase interface {
	GetAll() ([]entities.Order, error)
	Create(order dto.OrderDto) (*entities.Order, error)
	UpdateStatus(id uint32, status string) (*entities.Order, error)
}
