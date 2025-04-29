package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}
type statusResponse struct {
	Status string `json:"status"`
}
type statusFloat struct {
	Status float64 `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error()
	c.AbortWithStatusJSON(statusCode, Error{message})
}

type GetAllCartItemResponse struct {
	Data []models.Cart `json:"data"`
}
type GetCartItemResponse struct {
	Data models.Cart `json:"data"`
}
