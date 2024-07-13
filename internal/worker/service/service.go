package service

import (
	"context"

	"github.com/morf1lo/notification-system/internal/worker/model"
	"github.com/morf1lo/notification-system/internal/user/pb"
	"github.com/morf1lo/notification-system/pkg/mq"
)

type Feed interface {
	ProcessFeeds(ctx context.Context)
}

type Mailer interface {
	Send(to string, message *model.Article) error
}

type Service struct {
	Feed
}

func New(userService pb.UserClient, rabbitMQ *mq.MQConn) *Service {
	mailer := NewMailerService()

	return &Service{
		Feed: NewFeedService(userService, rabbitMQ, mailer),
	}
}
