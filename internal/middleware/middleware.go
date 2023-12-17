package middleware

import "github.com/gin-gonic/gin"

type Middleware interface {
	RegisterMiddleware(r *gin.Engine) *gin.Engine
}
