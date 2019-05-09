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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	Client = client
	err = Client.Ping(context.Background(), readpref.Primary())
	if err == nil {
		log.Println("Connected to MongoDB!")
	} else {
		log.Fatalln("Could not connect to MongoDB! Please check if mongo is running.")
	}
	DB = Client.Database(os.Getenv("DATABASE"))

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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	Client = client
	DB = Client.Database(os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	err = Client.Ping(context.Background(), readpref.Primary())
	if err == nil {
		fmt.Println("Connected to MongoDB for testing!")
	} else {
		fmt.Println("Could not connect to MongoDB!")
	}

	return Client, DB
}
