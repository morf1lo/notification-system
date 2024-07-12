package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/morf1lo/notification-system/internal/user/repository/postgres"
)

type Repository struct {
	Postgres *postgres.PostgresRepo
}

func New(db *pgx.Conn) *Repository {
	return &Repository{
		Postgres: postgres.NewPostgresRepo(db),
	}
}
