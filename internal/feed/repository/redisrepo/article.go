package redisrepo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/morf1lo/notification-system/internal/feed/model"
	"github.com/redis/go-redis/v9"
)

type ArticleRepo struct {
	rdb *redis.Client
}

func NewArticleRepo(rdb *redis.Client) *ArticleRepo {
	return &ArticleRepo{rdb: rdb}
}

func (r *ArticleRepo) Create(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	return r.rdb.Set(ctx, key, value, expiration).Err()
}

func (r *ArticleRepo) Find(ctx context.Context, key string) (*model.Article, error) {
	article, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var articleDB model.Article
	if err := json.Unmarshal([]byte(article), &articleDB); err != nil {
		return nil, err
	}

	return &articleDB, nil
}
