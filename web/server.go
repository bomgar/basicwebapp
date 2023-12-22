package web

import (
	"net/http"

	"github.com/bomgar/basicwebapp/web/controllers"
	"github.com/go-chi/chi/v5"
)

type RunSettings struct {
	ListenAddress string
	LogLevel      string
}

func Run(settings RunSettings) {
	r := chi.NewRouter()
	logger := newLogger(settings.LogLevel)

	r.Use(slogMiddleware(logger))

	controllers := controllers.Setup()

	setupRoutes(r, controllers)
	http.ListenAndServe(settings.ListenAddress, r)
}
