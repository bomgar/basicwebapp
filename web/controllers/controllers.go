package controllers

import "log/slog"

type Controllers struct {
	AuthController *AuthController
}

func Setup(logger *slog.Logger) *Controllers {
	return &Controllers{
		AuthController: &AuthController{
			logger: logger,
		},
	}
}
