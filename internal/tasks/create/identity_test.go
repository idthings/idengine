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
}

type mockSecretStore struct {
	returnValue error
}

func (m mockSecretStore) StoreSecret(id string, secret string) error {
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
