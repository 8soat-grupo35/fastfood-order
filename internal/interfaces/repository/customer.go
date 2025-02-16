package repository

import "github.com/8soat-grupo35/fastfood-order/internal/entities"

//go:generate mockgen -source=customer.go -destination=mock/customer.go
type CustomerRepository interface {
	GetAll() ([]entities.Customer, error)
	GetOne(entities.Customer) (*entities.Customer, error)
	Create(customer entities.Customer) (*entities.Customer, error)
	Update(customerId uint32, customer entities.Customer) (*entities.Customer, error)
	Delete(customerId uint32) error
}
