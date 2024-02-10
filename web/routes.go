package web

import (
	"log/slog"
	"net/http"

	"github.com/bomgar/basicwebapp/web/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoutes(c *controllers.Controllers, logger *slog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(slogMiddleware(logger))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Post("/register", c.AuthController.Register)
	r.Post("/login", c.AuthController.Login)
	r.With(c.AuthController.AuthenticatedMiddleware).Get("/whoami", c.AuthController.WhoAmI)

	return r
}
