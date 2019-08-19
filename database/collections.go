package database

import "go.mongodb.org/mongo-driver/mongo"

func UserCollection() *mongo.Collection {
	return DB.Collection("users")
}
