package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) Repository {
	collection := client.Database("todo_db").Collection("tasks")

	return &MongoRepository{
		client:     client,
		collection: collection,
	}
}
