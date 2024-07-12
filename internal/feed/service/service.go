package service

import (
	"context"

	"github.com/morf1lo/notification-system/internal/feed/model"
	"github.com/morf1lo/notification-system/internal/feed/repository"
	"github.com/morf1lo/notification-system/pkg/mq"
)

type Feed interface {
	Publish(ctx context.Context, article *model.Article) error
	FindByID(ctx context.Context, id int64) (*model.Article, error)
}

type Service struct {
	Feed
}

func New(repo *repository.Repository, rabbitMQ *mq.MQConn) *Service {
	return &Service{
		Feed: NewFeedService(repo, rabbitMQ),
	}
}
