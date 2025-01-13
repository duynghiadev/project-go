package service

import (
	"context"

	"todo-service/pkg/model"
	"todo-service/pkg/usecase/task/repository"
)

type TodoService interface {
	AddTask(ctx context.Context, task model.Task) error
	UpdateTask(ctx context.Context, ID string, task model.Task) error
	CompleteTask(ctx context.Context, ID string) error
	GetAllTasks(ctx context.Context) ([]model.Task, error)
}

type Service struct {
	Repo repository.Repository
}

func NewService(repo repository.Repository) TodoService {
	return &Service{Repo: repo}
}
