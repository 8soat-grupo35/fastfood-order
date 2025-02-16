package controllers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/gateways"
	controllersInterface "github.com/8soat-grupo35/fastfood-order/internal/interfaces/controllers"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"
	"github.com/8soat-grupo35/fastfood-order/internal/usecases"
	"gorm.io/gorm"
)

type CustomerController struct {
	UseCase usecase.CustomerUseCase
}

func NewCustomerController(db *gorm.DB) controllersInterface.CustomerController {
	gateway := gateways.NewCustomerGateway(db)
	return &CustomerController{
		UseCase: usecases.NewCustomerUseCase(gateway),
	}
}

func (c *CustomerController) GetAll() ([]entities.Customer, error) {
	return c.UseCase.GetAll()
}

func (c *CustomerController) GetByCpf(cpf string) (*entities.Customer, error) {
	return c.UseCase.GetByCpf(cpf)
}

func (c *CustomerController) Create(customer dto.CustomerDto) (*entities.Customer, error) {
	return c.UseCase.Create(customer)
}

func (c *CustomerController) Update(customerID uint32, customer dto.CustomerDto) (*entities.Customer, error) {
	return c.UseCase.Update(customerID, customer)
}

func (c *CustomerController) Delete(customerID uint32) error {
	return c.UseCase.Delete(customerID)
}
