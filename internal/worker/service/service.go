package service

import (
	"context"

	"github.com/morf1lo/notification-system/internal/user/pb"
	"github.com/morf1lo/notification-system/pkg/mq"
)

type Feed interface {
	ProcessFeeds(ctx context.Context)
}

type Service struct {
	Feed
}

func New(userService pb.UserClient, rabbitMQ *mq.MQConn) *Service {
	return &Service{
		Feed: NewFeedService(userService, rabbitMQ),
	}
}
