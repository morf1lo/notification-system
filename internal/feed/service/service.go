package service

import (
	"github.com/morf1lo/notification-system/internal/feed/model"
	"github.com/morf1lo/notification-system/pkg/mq"
)

type Feed interface {
	Publish(article *model.Article) error
}

type Service struct {
	Feed
}

func New(rabbitMQ *mq.MQConn) *Service {
	return &Service{
		Feed: NewFeedService(rabbitMQ),
	}
}
