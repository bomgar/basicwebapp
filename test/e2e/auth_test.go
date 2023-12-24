package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bomgar/basicwebapp/test/e2e/setup"
	"github.com/bomgar/basicwebapp/web/dto"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ts := setup.Setup(t)
	defer ts.Close()

	ts.RegisterUser(t, "fkbr@sxoe.kuci", "fkbr")
}

func TestLogin(t *testing.T) {
	ts := setup.Setup(t)
	defer ts.Close()

	ts.RegisterUser(t, "fkbr@sxoe.kuci", "fkbr")
	ts.LoginUser(t, "fkbr@sxoe.kuci", "fkbr")
}

func TestAuthWhoAmI(t *testing.T) {
	ts := setup.Setup(t)
	defer ts.Close()

	rs, err := ts.Server.Client().Get(ts.Server.URL + "/api/whoami")
	assert.Nil(t, err)
	rs.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, rs.StatusCode)

	ts.RegisterUser(t, "fkbr@sxoe.kuci", "fkbr")
	loginResponse, cookie := ts.LoginUser(t, "fkbr@sxoe.kuci", "fkbr")

	whoamIRequest, err := http.NewRequest("GET", ts.Server.URL+"/api/whoami", bytes.NewReader([]byte{}))
	assert.Nil(t, err)
	whoamIRequest.AddCookie(cookie)

	whoAmIResponse, err := ts.Server.Client().Do(whoamIRequest)
	assert.Nil(t, err)
	defer whoAmIResponse.Body.Close()

	assert.Equal(t, http.StatusOK, whoAmIResponse.StatusCode)
	whoAmIBody := dto.WhoAmIResponse{}
	err = json.NewDecoder(whoAmIResponse.Body).Decode(&whoAmIBody)
	assert.Nil(t, err)

	assert.Equal(t, loginResponse.UserId, whoAmIBody.UserId)
}
