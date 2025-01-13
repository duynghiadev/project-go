package service

import (
	"context"

	"todo-service/pkg/mapper"
	"todo-service/pkg/model"
)

func (s *Service) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	tasks, err := s.Repo.GetAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	return mapTo(tasks), nil
}

func mapTo(list []model.MongoTask) []model.Task {
	mappedTasks := make([]model.Task, 0)
	for _, task := range list {
		mappedTasks = append(mappedTasks, mapper.MapToModel(task))
	}

	return mappedTasks
}
