package gateways

import (
	"github.com/8soat-grupo35/fastfood-order/internal/entities"
	"github.com/8soat-grupo35/fastfood-order/internal/interfaces/repository"
	"gorm.io/gorm"
	"log"
)

type customerGateway struct {
	orm *gorm.DB
}

func NewCustomerGateway(orm *gorm.DB) repository.CustomerRepository {
	return &customerGateway{orm: orm}
}

func (c *customerGateway) GetAll() (customers []entities.Customer, err error) {
	result := c.orm.Find(&customers)

	if result.Error != nil {
		log.Println(result.Error)
		return customers, result.Error
	}

	return customers, err
}

func (c *customerGateway) GetOne(customerFilter entities.Customer) (customer *entities.Customer, err error) {
	result := c.orm.Where(customerFilter).First(&customer)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return customer, nil
}

func (c *customerGateway) Create(customer entities.Customer) (*entities.Customer, error) {
	result := c.orm.Create(&customer)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &customer, nil
}

func (c *customerGateway) Update(customerId uint32, customer entities.Customer) (*entities.Customer, error) {
	customerModel := entities.Customer{ID: customerId}
	result := c.orm.Model(&customerModel).Updates(&customer)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &customerModel, nil
}

func (c *customerGateway) Delete(customerId uint32) error {
	result := c.orm.Delete(&entities.Customer{}, customerId)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
