package controllers

import (
	"log/slog"
	"net/http"
)

type AuthController struct {
	logger *slog.Logger
}

func (c *AuthController) WhoAmI(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("don't know"))

	if err != nil {
		c.logger.Error("Failed to write response: %v", err)
	}
}
