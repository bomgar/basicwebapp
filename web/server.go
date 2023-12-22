package web

import (
	"net/http"

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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(settings.ListenAddress, r)

}
