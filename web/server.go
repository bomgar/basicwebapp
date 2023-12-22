package web

import (
	"log/slog"
	"net/http"

	"github.com/bomgar/basicwebapp/web/controllers"
)

type RunSettings struct {
	ListenAddress string
	LogLevel      string
}

func Run(settings RunSettings) {
	logger := newLogger(settings.LogLevel)

	controllers := controllers.Setup()
	r := SetupRoutes(controllers, logger)
	server := &http.Server{
		Addr:     settings.ListenAddress,
		Handler:  r,
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	server.ListenAndServe()
}
