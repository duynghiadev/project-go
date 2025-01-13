package repository

import (
	"context"

	"todo-service/pkg/model"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *MongoRepository) GetAllTasks(ctx context.Context) ([]model.MongoTask, error) {
	var tasks []model.MongoTask
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task model.MongoTask
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
