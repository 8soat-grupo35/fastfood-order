package usecases

import (
	"errors"
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"

	custom_errors "github.com/8soat-grupo35/fastfood-order/internal/api/errors"
)

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderUseCase(orderRepository repository.OrderRepository) usecase.OrderUseCase {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (service *orderService) GetAll() ([]entities.Order, error) {
	orders, err := service.orderRepository.GetAll()

	if err != nil {
		return []entities.Order{}, &custom_errors.DatabaseError{
			Message: "get order from repository has failed",
		}
	}

	return orders, nil
}

// Create implements ports.OrderService.
func (service *orderService) Create(order dto.OrderDto) (*entities.Order, error) {
	newOrder, err := entities.NewOrder(order)

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	orderSaved, err := service.orderRepository.Create(*newOrder)

	if err != nil {
		return nil, errors.New("create order on repository has failed")
	}

	return orderSaved, err
}

func (service *orderService) UpdateStatus(id uint32, status string) (*entities.Order, error) {
	order, err := service.orderRepository.GetById(id)
	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	order.Status = status
	validateError := order.Validate()
	if validateError != nil {
		return nil, errors.New(validateError.Error())
	}

	orderSaved, err := service.orderRepository.Update(id, *order)
	if err != nil {
		return nil, errors.New("update order on  repository has failed")
	}

	return orderSaved, err
}
