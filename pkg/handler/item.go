package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllItemResponse struct {
	Data []models.Item `json:"data"`
}

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

type getItemResponse struct {
	Data models.Item `json:"data"`
}

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
