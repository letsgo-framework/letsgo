package main

import (
	"github.com/joho/godotenv"
	"github.com/letsgo-framework/letsgo/database"
	letslog "github.com/letsgo-framework/letsgo/log"
	"github.com/letsgo-framework/letsgo/routes"
	"os"
)

func main() {

	// Load env
	err := godotenv.Load()
	letslog.InitLogFuncs()
	if err != nil {
		letslog.Error("Error loading .env file")
	} else {
		letslog.Info("env loaded")
	}

	database.Connect()

	srv := routes.PaveRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	srv.Run(port)

}
