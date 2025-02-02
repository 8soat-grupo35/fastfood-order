package gateways

import (
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository"
	"log"

	"gorm.io/gorm"
)

type itemGateway struct {
	orm *gorm.DB
}

func NewItemGateway(orm *gorm.DB) repository.ItemRepository {

	return &itemGateway{orm: orm}
}

func (c *itemGateway) GetAll(filter entities.Item) (items []entities.Item, err error) {
	result := c.orm.Where(filter).Find(&items)

	if result.Error != nil {
		log.Println(result.Error)
		return items, result.Error
	}

	return items, err
}

func (c *itemGateway) GetOne(itemFilter entities.Item) (item *entities.Item, err error) {
	result := c.orm.Where(itemFilter).First(&item)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return item, nil
}

func (c *itemGateway) Create(item entities.Item) (*entities.Item, error) {
	result := c.orm.Create(&item)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &item, nil
}

func (c *itemGateway) Update(itemId uint32, item entities.Item) (*entities.Item, error) {
	itemModel := entities.Item{ID: itemId}
	result := c.orm.Model(&itemModel).Updates(&item)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &itemModel, nil
}

func (c *itemGateway) Delete(itemId uint32) error {
	result := c.orm.Delete(&entities.Item{}, itemId)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
