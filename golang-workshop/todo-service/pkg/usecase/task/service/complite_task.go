package service

import "context"

func (s *Service) CompleteTask(ctx context.Context, ID string) error {
	return s.Repo.CompleteTask(ctx, ID)
}
