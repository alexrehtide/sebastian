package controllertestutils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Testable interface {
	Test(req *http.Request, ts ...int) (*http.Response, error)
}

type TestCase struct {
	Req Req
	Res Res
}

type Req struct {
	Method string
	Route  string
	Body   string
}

type Res struct {
	Code int
	Body string
}

func TestJsonT(t *testing.T, app Testable, tcase TestCase) {
	resp, err := TestJson(app, tcase.Req)
	require.NoError(t, err)

	assert.Equal(t, tcase.Res.Code, resp.Code)
	assert.Equal(t, tcase.Res.Body, resp.Body)
}

func TestJson(app Testable, test Req) (Res, error) {
	req := httptest.NewRequest(test.Method, test.Route, strings.NewReader(test.Body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := app.Test(req, 1)
	if err != nil {
		return Res{}, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return Res{Code: resp.StatusCode, Body: string(buf.String())}, nil
}
