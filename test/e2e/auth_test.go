package e2e

import (
	"testing"

	"github.com/bomgar/basicwebapp/test/e2e/setup"
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

	rs, err := ts.Server.Client().Get(ts.Server.URL + "/whoami")
	assert.Nil(t, err)
	defer rs.Body.Close()

	assert.Equal(t, 200, rs.StatusCode)

}
