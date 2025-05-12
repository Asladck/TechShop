package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get total price in cart
// @Security ApiKeyAuth
// @Tags buy
// @Description Получить общую стоимость товаров в корзине
// @ID get-price-cart
// @Produce json
// @Success 200 {object} statusFloat
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/cart/price [get]
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

// @Summary Buy single item
// @Security ApiKeyAuth
// @Tags buy
// @Description Купить конкретный товар (из корзины или напрямую)
// @ID buy-one-item
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Param input body models.Stock true "Stock information"
// @Success 200 {object} statusResponse
// @Failure 400 {object} handler.Error
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/buy/{id} [post]
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
