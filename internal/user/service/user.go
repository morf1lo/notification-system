package service

import (
	"context"

	"github.com/morf1lo/notification-system/internal/user/model"
	"github.com/morf1lo/notification-system/internal/user/pb"
	"github.com/morf1lo/notification-system/internal/user/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Subscribe(ctx context.Context, sub *model.Subscriber) error {
	if err := sub.Validate(); err != nil {
		return err
	}

	if s.repo.Postgres.Subscriber.ExistsByEmail(ctx, sub.Email) {
		return errUserIsAlreadyExists
	}

	err := s.repo.Postgres.Subscriber.Create(ctx, sub)
	return err
}

func (s *UserService) FindAll(ctx context.Context) ([]*pb.Subscriber, error) {
	subs, err := s.repo.Postgres.Subscriber.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return subs, nil
}
