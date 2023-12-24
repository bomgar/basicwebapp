package authcontroller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/bomgar/basicwebapp/services/authservice"
	"github.com/bomgar/basicwebapp/web/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/securecookie"
)

type AuthController struct {
	logger       *slog.Logger
	validator    *validator.Validate
	authService  *authservice.AuthService
	secureCookie *securecookie.SecureCookie
}

type Session struct {
	expires time.Time
	userId  int32
}

const SessionName = "fkbr-session"
const CookieName = "fkbr-cookie"

func New(logger *slog.Logger, validator *validator.Validate, authService *authservice.AuthService) *AuthController {

	hashKey := securecookie.GenerateRandomKey(64)
	var sc = securecookie.New(hashKey, nil)
	return &AuthController{
		logger:       logger.With("controller", "AuthController"),
		validator:    validator,
		authService:  authService,
		secureCookie: sc,
	}
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

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest := &dto.LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		c.logger.Warn("Failed to decode request", slog.Any("err", err))
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	userId, err := c.authService.Login(r.Context(), *loginRequest)
	if err != nil {
		c.logger.Warn("Failed to login", slog.Any("err", err))
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Login failed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = c.setSession(userId, w)
	if err != nil {
		c.logger.Error("Failed to write session", slog.Any("err", err))
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Login failed"))
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dto.LoginResponse{UserId: userId})
	if err != nil {
		c.logger.Error("Failed to write response: %v", err)
	}

}

func (c *AuthController) setSession(userId int32, w http.ResponseWriter) error {
	expires := time.Now().Add(30 * 24 * time.Hour)

	session := Session{
		userId:  userId,
		expires: expires,
	}
	jsonSession, err := json.Marshal(session)
	if err != nil {
		return err
	}
	if encoded, err := c.secureCookie.Encode(SessionName, jsonSession); err == nil {
		cookie := &http.Cookie{
			Name:     CookieName,
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
			Expires:  expires,
		}
		http.SetCookie(w, cookie)
	}
	return nil
}

func (c *AuthController) WhoAmI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("{}"))

	if err != nil {
		c.logger.Error("Failed to write response: %v", err)
	}
}
