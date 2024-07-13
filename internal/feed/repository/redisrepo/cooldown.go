package redisrepo

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CooldownRepo struct {
	rdb *redis.Client
}

func NewCooldownRepo(rdb *redis.Client) *CooldownRepo {
	return &CooldownRepo{rdb: rdb}
}

func (r *CooldownRepo) Create(ctx context.Context, key string, value int, expiration time.Duration) error {
	return r.rdb.Set(ctx, key, value, expiration).Err()
}

func (r *CooldownRepo) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.rdb.Get(ctx, key)
}

func (r *CooldownRepo) Increment(ctx context.Context, key string) error {
	_, err := r.rdb.Incr(ctx, key).Result()
	return err
}

func (r *CooldownRepo) SetExpiration(ctx context.Context, key string, expiration time.Duration) error {
	return r.rdb.Expire(ctx, key, expiration).Err()
}
