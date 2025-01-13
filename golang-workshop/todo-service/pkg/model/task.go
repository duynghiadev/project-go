package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type MongoTask struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`
}
