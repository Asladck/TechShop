package handler

import (
	"TechShop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      Регистрация пользователя
// @Description  Создаёт нового пользователя и возвращает его ID
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      models.User  true  "Информация о пользователе"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {object}  handler.Error
// @Failure      500    {object}  handler.Error
// @Router       /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password_hash" binding:"required"`
}

// @Summary      Авторизация пользователя
// @Description  Возвращает access и refresh токены
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      signInInput  true  "Данные для входа"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {object}  handler.Error
// @Failure      500    {object}  handler.Error
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.
			StatusBadRequest, err.Error())
		return
	}
	tokenA, tokenR, err := h.services.Authorization.GenerateToken(input.Username, input.Password, input.Email)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  tokenA,
		"refresh_token": tokenR,
	})
}

type refreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// @Summary      Обновление Access токена
// @Description  Генерирует новый access токен по refresh токену
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        input  body      refreshInput  true  "Refresh токен"
// @Success      200    {object}  map[string]interface{}
// @Failure      400    {object}  handler.Error
// @Failure      500    {object}  handler.Error
// @Router       /auth/refresh [post]
func (h *Handler) refreshHandler(c *gin.Context) {
	var input refreshInput
	if err := c.ShouldBindJSON(&input); err != nil {
		NewErrorResponse(c, http.
			StatusBadRequest, err.Error())
		return
	}
	userId, err := h.services.Authorization.ParseRefToken(input.RefreshToken)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	newAccessToken, err := h.services.Authorization.GenerateAccToken(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": newAccessToken,
	})
}
