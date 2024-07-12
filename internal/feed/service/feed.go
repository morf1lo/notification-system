package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/morf1lo/notification-system/internal/feed/model"
	"github.com/morf1lo/notification-system/internal/feed/repository"
	"github.com/morf1lo/notification-system/pkg/mq"
	"github.com/redis/go-redis/v9"
)

type FeedService struct {
	repo *repository.Repository
	rabbitMQ *mq.MQConn
}

func NewFeedService(repo *repository.Repository, rabbitMQ *mq.MQConn) *FeedService {
	return &FeedService{
		repo: repo,
		rabbitMQ: rabbitMQ,
	}
}

func (s *FeedService) Publish(ctx context.Context, article *model.Article) error {
	if err := s.repo.Postgres.Article.Create(ctx, article); err != nil {
		return err
	}

	articleJSON, err := json.Marshal(article)
	if err != nil {
		return err
	}

	return s.rabbitMQ.Publish(articleEmailNotificationMQ, articleJSON)
}

func (s *FeedService) FindByID(ctx context.Context, id int64) (*model.Article, error) {
	article, err := s.repo.Redis.Article.Find(ctx, ArticlePrefix(id))
	if err == nil {
		return article, nil
	}

	if err != redis.Nil {
		return nil, err
	}

	articleDB, err := s.repo.Postgres.Article.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	articleJSON, err := json.Marshal(articleDB)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Redis.Article.Create(ctx, ArticlePrefix(id), articleJSON, time.Hour * 24); err != nil {
		return nil, err
	}

	return articleDB, nil
}
