package respond

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func Error(w http.ResponseWriter, status int, message string, logger *slog.Logger) {

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	_, err := io.WriteString(w, message)
	if err != nil {
		logger.Error("response write error failed", slog.Any("err", err))
	}

}
func EncodeJson[T any](w http.ResponseWriter, status int, v T, logger *slog.Logger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		logger.Error("respond encode json failed", slog.Any("err", err))
	}
}
