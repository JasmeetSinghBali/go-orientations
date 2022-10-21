package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

/*create DBInstance & instantiate DBConnection*/
func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongoDB := os.Getenv("MONGODB_URL")

	dbConnection, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = dbConnection.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return dbConnection
}

var Client *mongo.Client = DBInstance()

/*access a collection from DB via dbConnectInstance*/
func OpenCollection(DBConnection *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = DBConnection.Database("cluster0").Collection(collectionName)
	return collection
}
