package authcontroller

import (
	"context"
	"net/http"
	"slices"
	"time"
)

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
