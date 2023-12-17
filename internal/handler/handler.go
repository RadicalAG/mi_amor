package handler

import "github.com/gin-gonic/gin"

type Handler interface {
	RegisterHandler(r *gin.Engine) *gin.Engine
}

func JsonSuccessFormater(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"data":    data,
	}
}
