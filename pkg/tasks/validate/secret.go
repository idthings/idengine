package validate

import (
	"log"
	"net/http"
)

const (
	authHeaderName = "X-idThings-Password"
)

// FetchSecretInterface fetches a secret from persistent storage
type FetchSecretInterface interface {
	FetchSecret(id string) (string, error)
}

// Secret runs
func Secret(store FetchSecretInterface, r *http.Request) (int, string) {

	// get the requesting device id
	id, err := extractGUID(r.URL.Path)
	if err != nil {
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
