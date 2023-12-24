package authcontroller

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"slices"
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
	Expires time.Time
	UserId  int32
}

type SessionContextKey string

var SessionKey SessionContextKey = "session"

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
		UserId:  userId,
		Expires: expires,
	}
	encoded, err := c.secureCookie.Encode(SessionName, session)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:     CookieName,
		Value:    encoded,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		Expires:  expires,
	}
	http.SetCookie(w, cookie)
	return nil
}

func (c *AuthController) WhoAmI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(dto.WhoAmIResponse{UserId: r.Context().Value(SessionKey).(Session).UserId})
	if err != nil {
		c.logger.Error("Failed to write response: %v", err)
	}
}

func (c *AuthController) AuthenticatedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieIndex := slices.IndexFunc(r.Cookies(), func(cookie *http.Cookie) bool {
			return cookie.Name == CookieName
		})
		if cookieIndex == -1 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		cookie := r.Cookies()[cookieIndex]
		session := Session{}
		err := c.secureCookie.Decode(SessionName, cookie.Value, &session)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if session.Expires.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), SessionKey, session)))
	})
}
