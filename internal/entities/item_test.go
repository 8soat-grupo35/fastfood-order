package entities

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewItemCreatesValidItem(t *testing.T) {
	itemDto := dto.ItemDto{
		Name:     "Burger",
		Category: "LANCHE",
		Price:    10.0,
		ImageUrl: "http://image.com",
	}

	item, err := NewItem(itemDto)

	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, "Burger", item.Name)
	assert.Equal(t, "LANCHE", item.Category)
	assert.Equal(t, float32(10.0), item.Price)
	assert.Equal(t, "http://image.com", item.ImageUrl)
}

func TestNewItemReturnsErrorForInvalidCategory(t *testing.T) {
	itemDto := dto.ItemDto{
		Name:     "Burger",
		Category: "INVALID",
		Price:    10.0,
		ImageUrl: "http://image.com",
	}

	item, err := NewItem(itemDto)

	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestNewItemReturnsErrorForInvalidPrice(t *testing.T) {
	itemDto := dto.ItemDto{
		Name:     "Burger",
		Category: "LANCHE",
		Price:    -1.0,
		ImageUrl: "http://image.com",
	}

	item, err := NewItem(itemDto)

	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestNewItemReturnsErrorForInvalidImageUrl(t *testing.T) {
	itemDto := dto.ItemDto{
		Name:     "Burger",
		Category: "LANCHE",
		Price:    10.0,
		ImageUrl: "invalid-url",
	}

	item, err := NewItem(itemDto)

	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestValidateCategoryReturnsErrorForInvalidCategory(t *testing.T) {
	item := Item{
		Category: "INVALID",
	}

	err := item.ValidateCategory()

	assert.Error(t, err)
}

func TestValidateCategoryReturnsNoErrorForValidCategory(t *testing.T) {
	item := Item{
		Category: "LANCHE",
	}

	err := item.ValidateCategory()

	assert.NoError(t, err)
}

func TestValidateReturnsErrorForInvalidItemName(t *testing.T) {
	item := Item{
		Name:     "B",
		Category: "LANCHE",
		Price:    10.0,
		ImageUrl: "http://image.com",
	}

	err := item.Validate()

	assert.Error(t, err)
}

func TestValidateReturnsErrorForInvalidPrice(t *testing.T) {
	item := Item{
		Name:     "Burger",
		Category: "LANCHE",
		Price:    -1.0,
		ImageUrl: "http://image.com",
	}

	err := item.Validate()

	assert.Error(t, err)
}

func TestValidateReturnsErrorForInvalidImageUrl(t *testing.T) {
	item := Item{
		Name:     "Burger",
		Category: "LANCHE",
		Price:    -1.0,
		ImageUrl: "invalid-url",
	}

	err := item.Validate()

	assert.Error(t, err)
}

func TestValidateReturnsNoErrorForValidItem(t *testing.T) {
	item := Item{
		Name:     "Burger",
		Category: "LANCHE",
		Price:    10.0,
		ImageUrl: "http://image.com",
	}

	err := item.Validate()

	assert.NoError(t, err)
}
