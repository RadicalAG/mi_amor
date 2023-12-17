package main

import (
	"context"
	"log"
	"radical/red_letter/config"
	"radical/red_letter/internal/handler"
	"radical/red_letter/internal/middleware"
	"radical/red_letter/internal/model"
	"radical/red_letter/internal/repository"
	"radical/red_letter/internal/server"
	"radical/red_letter/internal/service"

	"radical/red_letter/internal/db"
)

func main() {
	ctx := context.Background()
	configuration := config.Init()
	mdb := db.NewMongoDB(configuration.DB.Username, configuration.DB.Password, configuration.DB.Uri)
	client, cleanup := mdb.Connect(ctx)
	defer cleanup()
	s := server.NewHttpServer()

	errHandler := middleware.NewErrorMiddleware()
	s.AddMiddleware(errHandler)

	er := repository.NewEventRepository(client, db.DBName, model.EventCollectionName)

	es := service.NewEventService(er)

	th := handler.NewTestHandler()
	eh := handler.NewEventHandler(es)
	s.AddHandler(th, eh)

	s.Serve()
	log.Println("hey there!")
}
