package redisrepo

import (
	"context"
	"time"

	"github.com/morf1lo/notification-system/internal/feed/model"
	"github.com/redis/go-redis/v9"
)

type Article interface {
	Create(ctx context.Context, key string, value []byte, expiration time.Duration) error
	Find(ctx context.Context, key string) (*model.Article, error)
}

type RedisRepo struct {
	Article
}

func New(rdb *redis.Client) *RedisRepo {
	return &RedisRepo{
		Article: NewArticleRepo(rdb),
	}
}
