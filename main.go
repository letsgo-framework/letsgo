package main

import (
	"github.com/joho/godotenv"
	"github.com/letsgo-framework/letsgo/database"
	letslog "github.com/letsgo-framework/letsgo/log"
	"github.com/letsgo-framework/letsgo/routes"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {

	// Load env
	err := godotenv.Load()

	// Setup log writing
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
