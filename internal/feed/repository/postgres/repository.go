package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/morf1lo/notification-system/internal/feed/model"
)

type Article interface {
	Create(ctx context.Context, article *model.Article) error
	FindByID(ctx context.Context, id int64) (*model.Article, error)
}

type PostgresRepo struct {
	Article
}

func New(db *pgx.Conn) *PostgresRepo {
	return &PostgresRepo{
		Article: NewArticleRepo(db),
	}
}
