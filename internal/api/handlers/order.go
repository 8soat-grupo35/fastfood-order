package handlers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/controllers"
	controllersInterface "github.com/8soat-grupo35/fastfood-order/internal/interfaces/controllers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderHandler struct {
	orderController controllersInterface.OrderController
}

func NewOrderHandler(db *gorm.DB) OrderHandler {
	return OrderHandler{
		orderController: controllers.NewOrderController(db),
	}
}

// GetAll godoc
// @Summary      List Orders
// @Description  List All Orders
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Router       /v1/orders [get]
// @Success 200  {object} domain.Order
// @Failure 500  {object} error
func (h *OrderHandler) GetAll(echo echo.Context) error {
	orders, err := h.orderController.GetAll()

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, orders)
}

// Create godoc
// @Summary      Insert Order
// @Description  Insert Order
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param        Order	body dto.OrderDto true "Order to create"
// @Router       /v1/orders/checkout [post]
// @success 200 {array} presenters.OrderPresenter
// @Failure 500 {object} error
func (h *OrderHandler) Checkout(echo echo.Context) error {
	orderDto := dto.OrderDto{}

	err := echo.Bind(&orderDto)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	order, err := h.orderController.Checkout(orderDto)
	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, order)
}

// Create godoc
// @Summary      Update Order Status
// @Description  Update Order Status
// @Tags         Orders
// @Accept       json
// @Produce      json
// @Param		 		 id     path int         true "ID do item"
// @Param        Order	body dto.OrderStatusDto true "Status to update Order"
// @Router       /v1/orders/{id} [patch]
// @success 200 {object} domain.Order
// @Failure 500 {object} error
func (h *OrderHandler) UpdateStatus(echo echo.Context) error {
	orderDto := dto.OrderDto{}

	id, err := strconv.Atoi(echo.Param("id"))
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	bindError := echo.Bind(&orderDto)
	if bindError != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	order, err := h.orderController.UpdateStatus(uint32(id), orderDto.Status)
	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	if order == nil {
		return echo.JSON(http.StatusNoContent, err.Error())
	}

	return echo.JSON(http.StatusOK, order)
}
