package setup

import (
	"database/sql"
	"log/slog"
	"net/http/httptest"
	"os"

	"github.com/bomgar/basicwebapp/db"
	"github.com/bomgar/basicwebapp/web"
	"github.com/bomgar/basicwebapp/web/controllers"
)

type TestSetup struct {
	Server *httptest.Server
	DB     *sql.DB
}

func (ts TestSetup) Close() {
	ts.Server.Close()
	ts.DB.Close()
}

func TestServer() TestSetup {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	database := db.Connect("postgres://fkbr:fkbr@localhost:5432/fkbr", logger)
	db.Migrate(database, logger)

	controllers := controllers.Setup(logger)
	ts := httptest.NewTLSServer(web.SetupRoutes(controllers, logger))
	return TestSetup{
		Server: ts,
		DB:     database,
	}
}
