package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Error represents an API error response
// @Description API error information
type Error struct {
	Message string `json:"message"`
}

// StatusResponse represents a simple status response
// @Description Basic status response
type statusResponse struct {
	Status string `json:"status"`
}

// StatusFloat represents a numeric status response
// @Description Numeric status response
type statusFloat struct {
	Status float64 `json:"status"`
}

// NewErrorResponse logs error and returns error response
func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error()
	c.AbortWithStatusJSON(statusCode, Error{message})
}

// GetAllCartItemResponse represents response for getting all cart items
// @Description Response with list of cart items including product details
type GetAllCartItemResponse struct {
	Data []models.Cart `json:"data"`
}

// GetCartItemResponse represents response for getting single cart item
// @Description Response with single cart item details
type GetCartItemResponse struct {
	Data models.Cart `json:"data"`
}
