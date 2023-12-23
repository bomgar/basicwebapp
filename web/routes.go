package web

import (
	"log/slog"

	"github.com/bomgar/basicwebapp/web/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(c *controllers.Controllers, logger *slog.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(slogMiddleware(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Get("/whoami", c.AuthController.WhoAmI)
	r.Post("/register", c.AuthController.Register)
	r.Post("/login", c.AuthController.Login)

	return r
}
