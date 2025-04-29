package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addToCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	var input models.CartItem
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	cartId, err := h.services.Cart.AddToCart(userId, input.ItemID, input.ItemCount)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": cartId,
	})
}

func (h *Handler) getCartItemById(c *gin.Context) {
	var cart models.Cart
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	id := c.Param("id")
	cart, err = h.services.Cart.GetCartItemById(userId, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetCartItemResponse{
		Data: cart,
	})
}

func (h *Handler) getCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	var cartItems []models.Cart
	cartItems, err = h.services.Cart.GetCart(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, GetAllCartItemResponse{
		Data: cartItems,
	})
}
func (h *Handler) updateCartItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	cartId := c.Param("id")
	var input models.CartUpdate
	err = c.ShouldBindJSON(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Cart.Update(userId, cartId, input.ItemCount)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
func (h *Handler) deleteCartItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	cartId := c.Param("id")
	err = h.services.Cart.Delete(userId, cartId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
