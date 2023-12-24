package setup

import (
	"bytes"
	"encoding/json"
	"net/http"
	"slices"
	"testing"

	"github.com/bomgar/basicwebapp/web/controllers/authcontroller"
	"github.com/bomgar/basicwebapp/web/dto"
	"github.com/stretchr/testify/assert"
)

func (ts *TestSetup) RegisterUser(t *testing.T, email string, password string) {
	registerRequest := dto.RegisterRequest{
		Email:    email,
		Password: password,
	}
	body, err := json.Marshal(registerRequest)
	assert.Nil(t, err)

	rs, err := ts.Server.Client().Post(ts.Server.URL+"/register", "application/json", bytes.NewReader(body))
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, rs.StatusCode)

}

func (ts *TestSetup) LoginUser(t *testing.T, email string, password string) (dto.LoginResponse, *http.Cookie) {
	loginRequest := dto.LoginRequest{
		Email:    email,
		Password: password,
	}
	body, err := json.Marshal(loginRequest)
	assert.Nil(t, err)

	rs, err := ts.Server.Client().Post(ts.Server.URL+"/login", "application/json", bytes.NewReader(body))
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, rs.StatusCode)
	cookieIndex := slices.IndexFunc(rs.Cookies(), func(cookie *http.Cookie) bool {
		return cookie.Name == authcontroller.CookieName
	})
	assert.NotEqual(t, -1, cookieIndex)

	loginResponse := dto.LoginResponse{}
	err = json.NewDecoder(rs.Body).Decode(&loginResponse)
	assert.Nil(t, err)
	return loginResponse, rs.Cookies()[cookieIndex]
}
