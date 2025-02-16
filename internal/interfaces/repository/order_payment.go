package repository

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
)

//go:generate mockgen -source=order_payment.go -destination=mock/order_payment.go
type OrderPaymentRepository interface {
	Create(orderPayment dto.OrderPaymentDto) error
}
