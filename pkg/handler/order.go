package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// OrderIDResponse represents response for order operations
// @Description Response containing order details
// @Success 200 {object} OrderIdResponse
type getOrdersResponse struct {
	Data []models.Order `json:"data"`
}

// @Summary Get all orders
// @Security ApiKeyAuth
// @Tags order
// @Description Получить все заказы пользователя
// @ID get-orders
// @Produce json
// @Success 200 {object} getOrdersResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/order [get]
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

// @Summary Create orders from cart
// @Security ApiKeyAuth
// @Tags order
// @Description Создать заказы для всех товаров в корзине
// @ID create-orders-from-cart
// @Produce json
// @Success 200 {object} statusResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/order/create [post]
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

// @Summary Create order from cart item
// @Security ApiKeyAuth
// @Tags order
// @Description Создать заказ из конкретного товара в корзине
// @ID create-order-from-cart
// @Produce json
// @Param id path string true "Cart Item ID"
// @Success 200 {object} statusResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/order/create/{id} [post]
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

// GetOrdersResponse represents response for user's orders list
// @Description Response containing list of user's orders
// @Success 200 {object} getOrdersResponse
type OrderIdResponse struct {
	Data models.Order `json:"data"`
}

// @Summary Get order by ID
// @Security ApiKeyAuth
// @Tags order
// @Description Получить конкретный заказ по ID
// @ID get-order-by-id
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} OrderIdResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/order/{id} [get]
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

// @Summary Cancel order
// @Security ApiKeyAuth
// @Tags order
// @Description Отменить заказ
// @ID cancel-order
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} statusResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/order/{id}/cancel [post]
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

// @Summary Mark order as delivering
// @Security ApiKeyAuth
// @Tags order
// @Description Отметить заказ как "в процессе доставки"
// @ID delivering-order
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} statusResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/order/{id}/delivering [post]
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

// @Summary Mark order as delivered
// @Security ApiKeyAuth
// @Tags order
// @Description Отметить заказ как "доставленный"
// @ID delivered-order
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} statusResponse
// @Failure 401 {object} handler.Error
// @Failure 500 {object} handler.Error
// @Router /api/order/{id}/delivered [post]
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
