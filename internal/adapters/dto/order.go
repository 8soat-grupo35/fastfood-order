package dto

const (
	PAYMENT_STATUS_WAITING = 1
)

type OrderItemDto struct {
	Id       uint32 `json:"id"`
	Quantity uint32 `json:"quantity"`
} //@name OrderItemDto

type OrderDto struct {
	Items      []OrderItemDto `json:"items"`
	CustomerID uint32         `json:"customer_id"`
	Status     string         `json:"status"`
} //@name OrderDto

type OrderStatusDto struct {
	Status string `json:"status"`
} //@name OrderStatusDto

type OrderPaymentStatusDto struct {
	Status string `json:"status"`
} //@name OrderPaymentStatusDto

type OrderPaymentDto struct {
	OrderID         int `json:"id"`
	PaymentStatusID int `json:"payment_status_id"`
} //@name OrderPaymentDto
