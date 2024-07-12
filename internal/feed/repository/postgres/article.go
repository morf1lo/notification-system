package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/morf1lo/notification-system/internal/feed/model"
)

type ArticleRepo struct {
	db *pgx.Conn
}

func NewArticleRepo(db *pgx.Conn) *ArticleRepo {
	return &ArticleRepo{db: db}
}

func (r *ArticleRepo) Create(ctx context.Context, article *model.Article) error {
	_, err := r.db.Exec(ctx, "insert into articles(title, body, author) values($1, $2, $3)", article.Title, article.Body, article.Author)
	return err
}

func (r *ArticleRepo) FindByID(ctx context.Context, id int64) (*model.Article, error) {
	var article model.Article
	if err := r.db.QueryRow(ctx, "select a.id, a.title, a.body, a.author, a.date_added from articles a where a.id = $1", id).Scan(&article.ID, &article.Title, &article.Body, &article.Author, &article.DateAdded); err != nil {
		return nil, err
	}

	return &article, nil
}
