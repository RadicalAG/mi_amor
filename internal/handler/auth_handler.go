package handler

import (
	"net/http"
	"radical/red_letter/internal/api_error"
	"radical/red_letter/internal/dto"
	"radical/red_letter/internal/generator"
	"radical/red_letter/internal/internal_error"
	"radical/red_letter/internal/middleware"
	"radical/red_letter/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service        service.AuthService
	authmiddleware middleware.AuthMiddleware
	tokenClaim     generator.TokenClaim
}

func NewAuthHandler(service service.AuthService, authmiddleware middleware.AuthMiddleware, tokenClaim generator.TokenClaim) *AuthHandler {
	return &AuthHandler{
		service:        service,
		authmiddleware: authmiddleware,
		tokenClaim:     tokenClaim,
	}
}

func (t *AuthHandler) RegisterHandler(r *gin.Engine) *gin.Engine {
	r.POST("/auth/register", t.RegisterUser)
	r.POST("/auth/login", t.LoginUser)
	r.Use(t.authmiddleware.TokenAuthorization()).GET("/auth/me", t.Me)
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

func (t *AuthHandler) Me(c *gin.Context) {
	res, exists := c.Get("tokenClaims")
	if !exists {
		c.Error(api_error.FromError(internal_error.BadRequestError("unauthorized")))
		return
	}

	c.JSON(http.StatusOK, JsonSuccessFormater("Successfully Get User", res))
}
