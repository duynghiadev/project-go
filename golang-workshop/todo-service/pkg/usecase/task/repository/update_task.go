package repository

import (
	"context"

	"todo-service/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) UpdateTask(ctx context.Context, ID primitive.ObjectID, task model.MongoTask) error {
	filter := bson.D{{"_id", ID}}

	update := bson.D{{"$set",
		bson.D{
			{"title", task.Title},
			{"completed", task.Completed},
			{"description", task.Description},
		}}}

	_, err := r.collection.UpdateOne(ctx, filter, update)

	return err
}
