package controllers

import (
	"fmt"
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/gateways"
	controllersInterface "github.com/8soat-grupo35/fastfood-order/internal/interfaces/controllers"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/http"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"
	"github.com/8soat-grupo35/fastfood-order/internal/presenters"
	"github.com/8soat-grupo35/fastfood-order/internal/usecases"

	"gorm.io/gorm"
)

type OrderController struct {
	UseCase             usecase.OrderUseCase
	OrderPaymentUseCase usecase.OrderPaymentUseCase
}

func NewOrderController(db *gorm.DB, httpClient http.Client) controllersInterface.OrderController {
	orderGateway := gateways.NewOrderGateway(db)
	orderPaymentGateway := gateways.NewOrderPaymentGateway(httpClient)
	return &OrderController{
		UseCase:             usecases.NewOrderUseCase(orderGateway),
		OrderPaymentUseCase: usecases.NewOrderPaymentUseCase(orderPaymentGateway),
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

	err = o.OrderPaymentUseCase.Create(*order)
	if err != nil {
		fmt.Println("Error creating order payment")
		return nil, err
	}

	return &presenters.OrderPresenter{Id: order.ID}, nil
}

func (o *OrderController) UpdateStatus(id uint32, status string) (*entities.Order, error) {
	order, err := o.UseCase.UpdateStatus(id, status)

	if err != nil {
		return nil, err
	}

	return order, nil
}
