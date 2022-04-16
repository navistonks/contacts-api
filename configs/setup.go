package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	// Create mongo client with database URI
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectDB()

// Get database collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("contacts-app").Collection(collectionName)
	return collection
}
