package web

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/bomgar/basicwebapp/db"
	"github.com/bomgar/basicwebapp/services"
	"github.com/bomgar/basicwebapp/web/controllers"
)

type RunSettings struct {
	ListenAddress string
	LogLevel      string
	DatabaseUrl   string
}

func Run(settings RunSettings) {
	logger := newLogger(settings.LogLevel)
	database := db.Connect(settings.DatabaseUrl, logger)
	err := db.Migrate(settings.DatabaseUrl, logger)
	if err != nil {
		logger.Error("Failed to migrate db", slog.Any("err", err))
		os.Exit(1)
	}

	services := services.Setup(logger, database)
	controllers := controllers.Setup(logger, services)
	r := SetupRoutes(controllers, logger)
	server := &http.Server{
		Addr:     settings.ListenAddress,
		Handler:  r,
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Error("Stop server: %v", err)
		os.Exit(1)
	}

}
