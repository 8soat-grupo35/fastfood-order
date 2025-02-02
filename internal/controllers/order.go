package controllers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/gateways"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"
	"github.com/8soat-grupo35/fastfood-order/internal/presenters"
	"github.com/8soat-grupo35/fastfood-order/internal/usecases"

	"gorm.io/gorm"
)

type OrderController struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderController(db *gorm.DB) *OrderController {
	orderGateway := gateways.NewOrderGateway(db)
	return &OrderController{
		orderUseCase: usecases.NewOrderUseCase(orderGateway),
	}
}

func (o *OrderController) GetAll() ([]entities.Order, error) {

	return o.orderUseCase.GetAll()
}

func (o *OrderController) Checkout(orderDto dto.OrderDto) (*presenters.OrderPresenter, error) {
	order, err := o.orderUseCase.Create(orderDto)

	if err != nil {
		return nil, err
	}

	//TODO: Call payment use case

	return &presenters.OrderPresenter{Id: order.ID}, nil
}

func (o *OrderController) UpdateStatus(id uint32, status string) (*entities.Order, error) {
	order, err := o.orderUseCase.UpdateStatus(id, status)

	if err != nil {
		return nil, err
	}

	return order, nil
}
