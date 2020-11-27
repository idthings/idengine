package create

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var testIdentityItems = []struct {
	comment         string
	method          string
	url             string
	mockReturnValue error
	expectedInt     int
}{
	{
		comment:         "valid request",
		method:          "GET",
		url:             "http://localhost/identities/new/",
		mockReturnValue: nil,
		expectedInt:     200,
	},
	{
		comment:         "data store error",
		method:          "GET",
		url:             "http://localhost/identities/new/",
		mockReturnValue: errors.New("data store error"),
		expectedInt:     500,
	},
	{
		comment:         "valid request, json response",
		method:          "GET",
		url:             "http://localhost/identities/new/?format=json",
		mockReturnValue: nil,
		expectedInt:     200,
	},
}

type mockSecretStore struct {
	returnValue error
}

func (m mockSecretStore) StoreSecret(id string, secret string, expirationDays int) error {
	return m.returnValue
}

func TestGetIdentityGUID(t *testing.T) {

	var mock mockSecretStore
	var req *http.Request

	for _, item := range testIdentityItems {

		mock.returnValue = item.mockReturnValue

		byteBuf := bytes.NewBuffer([]byte(""))
		req, _ = http.NewRequest(item.method, item.url, byteBuf)

		status, _ := Identity(mock, req)
		assert.Equal(t, item.expectedInt, status, item.comment)
	}
}

var testJSONResponseItems = []struct {
	comment     string
	id          string
	secret      string
	returnValue string
}{
	{
		comment:     "valid id, secret",
		id:          "cbfbe13d-0ab4-487e-89bf-276dcd646a30",
		secret:      "UCg0&3DBR%C%q0D!5!*9",
		returnValue: `{"id":"cbfbe13d-0ab4-487e-89bf-276dcd646a30","secret":"UCg0&3DBR%C%q0D!5!*9"}`,
	},
}

func TestJSONResponse(t *testing.T) {
	for _, item := range testJSONResponseItems {
		i := identity{item.id, item.secret}
		assert.Equal(t, item.returnValue+"\n", responseAsJSONString(i), item.comment)
	}
}
