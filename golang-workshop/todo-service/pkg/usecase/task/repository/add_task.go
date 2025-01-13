package repository

import (
	"context"

	"todo-service/pkg/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) AddTask(ctx context.Context, task model.MongoTask) error {
	task.ID = primitive.NewObjectID()

	_, err := r.collection.InsertOne(ctx, task)

	return err
}
