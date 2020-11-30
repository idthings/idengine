package create

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type headerMap map[string]string

var testRotateItems = []struct {
	comment           string
	method            string
	url               string
	headers           headerMap
	fetchReturnString string
	fetchReturnError  error
	storeReturnError  error
	expectedStatus    int
	expectedResponse  string
}{
	{
		comment: "valid request",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		fetchReturnString: "UCg0&3DBR%C%q0D!5!*9",
		fetchReturnError:  nil,
		storeReturnError:  nil,
		expectedStatus:    200,
		expectedResponse:  "OK\n",
	},
	{
		comment: "valid request stream format",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30?format=stream",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		fetchReturnString: "UCg0&3DBR%C%q0D!5!*9",
		fetchReturnError:  nil,
		storeReturnError:  nil,
		expectedStatus:    200,
		expectedResponse:  "OK\n",
	},
	{
		comment: "valid request, json response",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30?format=json",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		fetchReturnString: "UCg0&3DBR%C%q0D!5!*9",
		fetchReturnError:  nil,
		storeReturnError:  nil,
		expectedStatus:    200,
		expectedResponse:  "OK\n",
	},
	{
		comment: "valid request, long path",
		method:  "GET",
		url:     "http://localhost/some/extension/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		fetchReturnString: "UCg0&3DBR%C%q0D!5!*9",
		fetchReturnError:  nil,
		storeReturnError:  nil,
		expectedStatus:    200,
		expectedResponse:  "OK\n",
	},
	{
		comment: "invalid password",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		fetchReturnString: "UCg0&3DBR%C%q0D!5!*",
		fetchReturnError:  nil,
		storeReturnError:  nil,
		expectedStatus:    401,
		expectedResponse:  "Unauthorized\n",
	},
	{
		comment: "invalid guid",
		method:  "GET",
		url:     "http://localhost/identities/-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		fetchReturnString: "UCg0&3DBR%C%q0D!5!*9",
		fetchReturnError:  nil,
		storeReturnError:  nil,
		expectedStatus:    404,
		expectedResponse:  "Not Found",
	},
	{
		comment: "no password",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "",
		},
		fetchReturnString: "UCg0&3DBR%C%q0D!5!*9",
		fetchReturnError:  nil,
		storeReturnError:  nil,
		expectedStatus:    400,
		expectedResponse:  "Bad Request: missing header",
	},
	{
		comment:           "no password header",
		method:            "GET",
		url:               "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers:           headerMap{},
		fetchReturnString: "UCg0&3DBR%C%q0D!5!*9",
		fetchReturnError:  nil,
		storeReturnError:  nil,
		expectedStatus:    400,
		expectedResponse:  "Bad Request: missing header",
	},
	{
		comment: "fetch data store error",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		fetchReturnString: "UCg0&3DBR%C%q0D!5!*9",
		fetchReturnError:  errors.New("data store error"),
		storeReturnError:  nil,
		expectedStatus:    500,
		expectedResponse:  "Internal error\n",
	},
	{
		comment: "store data store error",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Password": "UCg0&3DBR%C%q0D!5!*9",
		},
		fetchReturnString: "UCg0&3DBR%C%q0D!5!*9",
		fetchReturnError:  nil,
		storeReturnError:  errors.New("data store error"),
		expectedStatus:    500,
		expectedResponse:  "Internal error\n",
	},
}

type mockRotateSecretStore struct {
	fetchReturnString string
	fetchReturnError  error
	storeReturnError  error
}

func (m mockRotateSecretStore) StoreSecret(ctx context.Context, id string, secret string) error {
	return m.storeReturnError
}

func (m mockRotateSecretStore) FetchSecret(ctx context.Context, id string) (string, error) {
	return m.fetchReturnString, m.fetchReturnError
}

func TestRotateSecretGUID(t *testing.T) {

	var mock mockRotateSecretStore
	var req *http.Request

	for _, item := range testRotateItems {

		mock.fetchReturnString = item.fetchReturnString
		mock.fetchReturnError = item.fetchReturnError
		mock.storeReturnError = item.storeReturnError

		byteBuf := bytes.NewBuffer([]byte(""))
		req, _ = http.NewRequest(item.method, item.url, byteBuf)

		for k, v := range item.headers {
			req.Header.Add(k, v)
		}

		status, _ := RotateSecret(mock, req)
		assert.Equal(t, item.expectedStatus, status, item.comment)
	}
}
