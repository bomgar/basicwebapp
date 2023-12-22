package e2e

import (
	"testing"

	"github.com/bomgar/basicwebapp/test/e2e/setup"
	"github.com/stretchr/testify/assert"
)

func TestAuthWhoAmI(t *testing.T) {
	ts := setup.TestServer()
	defer ts.Close()
	rs, err := ts.Client().Get(ts.URL + "/whoami")
	assert.Nil(t, err)

	assert.Equal(t, 200, rs.StatusCode)

}
