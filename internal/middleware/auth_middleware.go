package middleware

import (
	"radical/red_letter/internal/api_error"
	"radical/red_letter/internal/generator"
	"radical/red_letter/internal/internal_error"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	tokenGenerator generator.TokenGenerator
}

func NewAuthMiddleware(tokenGenerator generator.TokenGenerator) *authMiddleware {
	return &authMiddleware{
		tokenGenerator: tokenGenerator,
	}
}

type AuthMiddleware interface {
	TokenAuthorization() gin.HandlerFunc
}

func (a *authMiddleware) TokenAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Error(api_error.FromError(internal_error.Unauthorized("unauthorized")))
			c.Abort()
			return
		}

		data, err := a.tokenGenerator.ValidateAndDecodeToken(tokenString)
		if err != nil {
			c.Error(api_error.FromError(internal_error.Unauthorized("unauthorized")))
			c.Abort()
			return
		}

		c.Set("tokenClaims", data)
		c.Next()
	}
}
