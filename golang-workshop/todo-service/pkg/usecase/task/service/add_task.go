package service

import (
	"context"

	"todo-service/pkg/mapper"
	"todo-service/pkg/model"
)

func (s *Service) AddTask(ctx context.Context, task model.Task) error {
	return s.Repo.AddTask(ctx, mapper.MapToDto(task))
}
