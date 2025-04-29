package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getOrdersResponse struct {
	Data []models.Order `json:"data"`
}

func (h *Handler) getOrders(c *gin.Context) {
	userId, err := getUserId(c)
	var orders []models.Order
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	orders, err = h.services.Order.GetOrders(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getOrdersResponse{
		Data: orders,
	})
}
func (h *Handler) createOrdersFromCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = h.services.Order.CreateOrdersFromCart(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
func (h *Handler) createOrderFromCart(c *gin.Context) {
	userId, err := getUserId(c)
	var cartId string
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	cartId = c.Param("id")
	err = h.services.Order.CreateOrderFromCart(userId, cartId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type OrderIdResponse struct {
	Data models.Order `json:"data"`
}

func (h *Handler) getOrderById(c *gin.Context) {
	userId, err := getUserId(c)
	var order models.Order
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	orderId := c.Param("id")
	order, err = h.services.Order.GetOrderById(userId, orderId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, OrderIdResponse{
		Data: order,
	})
}
func (h *Handler) cancelOrder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	orderId := c.Param("id")
	err = h.services.Order.CancelOrder(userId, orderId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
func (h *Handler) deliveringOrder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	orderId := c.Param("id")
	err = h.services.Order.DeliveringOrder(userId, orderId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
func (h *Handler) deliveredOrder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	orderId := c.Param("id")
	err = h.services.Order.DeliveredOrder(userId, orderId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
