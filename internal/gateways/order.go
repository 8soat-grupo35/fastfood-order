package gateways

import (
	"fmt"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository"
	"log"

	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type orderGateway struct {
	orm *gorm.DB
}

func NewOrderGateway(orm *gorm.DB) repository.OrderRepository {
	return &orderGateway{orm: orm}
}

func (c *orderGateway) GetAll() (orders []entities.Order, err error) {
	expressionOrderBy := fmt.Sprintf(
		"CASE status WHEN '%s' THEN 1 WHEN '%s' THEN 2 WHEN '%s' THEN 3 ELSE 4 END",
		entities.DONE_STATUS,
		entities.IN_PREPARATION_STATUS,
		entities.RECEIVED_STATUS,
	)

	result := c.orm.Preload(clause.Associations).
		Preload("Items.Item").
		Where("status != ?", entities.FINISHED_STATUS).
		Order(expressionOrderBy).
		Order("created_at ASC").
		Find(&orders)

	if result.Error != nil {
		log.Println(result.Error)
		return orders, result.Error
	}

	return orders, err
}

func (c *orderGateway) GetById(id uint32) (*entities.Order, error) {
	order := entities.Order{
		ID: id,
	}
	result := c.orm.Preload(clause.Associations).Preload("Items.Item").First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (c *orderGateway) Create(order entities.Order) (*entities.Order, error) {
	result := c.orm.Create(&order)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &order, nil
}

func (c *orderGateway) Update(id uint32, order entities.Order) (*entities.Order, error) {
	result := c.orm.Session(&gorm.Session{FullSaveAssociations: false}).Updates(&order)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &order, nil
}
