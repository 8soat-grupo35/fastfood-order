package controllers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/gateways"
	controllersInterface "github.com/8soat-grupo35/fastfood-order/internal/interfaces/controllers"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"
	"github.com/8soat-grupo35/fastfood-order/internal/presenters"
	"github.com/8soat-grupo35/fastfood-order/internal/usecases"

	"gorm.io/gorm"
)

type OrderController struct {
	UseCase usecase.OrderUseCase
}

func NewOrderController(db *gorm.DB) controllersInterface.OrderController {
	orderGateway := gateways.NewOrderGateway(db)
	return &OrderController{
		UseCase: usecases.NewOrderUseCase(orderGateway),
	}
}

func (o *OrderController) GetAll() ([]entities.Order, error) {

	return o.UseCase.GetAll()
}

func (o *OrderController) Checkout(orderDto dto.OrderDto) (*presenters.OrderPresenter, error) {
	order, err := o.UseCase.Create(orderDto)

	if err != nil {
		return nil, err
	}

	//TODO: Call payment use case

	return &presenters.OrderPresenter{Id: order.ID}, nil
}

func (o *OrderController) UpdateStatus(id uint32, status string) (*entities.Order, error) {
	order, err := o.UseCase.UpdateStatus(id, status)

	if err != nil {
		return nil, err
	}

	return order, nil
}
