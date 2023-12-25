package request

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bomgar/basicwebapp/test/e2e/setup"
	"github.com/stretchr/testify/assert"
)

func GetJson[R any](t *testing.T, ts *setup.TestSetup, url string, cookie *http.Cookie, result *R) {
	request, err := http.NewRequest("GET", ts.Server.URL+url, bytes.NewReader([]byte{}))
	request.AddCookie(cookie)
	assert.Nil(t, err)
	response, err := ts.Server.Client().Do(request)
	assert.Nil(t, err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	err = json.NewDecoder(response.Body).Decode(&result)
	assert.Nil(t, err)
}
