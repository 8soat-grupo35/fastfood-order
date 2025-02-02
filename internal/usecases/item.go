package usecases

import (
	"errors"
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/usecase"

	custom_errors "github.com/8soat-grupo35/fastfood-order/internal/api/errors"

	"log"
)

type itemService struct {
	itemRepository repository.ItemRepository
}

func NewItemUseCase(itemRepository repository.ItemRepository) usecase.ItemUseCase {
	return &itemService{
		itemRepository: itemRepository,
	}
}

func (service *itemService) GetAll(category string) ([]entities.Item, error) {
	filter := entities.Item{}

	if category != "" {
		filter.Category = category
		err := filter.ValidateCategory()

		if err != nil {
			return []entities.Item{}, &custom_errors.BadRequestError{
				Message: err.Error(),
			}
		}
	}

	items, err := service.itemRepository.GetAll(filter)

	if err != nil {
		return []entities.Item{}, &custom_errors.DatabaseError{
			Message: "get item from repository has failed",
		}
	}

	return items, nil
}

// Create implements ports.ItemService.
func (service *itemService) Create(item dto.ItemDto) (*entities.Item, error) {

	newItem, err := entities.NewItem(item)

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	itemSaved, err := service.itemRepository.Create(*newItem)

	if err != nil {
		return nil, errors.New("create item on repository has failed")
	}

	return itemSaved, err
}

// Update implements ports.ItemService.
func (service *itemService) Update(itemId uint32, item dto.ItemDto) (*entities.Item, error) {

	itemToUpdate, err := entities.NewItem(item)

	if err != nil {
		return nil, &custom_errors.BadRequestError{
			Message: err.Error(),
		}
	}

	itemAlreadySaved, err := service.itemRepository.GetOne(entities.Item{
		ID: itemId,
	})

	if err != nil {
		log.Println(err.Error())
		return nil, &custom_errors.DatabaseError{
			Message: "error on obtain item to update in repository",
		}
	}

	if itemAlreadySaved == nil {
		return nil, &custom_errors.NotFoundError{
			Message: "item not found to update",
		}
	}

	itemUpdated, err := service.itemRepository.Update(itemId, *itemToUpdate)

	if err != nil {
		log.Println(err.Error())
		return nil, &custom_errors.DatabaseError{
			Message: "updated item on repository has failed",
		}
	}

	return itemUpdated, err
}

// Delete implements ports.ItemService.
func (service *itemService) Delete(itemId uint32) error {
	itemAlreadySaved, err := service.itemRepository.GetOne(entities.Item{
		ID: itemId,
	})

	if err != nil {
		log.Println(err.Error())
		return &custom_errors.DatabaseError{
			Message: "error on obtain item to delete in repository",
		}
	}

	if itemAlreadySaved == nil {
		return &custom_errors.NotFoundError{
			Message: "item not found to delete",
		}
	}

	err = service.itemRepository.Delete(itemId)

	if err != nil {
		return &custom_errors.DatabaseError{
			Message: "error on delete in repository",
		}
	}

	return err
}
