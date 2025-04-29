package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getPriceInCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	var total float64
	total, err = h.services.Buy.GetPriceCart(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusFloat{
		Status: total,
	})
}
func (h *Handler) buyOneItem(c *gin.Context) {
	var stock models.Stock
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = c.ShouldBindJSON(&stock)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	itemId := c.Param("id")
	err = h.services.Buy.BuyItem(userId, itemId, stock.Stock)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "item was bought",
	})
}
