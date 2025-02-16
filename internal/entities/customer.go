package entities

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gorm.io/gorm"
)

type Customer struct {
	ID    uint32 `gorm:"primary_key;auto_increment"`
	Name  string `gorm:"size:255;not null;"`
	Email string `gorm:"size:255;not null;"`
	CPF   string `gorm:"size:11;not null;"`
	gorm.Model
} //@name domain.Customer

func NewCustomer(customer dto.CustomerDto) (*Customer, error) {
	newCustomer := Customer{
		Name:  customer.Name,
		Email: customer.Email,
		CPF:   customer.CPF,
	}

	err := newCustomer.Validate()

	if err != nil {
		return nil, err
	}

	return &newCustomer, err
}

func (c Customer) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(
			&c.Name,
			validation.Required,
			validation.Length(5, 255),
		),
		validation.Field(
			&c.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(
			&c.CPF,
			validation.Required,
			validation.Length(11, 11),
			is.Digit,
		),
	)
}
