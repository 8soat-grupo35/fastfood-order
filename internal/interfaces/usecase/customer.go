package usecase

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
)

type CustomerUseCase interface {
	GetAll() ([]entities.Customer, error)
	Create(dto.CustomerDto) (*entities.Customer, error)
	GetByCpf(cpf string) (*entities.Customer, error)
	Update(customerId uint32, customer dto.CustomerDto) (*entities.Customer, error)
	Delete(customerId uint32) error
}
