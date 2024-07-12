package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/morf1lo/notification-system/internal/user/config"
)

func Connect(ctx context.Context, cfg *config.DBConfig) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	db, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
