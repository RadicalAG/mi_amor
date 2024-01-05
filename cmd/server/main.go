package main

import (
	"context"
	"log"
	"radical/red_letter/config"
	"radical/red_letter/internal/generator"
	"radical/red_letter/internal/handler"
	"radical/red_letter/internal/middleware"
	"radical/red_letter/internal/model"
	"radical/red_letter/internal/repository"
	"radical/red_letter/internal/server"
	"radical/red_letter/internal/service"
	"radical/red_letter/internal/utils"

	"radical/red_letter/internal/db"
)

func main() {
	ctx := context.Background()
	configuration := config.Init()
	mdb := db.NewMongoDB(configuration.DB.Username, configuration.DB.Password, configuration.DB.Uri)
	client, cleanup := mdb.Connect(ctx)
	defer cleanup()
	s := server.NewHttpServer()

	tc := generator.NewTokenClaim(configuration.JWT.SecretKey)

	am := middleware.NewAuthMiddleware(tc)
	errHandler := middleware.NewErrorMiddleware()
	s.AddMiddleware(errHandler)

	v := utils.NewValidator()

	er := repository.NewEventRepository(client, db.DBName, model.EventCollectionName)
	ur := repository.NewUserRepository(client, db.DBName, model.UserCollectionName)

	es := service.NewEventService(er)
	as := service.NewAuthService(ur, v, tc)

	th := handler.NewTestHandler()
	eh := handler.NewEventHandler(es)
	ah := handler.NewAuthHandler(as, am, tc)
	s.AddHandler(th, eh, ah)

	s.Serve()
	log.Println("hey there!")
}
