package controllers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/bomgar/basicwebapp/services/authservice"
	"github.com/bomgar/basicwebapp/web/dto"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	logger      *slog.Logger
	validator   *validator.Validate
	authService *authservice.AuthService
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	registerRequest := &dto.RegisterRequest{}

	err := json.NewDecoder(r.Body).Decode(&registerRequest)
	if err != nil {
		c.logger.Error("Failed to decode request", slog.Any("err", err))
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	validationErrors := c.validator.Struct(registerRequest)
	if validationErrors != nil {
		c.logger.Info("Validation failed.", slog.Any("err", validationErrors))
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(validationErrors.Error()))
		return
	}

	c.logger.Info("Received register request.", slog.String("username", registerRequest.Email))

    err = c.authService.Register(r.Context(), *registerRequest)
	if err != nil {
		c.logger.Error("Failed to register user", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}

func (c *AuthController) WhoAmI(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("don't know"))

	if err != nil {
		c.logger.Error("Failed to write response: %v", err)
	}
}
