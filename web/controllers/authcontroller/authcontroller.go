package authcontroller

import (
	"context"
	"log/slog"
	"net/http"
	"slices"
	"time"

	"github.com/bomgar/basicwebapp/services/authservice"
	"github.com/bomgar/basicwebapp/web/dto"
	"github.com/bomgar/basicwebapp/web/receive"
	"github.com/bomgar/basicwebapp/web/respond"
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

	// TODO provide key with cli
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
	registerRequest, err := receive.ReceiveAndValidate[dto.RegisterRequest](r, c.validator)
	if err != nil {
		c.logger.Info("Validation failed.", slog.Any("err", err))
		respond.Error(w, http.StatusBadRequest, err.Error(), c.logger)
	}

	c.logger.Info("Received register request.", slog.String("username", registerRequest.Email))

	err = c.authService.Register(r.Context(), registerRequest)
	if err != nil {
		c.logger.Error("Failed to register user", slog.Any("err", err))
		respond.Error(w, http.StatusBadRequest, "Failed to register user", c.logger)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest, err := receive.ReceiveAndValidate[dto.LoginRequest](r, c.validator)
	if err != nil {
		c.logger.Info("Validation failed.", slog.Any("err", err))
		respond.Error(w, http.StatusBadRequest, err.Error(), c.logger)
		return
	}

	userId, err := c.authService.Login(r.Context(), loginRequest)
	if err != nil {
		c.logger.Info("Login failed.", slog.Any("err", err))
		respond.Error(w, http.StatusBadRequest, "Login failed.", c.logger)
		return
	}

	err = c.setSession(userId, w)
	if err != nil {
		c.logger.Info("Set session failed.", slog.Any("err", err))
		respond.Error(w, http.StatusInternalServerError, "Login failed.", c.logger)
		return
	}

	respond.EncodeJson(w, http.StatusOK, dto.LoginResponse{UserId: userId}, c.logger)
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

	respond.EncodeJson(w, http.StatusOK, dto.WhoAmIResponse{UserId: r.Context().Value(SessionKey).(Session).UserId}, c.logger)
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
