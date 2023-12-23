package authservice

import (
	"fmt"
	"log/slog"

	"github.com/bomgar/basicwebapp/web/dto"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Logger *slog.Logger
}

func (s *AuthService) Register(registerRequest dto.RegisterRequest) error {

	_, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	return nil
}
