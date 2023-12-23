package db

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(databaseUrl string, logger *slog.Logger) *sql.DB {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		logger.Error("Connect to database: %v", err)
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		logger.Error("Ping database: %v", err)
		os.Exit(1)
	}
	return db
}
