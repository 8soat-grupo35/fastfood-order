package usecases

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"
)

type orderPaymentUseCase struct {
	orderPaymentRepository repository.OrderPaymentRepository
}

func NewOrderPaymentUseCase(orderPaymentRepository repository.OrderPaymentRepository) usecase.OrderPaymentUseCase {
	return &orderPaymentUseCase{orderPaymentRepository: orderPaymentRepository}
}

func (o *orderPaymentUseCase) Create(order entities.Order) error {
	newOrderPayment := dto.OrderPaymentDto{
		OrderID: int(order.ID),
	}
	return o.orderPaymentRepository.Create(newOrderPayment)
}
