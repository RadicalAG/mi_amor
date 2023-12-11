package main

import (
	"radical/red_letter/internal/handler"
	"radical/red_letter/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	s := server.NewHttpServer(r)
	th := handler.NewTestHandler()
	s.AddHandler(th)
	s.Serve()
}
