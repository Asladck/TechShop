package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "UserId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusBadRequest, "empty auth head")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusBadRequest, "invalid auth header")
		return
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}
func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return "", errors.New("user id not found")
	}
	idInt, ok := id.(string)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is invalid type")
		return "", errors.New("user id not found")
	}
	return idInt, nil
}
