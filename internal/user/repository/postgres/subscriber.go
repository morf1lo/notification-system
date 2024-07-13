package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/morf1lo/notification-system/internal/user/model"
	"github.com/morf1lo/notification-system/internal/user/pb"
)

type SubscriberRepo struct {
	db *pgx.Conn
}

func NewSubscriberRepo(db *pgx.Conn) *SubscriberRepo {
	return &SubscriberRepo{db: db}
}

func (r *SubscriberRepo) Create(ctx context.Context, user *model.Subscriber) error {
	_, err := r.db.Exec(ctx, "insert into subscribers(email) values($1)", user.Email)
	return err
}

func (r *SubscriberRepo) FindByID(ctx context.Context, id int64) (*model.Subscriber, error) {
	var sub model.Subscriber
	if err := r.db.QueryRow(ctx, "select s.id, s.email, s.date_added from subscribers s where s.id = $1", id).Scan(&sub.ID, &sub.Email, &sub.DateAdded); err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *SubscriberRepo) FindByEmail(ctx context.Context, email string) (*model.Subscriber, error) {
	var sub model.Subscriber
	if err := r.db.QueryRow(ctx, "select s.id, s.email, s.date_added from subscribers s where s.email = $1", email).Scan(&sub.ID, &sub.Email, &sub.DateAdded); err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *SubscriberRepo) ExistsByEmail(ctx context.Context, email string) bool {
	var exists bool
	r.db.QueryRow(ctx, "select exists(select 1 from subscribers s where s.email = $1)", email).Scan(&exists)
	return exists
}

func (r *SubscriberRepo) GetCountOfSubscribers(ctx context.Context) int {
	var count int
	r.db.QueryRow(ctx, "select count(*) from subscribers").Scan(&count)
	return count
}

func (r *SubscriberRepo) FindAll(ctx context.Context) ([]*pb.Subscriber, error) {
	var count int64
	if err := r.db.QueryRow(ctx, "select count(s.id) from subscribers s").Scan(&count); err != nil {
		return nil, err
	}

	const batchSize = 100
	var subs []*pb.Subscriber
	var maxID int64

	for len(subs) < int(count) {
		rows, err := r.db.Query(ctx, "select s.id, s.email from subscribers s where s.id > $1 order by s.id limit $2", maxID, batchSize)
		if err != nil {
			return nil, err
		}

		var batch []*pb.Subscriber
		for rows.Next() {
			var sub pb.Subscriber
			if err := rows.Scan(&sub.Id, &sub.Email); err != nil {
				return nil, err
			}
			
			batch = append(batch, &sub)
		}

		if err := rows.Err(); err != nil {
			return nil, err
		}

		rows.Close()

		if len(batch) == 0 {
			break
		}

		subs = append(subs, batch...)
		maxID = subs[len(subs)-1].Id
	}

	if len(subs) == 0 {
		return nil, errNoSubs
	}

	return subs, nil
}
