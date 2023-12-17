package db

import (
	"log"
	"os"

	pg "github.com/go-pg/pg"
)

var DBName string = "red_letter"

func Connect() {
	opts := &pg.Options{
		User:     "postgres",
		Password: "Password123",
		Addr:     "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to database.\n")
		os.Exit(100)
	}
	log.Printf("Database connected successfully.\n")
	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Error while closing connection, reason: %v\n", closeErr)
		os.Exit(100)
	}
	log.Printf("Connection close successfully\n")
	return
}
