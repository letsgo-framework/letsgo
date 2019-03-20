package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gitlab.com/letsgo/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var db *mongo.Database

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect DB
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbHost := os.Getenv("DATABASE_HOST")
		dbPort := os.Getenv("DATABASE_PORT")
		dbURL = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	db = client.Database(os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	srv := routes.PaveRoutes()
	srv.Run(os.Getenv("PORT"))

}
