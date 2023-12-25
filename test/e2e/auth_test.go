package e2e

import (
	"net/http"
	"testing"

	"github.com/bomgar/basicwebapp/test/e2e/request"
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

	whoAmIBody := dto.WhoAmIResponse{}
	request.GetJson(t, ts, "/api/whoami", cookie, &whoAmIBody)

	assert.Equal(t, loginResponse.UserId, whoAmIBody.UserId)
}
