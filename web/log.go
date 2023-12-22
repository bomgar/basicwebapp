package web

import (
	"log/slog"
	"net/http"
	"os"
	"time"
	"github.com/go-chi/chi/v5/middleware"
)

func newLogger() *slog.Logger {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    return logger

}

func slogMiddleware(log *slog.Logger) func(next http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

            defer func() {
                log.Info("Request",
                    slog.String("method", r.Method),
                    slog.String("url", r.URL.String()),
                    slog.String("proto", r.Proto),
                    slog.Int64("duration", time.Since(start).Microseconds()),
                    slog.String("remote", r.RemoteAddr),
                    slog.Int("status", ww.Status()),
                    slog.Int("bytesWritten", ww.BytesWritten()),
                    slog.String("user-agent", r.UserAgent()),
                )
            }()

            next.ServeHTTP(ww, r)
        })
    }
}
