package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Add item to cart
// @Security ApiKeyAuth
// @Tags cart
// @Description Добавить товар в корзину пользователя
// @ID add-to-cart
// @Accept  json
// @Produce  json
// @Param input body models.CartItem true "item info"
// @Success 200 {object} map[string]interface{} "Returns cart ID"
// @Failure 400 {object} handler.Error
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/cart [post]
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

// @Summary Get cart item by ID
// @Security ApiKeyAuth
// @Tags cart
// @Description Получить конкретный товар из корзины
// @ID get-cart-item
// @Produce  json
// @Param id path string true "Cart Item ID"
// @Success 200 {object} GetCartItemResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/cart/{id} [get]
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

// @Summary Get all cart items
// @Security ApiKeyAuth
// @Tags cart
// @Description Получить все товары в корзине пользователя
// @ID get-cart
// @Produce  json
// @Success 200 {object} GetAllCartItemResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/cart [get]
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

// @Summary Update cart item
// @Security ApiKeyAuth
// @Tags cart
// @Description Обновить количество товара в корзине
// @ID update-cart-item
// @Accept  json
// @Produce  json
// @Param id path string true "Cart Item ID"
// @Param input body models.CartUpdate true "update info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} handler.Error
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/cart/{id} [put]
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

// @Summary Delete cart item
// @Security ApiKeyAuth
// @Tags cart
// @Description Удалить товар из корзины
// @ID delete-cart-item
// @Produce  json
// @Param id path string true "Cart Item ID"
// @Success 200 {object} statusResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/cart/{id} [delete]
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
