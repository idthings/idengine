package validate

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
)

const (
	authHeaderName = "X-idThings-Password"
	uriPrefix      = "/identities/"
)

// FetchSecretInterface fetches a secret from persistent storage
type FetchSecretInterface interface {
	FetchSecret(id string) (string, error)
}

// Secret runs
func Secret(store FetchSecretInterface, r *http.Request) (int, string) {

	// extract guid from http request
	id := strings.Replace(r.URL.RequestURI(), uriPrefix, "", 1)
	_, err := uuid.Parse(id)
	if err != nil {
		log.Println("validate.Secret(): id not a valid guid")
		return http.StatusNotFound, "Not Found"
	}

	passwordString := r.Header.Get(authHeaderName) // case insensitive
	if len(passwordString) == 0 {
		log.Println("validate.Secret(): auth header string empty")
		return http.StatusBadRequest, "Bad Request: missing header"
	}

	secret, err := store.FetchSecret(id)
	if err != nil {
		log.Println("validate.Secret():", err.Error())
		return http.StatusInternalServerError, "Internal error\n"
	}

	if secret == passwordString {
		return http.StatusOK, "OK\n"
	}

	return http.StatusUnauthorized, "Unauthorized\n"
}
