package setup

import (
	"log/slog"
	"net/http/httptest"
	"os"

	"github.com/bomgar/basicwebapp/web"
	"github.com/bomgar/basicwebapp/web/controllers"
)

func TestServer() *httptest.Server {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	controllers := controllers.Setup()
	ts := httptest.NewTLSServer(web.SetupRoutes(controllers, logger))
	return ts
}
