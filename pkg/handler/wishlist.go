package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getWishItemsResponse struct {
	Data []models.Item
}

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
