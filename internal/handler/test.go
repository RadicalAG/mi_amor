package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type testHandler struct {
}

func NewTestHandler() *testHandler {
	return &testHandler{}
}

func (t *testHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (t *testHandler) RegisterHandler(r *gin.Engine) *gin.Engine {
	r.GET("/ping", t.Ping)
	return r
}
