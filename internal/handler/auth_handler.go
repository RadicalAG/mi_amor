package handler

import (
	"net/http"
	"radical/red_letter/internal/api_error"
	"radical/red_letter/internal/dto"
	"radical/red_letter/internal/middleware"
	"radical/red_letter/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service    service.AuthService
	middleware middleware.AuthMiddleware
}

func NewAuthHandler(service service.AuthService, middleware middleware.AuthMiddleware) *AuthHandler {
	return &AuthHandler{
		service:    service,
		middleware: middleware,
	}
}

func (t *AuthHandler) RegisterHandler(r *gin.Engine) *gin.Engine {
	r.POST("/auth/register", t.RegisterUser)
	r.POST("/auth/login", t.LoginUser)
	secured := r.Group("/auth/me").Use(middleware.NewAuthMiddleware().TokenAuthorization())
	{
		secured.GET("/ping", t.Ping)
	}
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

func (t *AuthHandler) LoginUser(c *gin.Context) {
	var req dto.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(api_error.NewApiError(http.StatusBadRequest, "invalid body"))
		return
	}

	token, err := t.service.LoginUser(c, req.Email, req.Password)
	if err != nil {
		c.Error(api_error.FromError(err))
		return
	}
	res := dto.LoginUserResponse{Token: token}

	c.JSON(http.StatusCreated, JsonSuccessFormater("Logged in Successfully", res))
}

func (t *AuthHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
