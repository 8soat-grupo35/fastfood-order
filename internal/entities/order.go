package entities

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	RECEIVED_STATUS       = "RECEBIDO"
	IN_PREPARATION_STATUS = "EM_PREPARACAO"
	DONE_STATUS           = "PRONTO"
	FINISHED_STATUS       = "FINALIZADO"
)

type OrderItem struct {
	ID       uint32 `gorm:"primarykey;autoIncrement" json:"-"`
	OrderID  uint32 `json:"-"`
	ItemID   uint32 `json:"id"`
	Quantity uint32 `json:"quantity"`
	Item     Item   `gorm:"references:ID" json:"-"`
} //@name domain.OrderItem

type Order struct {
	ID         uint32      `gorm:"primarykey;autoIncrement" json:"id"`
	Items      []OrderItem `gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE" json:"items"`
	CustomerID uint32      `json:"customer_id"`
	Status     string      `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
} //@name domain.Order

func NewOrder(orderDto dto.OrderDto) (*Order, error) {
	newOrder := Order{
		CustomerID: orderDto.CustomerID,
		Items:      OrderItemToDomain(orderDto),
		Status:     RECEIVED_STATUS,
	}

	err := newOrder.Validate()

	if err != nil {
		return nil, err
	}

	return &newOrder, err
}

func OrderItemToDomain(orderDto dto.OrderDto) (list []OrderItem) {

	for _, orderItemDto := range orderDto.Items {
		list = append(list, OrderItem{
			ItemID:   orderItemDto.Id,
			Quantity: orderItemDto.Quantity,
		})
	}

	return list
}

func (order Order) Validate() error {
	return validation.ValidateStruct(
		&order,
		validation.Field(
			&order.Items,
			validation.Required.When(len(order.Items) > 0).Error("must be one or more item"),
			validation.Required,
		),
		validation.Field(
			&order.CustomerID,
			validation.Required,
		),
		validation.Field(
			&order.Status,
			validation.Required,
			validation.In(DONE_STATUS, IN_PREPARATION_STATUS, RECEIVED_STATUS, FINISHED_STATUS),
		),
	)
}
