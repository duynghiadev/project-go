package service

import (
	"context"

	"todo-service/pkg/mapper"
	"todo-service/pkg/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) UpdateTask(ctx context.Context, ID string, task model.Task) error {
	premID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	return s.Repo.UpdateTask(ctx, premID, mapper.MapToDto(task))
}
