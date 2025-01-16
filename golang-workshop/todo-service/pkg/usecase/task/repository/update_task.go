package repository

import (
	"context"

	"todo-service/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoRepository) UpdateTask(ctx context.Context, ID primitive.ObjectID, task model.MongoTask) error {
	// Filter để tìm tài liệu cần cập nhật
	filter := bson.M{"_id": ID}

	// Update với các trường cần cập nhật
	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"completed":   task.Completed,
			"description": task.Description,
		},
	}

	// Gọi UpdateOne để cập nhật tài liệu
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}
