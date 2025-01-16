package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) CompleteTask(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	update := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "completed", Value: true}}}}
	_, err = r.collection.UpdateOne(ctx, filter, update)

	return err
}
