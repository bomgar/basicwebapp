package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Run(port int) {
	r := chi.NewRouter()
	logger := newLogger()

	r.Use(slogMiddleware(logger))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)

}
