package db

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseUrl string, logger *slog.Logger) *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		logger.Error("Connect to database failed.", slog.Any("err", err))
		os.Exit(1)
	}
	err = db.Ping(context.Background())
	if err != nil {
		logger.Error("Ping database failed.", slog.Any("err", err))
		os.Exit(1)
	}
	return db
}
