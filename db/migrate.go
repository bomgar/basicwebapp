package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type GooseLoggerAdapter struct {
	logger *slog.Logger
}

func (l GooseLoggerAdapter) Fatalf(format string, v ...interface{}) {
	l.logger.Error(strings.TrimSpace(fmt.Sprintf(format, v...)))

}

func (l GooseLoggerAdapter) Printf(format string, v ...interface{}) {
	l.logger.Info(strings.TrimSpace(fmt.Sprintf(format, v...)))
}

func Migrate(databaseUrl string, logger *slog.Logger) error {
	goose.SetBaseFS(embedMigrations)
	goose.SetLogger(GooseLoggerAdapter{logger: logger})
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return fmt.Errorf("Connect to database failed: %w", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("db migration failed: %w", err)
	}

	return nil

}
