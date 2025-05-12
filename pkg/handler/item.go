package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllItemsResponse represents response for listing all products
// @Description Response containing list of all products
// @Success 200 {object} getAllItemResponse
type getAllItemResponse struct {
	Data []models.Item `json:"data"`
}

// @Summary      Получить все товары
// @Description  Возвращает список всех товаров
// @Tags         items
// @Produce      json
// @Success      200  {object}  getAllItemResponse
// @Failure      500  {object}  handler.Error
// @Router       /items [get]
func (h *Handler) getItems(c *gin.Context) {
	items, err := h.services.Item.GetItems()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllItemResponse{
		Data: items,
	})
}

// GetItemResponse represents response for single product
// @Description Response containing detailed product information
// @Success 200 {object} getItemResponse
type getItemResponse struct {
	Data models.Item `json:"data"`
}

// @Summary      Получить товар по ID
// @Description  Возвращает один товар по его идентификатору
// @Tags         items
// @Produce      json
// @Param        id   path      string  true  "Item ID"
// @Success      200  {object}  getItemResponse
// @Failure      500  {object}  handler.Error
// @Router       /items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	id := c.Param("id")
	item, err := h.services.Item.GetItemById(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getItemResponse{
		Data: item,
	})
}
