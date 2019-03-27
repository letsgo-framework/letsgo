package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var DB *mongo.Database
var Client *mongo.Client

func Connect() (*mongo.Client, *mongo.Database) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbHost := os.Getenv("DATABASE_HOST")
		dbPort := os.Getenv("DATABASE_PORT")
		dbURL = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	DB = Client.Database(os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return Client, DB
}

func TestConnect() (*mongo.Client, *mongo.Database) {
	err := godotenv.Load("../.env.testing")
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbHost := os.Getenv("DATABASE_HOST")
		dbPort := os.Getenv("DATABASE_PORT")
		dbURL = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	DB = Client.Database(os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB for testing!")

	return Client, DB
}