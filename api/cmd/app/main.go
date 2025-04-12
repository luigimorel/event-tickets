package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/morelmiles/go-events/internals/routes"
	"github.com/morelmiles/go-events/pkg/database"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := routes.SetupRoutes(); err != nil {
		log.Fatal("Error setting up routes: ", err)
	}
}
