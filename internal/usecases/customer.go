package usecases

import (
	"errors"
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"

	custom_errors "github.com/8soat-grupo35/fastfood-order/internal/api/errors"
)

type customerUseCase struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerUseCase(customerRepository repository.CustomerRepository) usecase.CustomerUseCase {
	return &customerUseCase{
		customerRepository: customerRepository,
	}
}

func (useCase *customerUseCase) GetAll() ([]entities.Customer, error) {
	customers, err := useCase.customerRepository.GetAll()

	if err != nil {
		return []entities.Customer{}, errors.New("get customer from repository has failed")
	}

	return customers, nil
}

// Create implements ports.CustomerService.
func (useCase *customerUseCase) Create(customer dto.CustomerDto) (*entities.Customer, error) {
	newCustomer, err := entities.NewCustomer(customer)

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	customerSaved, err := useCase.customerRepository.Create(*newCustomer)

	if err != nil {
		return nil, errors.New("create customer on repository has failed")
	}

	return customerSaved, err
}

func (useCase *customerUseCase) GetByCpf(cpf string) (*entities.Customer, error) {

	customer, err := useCase.customerRepository.GetOne(entities.Customer{
		CPF: cpf,
	})

	if err != nil {
		return nil, errors.New("error on obtain customer by CPF in repository")
	}

	return customer, err
}

// Update implements ports.CustomerService.
func (useCase *customerUseCase) Update(customerId uint32, customer dto.CustomerDto) (*entities.Customer, error) {

	customerToUpdate, err := entities.NewCustomer(customer)

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	customerAlreadySaved, err := useCase.customerRepository.GetOne(entities.Customer{
		ID: customerId,
	})

	if err != nil {
		return nil, errors.New("error on obtain customer to update in repository")
	}

	if customerAlreadySaved == nil {
		return nil, errors.New("customer not found to update")
	}

	customerUpdated, err := useCase.customerRepository.Update(customerId, *customerToUpdate)

	if err != nil {
		return nil, errors.New("updated customer on repository has failed")
	}

	return customerUpdated, err
}

// Delete implements ports.CustomerService.
func (useCase *customerUseCase) Delete(customerId uint32) error {
	customerAlreadySaved, err := useCase.customerRepository.GetOne(entities.Customer{
		ID: customerId,
	})

	if err != nil {
		return errors.New("error on obtain customer to delete in repository")
	}

	if customerAlreadySaved == nil {
		return errors.New("customer not found to delete")
	}

	err = useCase.customerRepository.Delete(customerId)

	if err != nil {
		return errors.New("error on delete in repository")
	}

	return err
}
