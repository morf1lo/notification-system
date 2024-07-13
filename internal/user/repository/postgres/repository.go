package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/morf1lo/notification-system/internal/user/model"
	"github.com/morf1lo/notification-system/internal/user/pb"
)

type Subscriber interface {
	Create(ctx context.Context, user *model.Subscriber) error
	FindByID(ctx context.Context, id int64) (*model.Subscriber, error)
	FindByEmail(ctx context.Context, email string) (*model.Subscriber, error)
	ExistsByEmail(ctx context.Context, email string) bool
	GetCountOfSubscribers(ctx context.Context) int
	FindAll(ctx context.Context) ([]*pb.Subscriber, error)
}

type PostgresRepo struct {
	Subscriber
}

func NewPostgresRepo(db *pgx.Conn) *PostgresRepo {
	return &PostgresRepo{
		Subscriber: NewSubscriberRepo(db),
	}
}
