package usecase

import (
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
)

//go:generate mockgen -source=order_payment.go -destination=mock/order_payment.go
type OrderPaymentUseCase interface {
	Create(order entities.Order) error
}
