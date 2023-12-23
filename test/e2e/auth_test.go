package e2e

import (
	"testing"

	"github.com/bomgar/basicwebapp/test/e2e/setup"
	"github.com/stretchr/testify/assert"
)

func TestAuthWhoAmI(t *testing.T) {
	ts := setup.Setup()
	defer ts.Close()

	rs, err := ts.Server.Client().Get(ts.Server.URL + "/whoami")
	assert.Nil(t, err)
	defer rs.Body.Close()

	assert.Equal(t, 200, rs.StatusCode)

}
