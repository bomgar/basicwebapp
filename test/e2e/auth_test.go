package e2e

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/bomgar/basicwebapp/test/e2e/setup"
	"github.com/bomgar/basicwebapp/web/dto"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ts := setup.Setup()
	defer ts.Close()

	registerRequest := dto.RegisterRequest{
		Email: "fkbr",
		Password: "sxoe",
	}
	body, err := json.Marshal(registerRequest)
	assert.Nil(t, err)

	rs, err := ts.Server.Client().Post(ts.Server.URL+"/register", "application/json", bytes.NewReader(body))
	assert.Nil(t, err)

	assert.Equal(t, 200, rs.StatusCode)
}

func TestAuthWhoAmI(t *testing.T) {
	ts := setup.Setup()
	defer ts.Close()

	rs, err := ts.Server.Client().Get(ts.Server.URL + "/whoami")
	assert.Nil(t, err)
	defer rs.Body.Close()

	assert.Equal(t, 200, rs.StatusCode)

}
