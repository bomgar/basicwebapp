package web

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func newLogger(level string) *slog.Logger {
	var l slog.Level
	switch strings.ToUpper(level) {
	case "DEBUG":
		l = slog.LevelDebug
	case "INFO":
		l = slog.LevelInfo
	case "WARN":
		l = slog.LevelWarn
	case "ERROR":
		l = slog.LevelError
	default:
		l = slog.LevelInfo
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: l,
	}))
	return logger

}

func slogMiddleware(log *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			defer func() {
				log.Info(fmt.Sprintf("ACCESS %s %s", r.Method, r.URL),
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
