package testsetup

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func (ts *TestSetup) GetJsonWithCookie(t *testing.T, url string, cookie *http.Cookie, result interface{}) {
	ts.GetJsonWithRequestCustomizer(t, url, result, func(r *http.Request) {
		r.AddCookie(cookie)
	})
}

func (ts *TestSetup) GetJson(t *testing.T, url string, result interface{}) {
	ts.GetJsonWithRequestCustomizer(t, url, result, func(r *http.Request) {})
}

func (ts *TestSetup) GetExpectErrorStatus(t *testing.T, url string, expectedStatus int) {
	request, err := http.NewRequest("GET", ts.Server.URL+url, bytes.NewReader([]byte{}))
	require.Nil(t, err)

	response, err := ts.Server.Client().Do(request)
	require.Nil(t, err)
	defer response.Body.Close()

	require.Equal(t, expectedStatus, response.StatusCode)
}

func (ts *TestSetup) GetJsonWithRequestCustomizer(t *testing.T, url string, result interface{}, customizer func(r *http.Request)) {
	request, err := http.NewRequest("GET", ts.Server.URL+url, bytes.NewReader([]byte{}))
	require.Nil(t, err)

	customizer(request)

	response, err := ts.Server.Client().Do(request)
	require.Nil(t, err)
	defer response.Body.Close()

	require.Equal(t, http.StatusOK, response.StatusCode)
	require.Equal(t, "application/json", response.Header.Get("Content-Type"))
	err = json.NewDecoder(response.Body).Decode(&result)
	require.Nil(t, err)
}

func (ts *TestSetup) PostJsonWithCookie(t *testing.T, url string, body interface{}, cookie *http.Cookie, result interface{}) {
	ts.PostJsonWithRequestCustomizer(t, url, body, result, func(r *http.Request) {
		r.AddCookie(cookie)
	})
}

func (ts *TestSetup) PostJson(t *testing.T, url string, body interface{}, result interface{}) {
	ts.PostJsonWithRequestCustomizer(t, url, body, result, func(r *http.Request) {})
}

func (ts *TestSetup) PostJsonWithRequestCustomizer(t *testing.T, url string, body interface{}, result interface{}, customizer func(r *http.Request)) {
	bodyBytes, err := json.Marshal(body)
	require.Nil(t, err)
	request, err := http.NewRequest("POST", ts.Server.URL+url, bytes.NewReader(bodyBytes))
	require.Nil(t, err)

	customizer(request)

	response, err := ts.Server.Client().Do(request)
	require.Nil(t, err)
	defer response.Body.Close()

	require.Equal(t, http.StatusOK, response.StatusCode)
	if result != nil {
		require.Equal(t, "application/json", response.Header.Get("Content-Type"))
		err = json.NewDecoder(response.Body).Decode(&result)
		require.Nil(t, err)
	}
}
