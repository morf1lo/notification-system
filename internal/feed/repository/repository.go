package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/morf1lo/notification-system/internal/feed/repository/postgres"
	"github.com/morf1lo/notification-system/internal/feed/repository/redisrepo"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	Postgres *postgres.PostgresRepo
	Redis    *redisrepo.RedisRepo
}

func New(db *pgx.Conn, rdb *redis.Client) *Repository {
	return &Repository{
		Postgres: postgres.New(db),
		Redis: redisrepo.New(rdb),
	}
}
