package entities

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gorm.io/gorm"
)

type Item struct {
	ID       uint32  `gorm:"primary_key;auto_increment"`
	Name     string  `gorm:"size:255;not null;"`
	Category string  `gorm:"size:30;not null;"`
	Price    float32 `gorm:"not null;"`
	ImageUrl string  `gorm:"size:255;not null;"`
	gorm.Model
} //@name domain.Item

func (item Item) ValidateCategory() error {

	return validation.ValidateStruct(
		&item,
		validation.Field(
			&item.Category,
			validation.In(item.allowedCategories()...).Error("must be a valid value between (lanche,sobremesa,acompanhamento,bebida)"),
		),
	)
}

func (item Item) allowedCategories() []interface{} {
	return []interface{}{
		"LANCHE",
		"ACOMPANHAMENTO",
		"BEBIDA",
		"SOBREMESA",
	}
}

func (item Item) Validate() error {
	allowedCategories := item.allowedCategories()
	return validation.ValidateStruct(
		&item,
		validation.Field(
			&item.Name,
			validation.Required,
			validation.Length(3, 255),
		),
		validation.Field(
			&item.Category,
			validation.Required,
			validation.In(allowedCategories...).Error("must be a valid value between (lanche,sobremesa,acompanhamento,bebida)"),
		),
		validation.Field(
			&item.Price,
			validation.Required,
			validation.Min(0.01),
		),
		validation.Field(
			&item.ImageUrl,
			validation.Required,
			is.URL,
		),
	)
}

func NewItem(item dto.ItemDto) (*Item, error) {
	newItem := Item{
		Name:     item.Name,
		Category: strings.ToUpper(item.Category),
		Price:    item.Price,
		ImageUrl: item.ImageUrl,
	}

	err := newItem.Validate()

	if err != nil {
		return nil, err
	}

	return &newItem, err
}
