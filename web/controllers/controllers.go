package controllers

import (
	"log/slog"

	"github.com/bomgar/basicwebapp/services"
	"github.com/bomgar/basicwebapp/web/controllers/authcontroller"
	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	AuthController *authcontroller.AuthController
}

func Setup(logger *slog.Logger, services *services.Services) *Controllers {
	validator := validator.New()
	return &Controllers{
		AuthController: authcontroller.New(
			logger.With("controller", "AuthController"),
			validator,
			services.AuthService,
		),
	}
}
