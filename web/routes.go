package web

import (
	"log/slog"

	"github.com/bomgar/basicwebapp/web/controllers"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(c *controllers.Controllers, logger *slog.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Use(slogMiddleware(logger))
	r.Get("/whoami", c.AuthController.WhoAmI)

	return r
}
