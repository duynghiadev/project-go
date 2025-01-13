package repository

import (
	"context"

	"todo-service/pkg/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	AddTask(ctx context.Context, task model.MongoTask) error
	UpdateTask(ctx context.Context, ID primitive.ObjectID, task model.MongoTask) error
	CompleteTask(ctx context.Context, ID string) error
	GetAllTasks(ctx context.Context) ([]model.MongoTask, error)
}
