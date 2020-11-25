package validate

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var testDigestItems = []struct {
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
			"X-idThings-Digest": "HMAC-SHA256,f62100c007ec7630a6d65c0d7d745dae5a21da5d8474722e6aa065c15b6ca9c0,1604573826351,my data",
		},
		mockReturnString: "UCg0&3DBR%C%q0D!5!*9",
		mockReturnError:  nil,
		expectedStatus:   401,
		expectedResponse: "Digest Expired",
	},
	{
		comment: "invalid guid in request",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-",
		headers: headerMap{
			"X-idThings-Digest": "digest-type,digest,timestamp,data",
		},
		mockReturnString: "UCg0&3DBR%C%q0D!5!*9",
		mockReturnError:  nil,
		expectedStatus:   404,
		expectedResponse: "Not Found",
	},
	{
		comment:          "missing digest header in request",
		method:           "GET",
		url:              "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		mockReturnString: "UCg0&3DBR%C%q0D!5!*9",
		mockReturnError:  nil,
		expectedStatus:   400,
		expectedResponse: "Bad Request: missing header",
	},
	{
		comment: "invalid digest header request",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Digest": "HMAC-SHA256,asdasd",
		},
		mockReturnString: "UCg0&3DBR%C%q0D!5!*9",
		mockReturnError:  nil,
		expectedStatus:   400,
		expectedResponse: "Bad Request: invalid header",
	},
	{
		comment: "unsupported digest type request",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Digest": "HMAC-SHA2020,718a2a26353efce2d754b26a50b8824b1dc2e143e0aca7e2f1b55f6f7d2d5943,1604573826351,my data",
		},
		mockReturnString: "UCg0&3DBR%C%q0D!5!*9",
		mockReturnError:  nil,
		expectedStatus:   401,
		expectedResponse: "Unauthorized",
	},
	{
		comment: "data store error",
		method:  "GET",
		url:     "http://localhost/identities/cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		headers: headerMap{
			"X-idThings-Digest": "HMAC-SHA256,digest,timestamp,data",
		},
		mockReturnString: "",
		mockReturnError:  errors.New("data store error"),
		expectedStatus:   500,
		expectedResponse: "Internal error",
	},
}

func TestDigest(t *testing.T) {

	var mock mockSecretStore
	var req *http.Request

	for _, item := range testDigestItems {

		mock.returnString = item.mockReturnString
		mock.returnError = item.mockReturnError

		byteBuf := bytes.NewBuffer([]byte(""))
		req, _ = http.NewRequest(item.method, item.url, byteBuf)

		for k, v := range item.headers {
			req.Header.Add(k, v)
		}

		status, response := Digest(mock, req)
		assert.Equal(t, item.expectedStatus, status, item.comment)
		assert.Equal(t, item.expectedResponse, response, item.comment)
	}
}
