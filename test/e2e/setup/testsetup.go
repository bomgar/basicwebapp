package setup

import (
	"log/slog"
	"net/http/httptest"
	"os"

	"github.com/bomgar/basicwebapp/db"
	"github.com/bomgar/basicwebapp/services"
	"github.com/bomgar/basicwebapp/web"
	"github.com/bomgar/basicwebapp/web/controllers"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TestSetup struct {
	Server      *httptest.Server
	DB          *pgxpool.Pool
	Services    *services.Services
	Controllers *controllers.Controllers
}

func (ts TestSetup) Close() {
	ts.Server.Close()
	ts.DB.Close()
}

func Setup() TestSetup {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	databaseUrl := "postgres://fkbr:fkbr@localhost:5432/fkbr"
	database := db.Connect(databaseUrl, logger)
	db.Migrate(databaseUrl, logger)

	services := services.Setup(logger, database)
	controllers := controllers.Setup(logger, services)
	ts := httptest.NewTLSServer(web.SetupRoutes(controllers, logger))
	return TestSetup{
		Server:      ts,
		DB:          database,
		Services:    services,
		Controllers: controllers,
	}
}
