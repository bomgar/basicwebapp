package authservice

import (
	"fmt"
	"log/slog"

	"github.com/bomgar/basicwebapp/db/q"
	"github.com/bomgar/basicwebapp/web/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type AuthService struct {
	Logger *slog.Logger
	DB     *pgxpool.Pool
}

func (s *AuthService) Register(ctx context.Context, registerRequest dto.RegisterRequest) error {

	pwHash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	err = s.DB.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
		queries := q.New(conn)
		queries.InsertUser(ctx, q.InsertUserParams{
			Email:          registerRequest.Email,
			HashedPassword: string(pwHash),
		})
		return nil
	})

	if err != nil {
		return fmt.Errorf("Could not use database: %w", err)
	}

	return nil
}
