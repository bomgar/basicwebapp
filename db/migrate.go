package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"os"
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

func Migrate(db *sql.DB, logger *slog.Logger) {
	goose.SetBaseFS(embedMigrations)
	goose.SetLogger(GooseLoggerAdapter{logger: logger})

	if err := goose.SetDialect("postgres"); err != nil {
		logger.Error("Failed to set dialect.", slog.Any("err", err))
		os.Exit(1)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		logger.Error("Failed to apply db migrations.", slog.Any("err", err))
		os.Exit(1)
	}

}
