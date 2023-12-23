package web

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/bomgar/basicwebapp/db"
	"github.com/bomgar/basicwebapp/web/controllers"
)

type RunSettings struct {
	ListenAddress string
	LogLevel      string
	DatabaseUrl   string
}

func Run(settings RunSettings) {
	logger := newLogger(settings.LogLevel)
	_ = db.Connect(settings.DatabaseUrl, logger)

	controllers := controllers.Setup(logger)
	r := SetupRoutes(controllers, logger)
	server := &http.Server{
		Addr:     settings.ListenAddress,
		Handler:  r,
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Error("Stop server: %v", err)
		os.Exit(1)
	}

}
