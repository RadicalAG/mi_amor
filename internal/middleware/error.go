package middleware

import (
	"net/http"
	"radical/red_letter/internal/api_error"

	"github.com/gin-gonic/gin"
)

type errorMiddleware struct {
}

func NewErrorMiddleware() *errorMiddleware {
	return &errorMiddleware{}
}

func (e *errorMiddleware) HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Errors.Last() == nil {
			return
		}
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case api_error.APIError:
				c.AbortWithStatusJSON(e.Status, map[string]string{"message": e.Message})
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": "Service Unavailable"})
				return
			}
		}
	}
}

func (e *errorMiddleware) RegisterMiddleware(r *gin.Engine) *gin.Engine {
	r.Use(e.HandleError())
	return r
}
