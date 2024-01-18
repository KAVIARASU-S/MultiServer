package config

import (
	"context"
	"log"
	"server/constants"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectDatabase (ctx context.Context) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	
	log.Println("Connecting to MongoDb")

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(constants.ConnectionString))

	if err != nil {
		log.Println("Error connecting to MongoDB: ", err)
		return nil,err
	} else {
		log.Println("Conneted to MongoDb successfully")
	}

	return mongoClient, nil
}

func GetCollection (mongoClient *mongo.Client, DB_Name string, CollectionName string) (*mongo.Collection) {
	collection := mongoClient.Database(DB_Name).Collection(CollectionName)
	return collection
}