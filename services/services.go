package services

import (
	"log/slog"

	"github.com/bomgar/basicwebapp/services/authservice"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Services struct {
	AuthService *authservice.AuthService
}

func Setup(logger *slog.Logger, DB *pgxpool.Pool) *Services {
	return &Services{
		AuthService: &authservice.AuthService{
			Logger: logger.With("service", "AuthService"),
			DB:     DB,
		},
	}
}
