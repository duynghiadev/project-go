package db

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var mongoOnce sync.Once
var clientInstanceError error

type Collection string

const (
	ProductsCollection Collection = "products"
)

var Database string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or failed to load")
	}

	Database = os.Getenv("DATABASE_NAME")
	if Database == "" {
		log.Fatal("DATABASE_NAME is not set in the environment variables")
	}
}

// GetMongoClient initializes a MongoDB client and ensures a single instance (singleton pattern)
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		mongoURI := os.Getenv("MONGO_URI")
		if mongoURI == "" {
			log.Fatal("MONGO_URI is not set in the environment variables")
		}

		clientOptions := options.Client().ApplyURI(mongoURI)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Printf("Failed to connect to MongoDB: %v", err)
			clientInstanceError = err
			return
		}

		// Kiểm tra kết nối
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Printf("Failed to ping MongoDB: %v", err)
			clientInstanceError = err
			return
		}

		clientInstance = client
		clientInstanceError = err

		log.Printf("Successfully connected to MongoDB Database: %s", Database)
	})

	return clientInstance, clientInstanceError
}
