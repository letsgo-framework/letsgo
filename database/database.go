/*
|--------------------------------------------------------------------------
| Mongo Database Connection
|--------------------------------------------------------------------------
|
| We are using mongo-go-driver to connect to mongodb
| Connect is used to make connection
| TestConnect is used to make connection while running tests
|
*/

package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	letslog "github.com/letsgo-framework/letsgo/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"time"
)

// DB is pointer of Mongo Database
var DB *mongo.Database

// Client pointer of Mongo Client
var Client *mongo.Client

// Connect to database
func Connect() (*mongo.Client, *mongo.Database) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbHost := os.Getenv("DATABASE_HOST")
		dbPort := os.Getenv("DATABASE_PORT")
		dbURL = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	if err != nil {
		letslog.Fatal(err)
	}
	Client = client
	err = Client.Ping(context.Background(), readpref.Primary())
	if err == nil {
		letslog.Info("Connected to MongoDB!")
	} else {
		letslog.Error("Could not connect to MongoDB! Please check if mongo is running.")
	}
	DB = Client.Database(os.Getenv("DATABASE"))

	return Client, DB
}

// TestConnect to database while testing
func TestConnect() (*mongo.Client, *mongo.Database) {
	err := godotenv.Load("../.env.testing")
	if err != nil {
		letslog.Error("Error loading .env.testing file")
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbHost := os.Getenv("DATABASE_HOST")
		dbPort := os.Getenv("DATABASE_PORT")
		dbURL = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	Client = client
	DB = Client.Database(os.Getenv("DATABASE"))
	if err != nil {
		letslog.Fatal(err)
	}
	err = Client.Ping(context.Background(), readpref.Primary())
	if err == nil {
		letslog.Info("Connected to MongoDB for testing!")
	} else {
		letslog.Error("Could not connect to MongoDB!")
	}

	return Client, DB
}
