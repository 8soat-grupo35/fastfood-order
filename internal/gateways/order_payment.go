package gateways

import (
	"bytes"
	"encoding/json"
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"

	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/http"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository"
)

type orderPaymentGateway struct {
	client http.Client
}

func NewOrderPaymentGateway(client http.Client) repository.OrderPaymentRepository {
	return &orderPaymentGateway{client: client}
}

func (o orderPaymentGateway) Create(orderPayment dto.OrderPaymentDto) error {
	orderData, err := json.Marshal(orderPayment)
	if err != nil {
		return err
	}
	_, err = o.client.Post("/order-payment", bytes.NewReader(orderData))
	if err != nil {
		return err
	}
	return nil
}
