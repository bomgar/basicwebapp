package controllers

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
)

type Controllers struct {
	AuthController *AuthController
}

func Setup(logger *slog.Logger) *Controllers {
	validator := validator.New()
	return &Controllers{
		AuthController: &AuthController{
			logger:    logger,
			validator: validator,
		},
	}
}
