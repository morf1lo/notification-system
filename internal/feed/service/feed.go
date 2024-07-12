package service

import (
	"encoding/json"

	"github.com/morf1lo/notification-system/internal/feed/model"
	"github.com/morf1lo/notification-system/pkg/mq"
)

type FeedService struct {
	rabbitMQ *mq.MQConn
}

func NewFeedService(rabbitMQ *mq.MQConn) *FeedService {
	return &FeedService{rabbitMQ: rabbitMQ}
}

func (s *FeedService) Publish(article *model.Article) error {
	articleJSON, err := json.Marshal(article)
	if err != nil {
		return err
	}

	return s.rabbitMQ.Publish(articleEmailNotificationMQ, articleJSON)
}
