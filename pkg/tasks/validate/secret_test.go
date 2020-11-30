package validate

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var testSecretItems = []struct {
	comment          string
	method           string
	url              string
	headers          headerMap
	mockReturnString string
	mockReturnError  error
	expectedStatus   int
	expectedResponse string
}{
	{
		comment: "valid request",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		mockReturnString: "UCg0&3DBR%C%q0D!5!*9",
		mockReturnError:  nil,
		expectedStatus:   200,
		expectedResponse: "OK",
	},
	{
		comment: "valid request, long path",
		method:  "GET",
		url:     "http://localhost/some/extension/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		mockReturnString: "UCg0&3DBR%C%q0D!5!*9",
		mockReturnError:  nil,
		expectedStatus:   200,
		expectedResponse: "OK",
	},
	{
		comment: "invalid password",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		mockReturnString: "UCg0&3DBR%C%q0D!5!*",
		mockReturnError:  nil,
		expectedStatus:   401,
		expectedResponse: "Unauthorized",
	},
	{
		comment: "invalid guid",
		method:  "GET",
		url:     "http://localhost/identities/-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		mockReturnString: "",
		mockReturnError:  nil,
		expectedStatus:   404,
		expectedResponse: "Not Found",
	},
	{
		comment: "no password",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "",
		},
		mockReturnString: "",
		mockReturnError:  nil,
		expectedStatus:   400,
		expectedResponse: "Bad Request: missing header",
	},
	{
		comment:          "no password header",
		method:           "GET",
		url:              "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers:          headerMap{},
		mockReturnString: "",
		mockReturnError:  nil,
		expectedStatus:   400,
		expectedResponse: "Bad Request: missing header",
	},
	{
		comment: "data store error",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		mockReturnString: "",
		mockReturnError:  errors.New("data store error"),
		expectedStatus:   500,
		expectedResponse: "Internal error",
	},
}

type mockSecretStore struct {
	returnString string
	returnError  error
}

func (m mockSecretStore) FetchSecret(ctx context.Context, id string) (string, error) {
	return m.returnString, m.returnError
}

type headerMap map[string]string

func TestSecret(t *testing.T) {

	var mock mockSecretStore
	var req *http.Request

	for _, item := range testSecretItems {

		mock.returnString = item.mockReturnString
		mock.returnError = item.mockReturnError

		byteBuf := bytes.NewBuffer([]byte(""))
		req, _ = http.NewRequest(item.method, item.url, byteBuf)

		for k, v := range item.headers {
			req.Header.Add(k, v)
		}

		status, response := Secret(mock, req)
		assert.Equal(t, item.expectedStatus, status, item.comment)
		assert.Equal(t, item.expectedResponse, response, item.comment)
	}
}
