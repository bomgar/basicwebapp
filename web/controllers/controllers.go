package controllers

import (
	"log/slog"

	"github.com/bomgar/basicwebapp/services"
	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	AuthController *AuthController
}

func Setup(logger *slog.Logger, services *services.Services) *Controllers {
	validator := validator.New()
	return &Controllers{
		AuthController: &AuthController{
			logger:      logger.With("controller", "AuthController"),
			validator:   validator,
			authService: services.AuthService,
		},
	}
}
