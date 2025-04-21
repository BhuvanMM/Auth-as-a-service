package database

import (
	"context"
	"log"
	"session-service/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Fatal("Mongo client creation failed:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Mongo connection failed:", err)
	}

	MongoClient = client
	log.Println("Connected to MongoDB")
}

func GetCollection() *mongo.Collection {
	return MongoClient.Database(config.DatabaseName).Collection(config.SessionCol)
}
