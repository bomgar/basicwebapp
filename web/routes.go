package web

import (
	"github.com/bomgar/basicwebapp/web/controllers"
	"github.com/go-chi/chi/v5"
)

func setupRoutes(r *chi.Mux, c *controllers.Controllers) {
	r.Get("/whoami", c.AuthController.WhoAmI)
}
