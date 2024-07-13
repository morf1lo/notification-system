package redisrepo

import (
	"context"
	"time"

	"github.com/morf1lo/notification-system/internal/feed/model"
	"github.com/redis/go-redis/v9"
)

type Article interface {
	Create(ctx context.Context, key string, value []byte, expiration time.Duration) error
	Get(ctx context.Context, key string) (*model.Article, error)
}

type Cooldown interface {
	Create(ctx context.Context, key string, value int, expiration time.Duration) error
	Get(ctx context.Context, key string) *redis.StringCmd
	Increment(ctx context.Context, key string) error
	SetExpiration(ctx context.Context, key string, expiration time.Duration) error
}

type RedisRepo struct {
	Article
	Cooldown
}

func New(rdb *redis.Client) *RedisRepo {
	return &RedisRepo{
		Article: NewArticleRepo(rdb),
		Cooldown: NewCooldownRepo(rdb),
	}
}
