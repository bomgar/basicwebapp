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

	return s.DB.AcquireFunc(ctx, func(conn *pgxpool.Conn) error {
		queries := q.New(conn)
		_, err = queries.InsertUser(ctx, q.InsertUserParams{
			Email:          registerRequest.Email,
			HashedPassword: string(pwHash),
		})
		if err != nil {
			return fmt.Errorf("Could not insert user: %w", err)
		}
		return nil
	})
}

func (s *AuthService) Login(ctx context.Context, loginRequest dto.LoginRequest) (int32, error) {
	conn, err := s.DB.Acquire(ctx)
	if err != nil {
		return -1, err
	}
	defer conn.Release()
	queries := q.New(conn)

	row, err := queries.SelectPasswordHashByUserEmail(ctx, loginRequest.Email)
	if err != nil {
		return -1, fmt.Errorf("Could not retrieve password hash: %w", err)
	}

	passwordHash := row.HashedPassword
	userId := row.ID

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(loginRequest.Password))
	if err != nil {
		return -1, fmt.Errorf("Password check failed: %w", err)
	}

	return userId, nil
}
