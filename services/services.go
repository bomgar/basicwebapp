package services

import (
	"log/slog"

	"github.com/bomgar/basicwebapp/services/authservice"
)

type Services struct {
	AuthService *authservice.AuthService
}

func Setup(logger *slog.Logger) *Services {
	return &Services{
		AuthService: &authservice.AuthService{
			Logger: logger.With("service", "AuthService"),
		},
	}
}
