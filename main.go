package main

import (
	"github.com/joho/godotenv"
	"github.com/letsGo/database"
	"github.com/letsGo/routes"
	"log"
	"os"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	srv := routes.PaveRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	srv.Run(port)

}
