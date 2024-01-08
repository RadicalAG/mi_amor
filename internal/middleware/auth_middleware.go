package middleware

import (
	"radical/red_letter/internal/api_error"
	"radical/red_letter/internal/generator"
	"radical/red_letter/internal/internal_error"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	tokenClaim generator.TokenClaim
}

func NewAuthMiddleware(tokenClaim generator.TokenClaim) *authMiddleware {
	return &authMiddleware{
		tokenClaim: tokenClaim,
	}
}

type AuthMiddleware interface {
	TokenAuthorization() gin.HandlerFunc
}

func (a *authMiddleware) TokenAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Error(api_error.FromError(internal_error.BadRequestError("unauthorized")))
			return
		}

		data, err := a.tokenClaim.ValidateAndDecodeToken(tokenString)
		if err != nil {
			c.Error(api_error.FromError(internal_error.BadRequestError("unauthorized")))
			return
		}

		c.Set("tokenClaims", data)
		c.Next()
	}
}
