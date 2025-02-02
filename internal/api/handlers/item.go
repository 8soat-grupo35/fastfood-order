package handlers

import (
	"github.com/8soat-grupo35/fastfood-order/internal/adapters/dto"
	"github.com/8soat-grupo35/fastfood-order/internal/controllers"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	itemController *controllers.ItemController
}

func NewItemHandler(db *gorm.DB) ItemHandler {
	return ItemHandler{
		itemController: controllers.NewItemController(db),
	}
}

// GetAll godoc
// @Summary      List Items
// @Description  List All Items
// @Tags         Items
// @Accept       json
// @Produce      json
// @Router       /v1/item [get]
// @success 200  {object} domain.Item
// @Failure 500 {object} error
func (h *ItemHandler) GetAll(echo echo.Context) error {
	category := echo.QueryParam("category")

	items, err := h.itemController.GetAllByCategory(category)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, items)
}

// Create godoc
// @Summary      Insert Item
// @Description  Insert Item
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param        ItemToInsert	body dto.ItemDto true "teste"
// @Router       /v1/item [post]
// @success 200 {array} domain.Item
// @Failure 500 {object} error
func (h *ItemHandler) Create(echo echo.Context) error {
	itemDto := dto.ItemDto{}

	err := echo.Bind(&itemDto)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	item, err := h.itemController.Create(itemDto)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, item)
}

// Update godoc
// @Summary      Update Item
// @Description  Update Item
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param		 id             path int         true "ID do item"
// @Param        ItemToInsert	body dto.ItemDto true "teste"
// @Router       /v1/item/{id} [put]
// @success 200 {array} domain.Item
// @Failure 500 {object} error
func (h *ItemHandler) Update(echo echo.Context) error {
	itemDto := dto.ItemDto{}

	err := echo.Bind(&itemDto)
	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	item, err := h.itemController.Update(id, itemDto)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, item)
}

// Delete godoc
// @Summary      Delete Item
// @Description  Delete Item
// @Tags         Items
// @Accept       json
// @Produce      json
// @Param		 id             path int         true "ID do item"
// @Router       /v1/item/{id} [delete]
// @success 200 {string}  string    "item deleted successfully"
// @Failure 500 {object} error
func (h *ItemHandler) Delete(echo echo.Context) error {
	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.itemController.Delete(id)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	return echo.JSON(http.StatusOK, "item deleted successfully")
}
