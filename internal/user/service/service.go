package service

import (
	"context"

	"github.com/morf1lo/notification-system/internal/user/model"
	"github.com/morf1lo/notification-system/internal/user/pb"
	"github.com/morf1lo/notification-system/internal/user/repository"
)

type User interface {
	Subscribe(ctx context.Context, sub *model.Subscriber) error
	FindAll(ctx context.Context) ([]*pb.Subscriber, error)
}

type Service struct {
	User
}

func New(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
	}
}
