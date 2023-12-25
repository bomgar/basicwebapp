package testsetup

import (
	"context"
	"log/slog"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bomgar/basicwebapp/db"
	"github.com/bomgar/basicwebapp/services"
	"github.com/bomgar/basicwebapp/web"
	"github.com/bomgar/basicwebapp/web/controllers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

type TestSetup struct {
	Server      *httptest.Server
	DB          *pgxpool.Pool
	Services    *services.Services
	Controllers *controllers.Controllers
}

func (ts *TestSetup) Close() {
	ts.Server.Close()
	ts.DB.Close()
}

func Setup(t *testing.T) *TestSetup {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	databaseUrl := "postgres://fkbr:fkbr@localhost:5432/fkbr"
	database := db.Connect(databaseUrl, logger)
	err := db.Migrate(databaseUrl, logger)

	require.Nil(t, err)

	err = cleanDb(database)
	require.Nil(t, err)

	services := services.Setup(logger, database)
	controllers := controllers.Setup(logger, services)
	ts := httptest.NewTLSServer(web.SetupRoutes(controllers, logger))
	return &TestSetup{
		Server:      ts,
		DB:          database,
		Services:    services,
		Controllers: controllers,
	}
}

func cleanDb(database *pgxpool.Pool) error {
	ctx := context.Background()
	return database.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
		deleteQueries := []string{
			"DELETE FROM users",
		}
		for _, query := range deleteQueries {
			_, err := conn.Exec(ctx, query)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
