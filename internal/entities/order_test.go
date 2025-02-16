package entities

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewOrderCreatesValidOrder(t *testing.T) {
	orderDto := dto.OrderDto{
		CustomerID: 1,
		Items: []dto.OrderItemDto{
			{Id: 1, Quantity: 2},
		},
	}

	order, err := NewOrder(orderDto)

	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, uint32(1), order.CustomerID)
	assert.Equal(t, RECEIVED_STATUS, order.Status)
	assert.Len(t, order.Items, 1)
	assert.Equal(t, uint32(1), order.Items[0].ItemID)
	assert.Equal(t, uint32(2), order.Items[0].Quantity)
}

func TestNewOrderReturnsErrorForInvalidOrder(t *testing.T) {
	orderDto := dto.OrderDto{
		CustomerID: 0,
		Items:      []dto.OrderItemDto{},
	}

	order, err := NewOrder(orderDto)

	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestValidateReturnsErrorForInvalidStatus(t *testing.T) {
	order := Order{
		CustomerID: 1,
		Items: []OrderItem{
			{ItemID: 1, Quantity: 2},
		},
		Status: "INVALID_STATUS",
	}

	err := order.Validate()

	assert.Error(t, err)
}

func TestValidateReturnsErrorForMissingCustomerID(t *testing.T) {
	order := Order{
		CustomerID: 0,
		Items: []OrderItem{
			{ItemID: 1, Quantity: 2},
		},
		Status: RECEIVED_STATUS,
	}

	err := order.Validate()

	assert.Error(t, err)
}

func TestValidateReturnsErrorForEmptyItems(t *testing.T) {
	order := Order{
		CustomerID: 1,
		Items:      []OrderItem{},
		Status:     RECEIVED_STATUS,
	}

	err := order.Validate()

	assert.Error(t, err)
}

func TestValidateReturnsNoErrorForValidOrder(t *testing.T) {
	order := Order{
		CustomerID: 1,
		Items: []OrderItem{
			{ItemID: 1, Quantity: 2},
		},
		Status:    RECEIVED_STATUS,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := order.Validate()

	assert.NoError(t, err)
}
