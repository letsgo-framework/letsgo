package main

import (
	"github.com/joho/godotenv"
	"github.com/letsgo-framework/letsgo/database"
	"github.com/letsgo-framework/letsgo/routes"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

func main() {

	// Configure Logging
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./log/letsgo.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})


	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("env loaded")
	}

	database.Connect()

	srv := routes.PaveRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	srv.Run(port)

}
