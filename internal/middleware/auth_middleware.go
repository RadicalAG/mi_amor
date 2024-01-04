package middleware

import (
	"radical/red_letter/internal/generator"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct{}

func NewAuthMiddleware() *authMiddleware {
	return &authMiddleware{}
}

type AuthMiddleware interface {
	TokenAuthorization() gin.HandlerFunc
}

func (a *authMiddleware) TokenAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := generator.NewTokenGenerator().ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
