package entities

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCustomerCreatesValidCustomer(t *testing.T) {
	customerDto := dto.CustomerDto{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		CPF:   "12345678901",
	}

	customer, err := NewCustomer(customerDto)

	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "John Doe", customer.Name)
	assert.Equal(t, "john.doe@example.com", customer.Email)
	assert.Equal(t, "12345678901", customer.CPF)
}

func TestNewCustomerReturnsErrorForInvalidName(t *testing.T) {
	customerDto := dto.CustomerDto{
		Name:  "John",
		Email: "john.doe@example.com",
		CPF:   "12345678901",
	}

	customer, err := NewCustomer(customerDto)

	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestNewCustomerReturnsErrorForInvalidEmail(t *testing.T) {
	customerDto := dto.CustomerDto{
		Name:  "John Doe",
		Email: "invalid-email",
		CPF:   "12345678901",
	}

	customer, err := NewCustomer(customerDto)

	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestNewCustomerReturnsErrorForInvalidCPF(t *testing.T) {
	customerDto := dto.CustomerDto{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		CPF:   "1234567890",
	}

	customer, err := NewCustomer(customerDto)

	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestValidateReturnsErrorForInvalidName(t *testing.T) {
	customer := Customer{
		Name:  "John",
		Email: "john.doe@example.com",
		CPF:   "12345678901",
	}

	err := customer.Validate()

	assert.Error(t, err)
}

func TestValidateReturnsErrorForInvalidEmail(t *testing.T) {
	customer := Customer{
		Name:  "John Doe",
		Email: "invalid-email",
		CPF:   "12345678901",
	}

	err := customer.Validate()

	assert.Error(t, err)
}

func TestValidateReturnsErrorForInvalidCPF(t *testing.T) {
	customer := Customer{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		CPF:   "1234567890",
	}

	err := customer.Validate()

	assert.Error(t, err)
}

func TestValidateReturnsNoErrorForValidCustomer(t *testing.T) {
	customer := Customer{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		CPF:   "12345678901",
	}

	err := customer.Validate()

	assert.NoError(t, err)
}
