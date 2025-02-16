package handlers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/controllers"
	controllersInterface "github.com/8soat-grupo35/fastfood-order/internal/interfaces/controllers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	customerController controllersInterface.CustomerController
}

func NewCustomerHandler(db *gorm.DB) CustomerHandler {
	return CustomerHandler{
		customerController: controllers.NewCustomerController(db),
	}
}

func (h *CustomerHandler) GetAll(echo echo.Context) error {
	customers, err := h.customerController.GetAll()

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, customers)
}

// Create godoc
// @Summary      Insert Customer
// @Description  Insert Customer
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        CustomerToInsert	body dto.CustomerDto true "teste"
// @Router       /v1/customer [post]
// @success 200 {array} domain.Customer
// @Failure 500 {object} error
func (h *CustomerHandler) Create(echo echo.Context) error {
	customerDto := dto.CustomerDto{}

	err := echo.Bind(&customerDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	customer, err := h.customerController.Create(customerDto)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) Update(echo echo.Context) error {
	customerDto := dto.CustomerDto{}

	err := echo.Bind(&customerDto)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	customer, err := h.customerController.Update(uint32(id), customerDto)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, customer)
}

func (h *CustomerHandler) Delete(echo echo.Context) error {
	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.customerController.Delete(uint32(id))

	if err != nil {
		return echo.JSON(http.StatusNotFound, err.Error())
	}

	return echo.JSON(http.StatusOK, "customer deleted successfully")
}

// GetByCpf godoc
// @Summary      Get Customer by CPF
// @Description  Retrieve a customer by their CPF
// @Tags         Customers
// @Accept       json
// @Produce      json
// @Param        cpf   path      string  true  "CPF of the customer"
// @Router       /v1/customer/cpf/{cpf} [get]
// @Success      200  {object}  domain.Customer
// @Failure      500  {object}  error
func (h CustomerHandler) GetByCpf(echo echo.Context) error {

	cpf := echo.Param("cpf")

	customer, err := h.customerController.GetByCpf(cpf)

	if err != nil {
		return echo.JSON(http.StatusNotFound, err.Error())
	}

	return echo.JSON(http.StatusOK, customer)
}
