package handler

import (
	"net/http"
	"radical/red_letter/internal/api_error"
	"radical/red_letter/internal/dto"
	"radical/red_letter/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (t *AuthHandler) RegisterHandler(r *gin.Engine) *gin.Engine {
	r.POST("/auth/register", t.RegisterUser)
	return r
}

func (t *AuthHandler) RegisterUser(c *gin.Context) {
	var req dto.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_error.NewApiError(http.StatusBadRequest, "invalid body"))
		return
	}

	err := t.service.RegisterUser(c, req.Name, req.Email, req.Password)
	if err != nil {
		c.Error(api_error.FromError(err))
		return
	}
	res := dto.RegisterUserResponse{}

	c.JSON(http.StatusCreated, JsonSuccessFormater("User Registered Successfully", res))
}
