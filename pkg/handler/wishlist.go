package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetWishlistItemsResponse represents response for wishlist items
// @Description Response containing list of wishlist products
// @Success 200 {object} getWishItemsResponse
type getWishItemsResponse struct {
	Data []models.Item
}

// @Summary Add item to wishlist
// @Security ApiKeyAuth
// @Tags wishlist
// @Description Добавить товар в вишлист пользователя
// @ID add-to-wishlist
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Item ID"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/wishlist/{id} [post]
func (h *Handler) addToWishlist(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	itemId := c.Param("id")
	id, err := h.services.WishList.AddToWishlist(userId, itemId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get user's wishlist
// @Security ApiKeyAuth
// @Tags wishlist
// @Description Получить вишлист пользователя
// @ID get-
// @Security ApiKeyAuth
// @Produce  json
// @Success 200 {object} getWishItemsResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/wishlist [get]
func (h *Handler) getWishlist(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	wishItems, err := h.services.WishList.GetWishlist(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getWishItemsResponse{
		Data: wishItems,
	})
}

// @Summary Delete item from wishlist
// @Security ApiKeyAuth
// @Tags wishlist
// @Description Удалить товар из вишлиста пользователя
// @ID delete-from-wishlist
// @Security ApiKeyAuth
// @Param id path string true "Item ID"
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/wishlist/{id} [delete]
func (h *Handler) deleteFromWishlist(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	itemId := c.Param("id")
	err = h.services.WishList.DeleteWishItem(userId, itemId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
