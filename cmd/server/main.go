package main

import (
	"log"
	"radical/red_letter/internal/handler"
	"radical/red_letter/internal/server"

	db "radical/red_letter/cmd/server/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	s := server.NewHttpServer(r)
	th := handler.NewTestHandler()
	s.AddHandler(th)
	s.Serve()
	log.Println("hey there!")
	db.Connect()
}
