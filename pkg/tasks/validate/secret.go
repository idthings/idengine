package validate

import (
	"context"
	"log"
	"net/http"
)

const (
	authHeaderName = "X-idThings-Password"
)

// FetchSecretInterface fetches a secret from persistent storage
type FetchSecretInterface interface {
	FetchSecret(ctx context.Context, id string) (string, error)
}

// Secret runs
func Secret(ctx context.Context, store FetchSecretInterface, r *http.Request) (int, string) {

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

	secret, err := store.FetchSecret(ctx, id)
	if err != nil {
		log.Println("validate.Secret():", err.Error())
		return http.StatusInternalServerError, "Internal error"
	}

	if secret == passwordString {
		return http.StatusOK, "OK"
	}

	return http.StatusUnauthorized, "Unauthorized"
}
