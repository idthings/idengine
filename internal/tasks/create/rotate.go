package create

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/idthings/idengine/internal/data"
	"log"
	"net/http"
	"strings"
)

const (
	authHeaderName = "X-idThings-Password"
)

// RotateSecretInterface rotates
type RotateSecretInterface interface {
	FetchSecret(id string) (string, error)
	StoreSecret(id string, secret string) error
}

// RotateSecret runs
func RotateSecret(store RotateSecretInterface, r *http.Request) (int, string) {

	// extract guid from http request, the last element
	elements := strings.Split(r.URL.RequestURI(), "/") // always returns at least one element
	id := elements[len(elements)-1]

	// valid guid string format
	_, err := uuid.Parse(id)
	if err != nil {
		log.Println("create.RotateSecret(): id not a valid guid")
		return http.StatusNotFound, "Not Found"
	}

	// check password was supplied
	passwordString := r.Header.Get(authHeaderName) // case insensitive
	if len(passwordString) == 0 {
		log.Println("validate.Secret(): auth header string empty")
		return http.StatusBadRequest, "Bad Request: missing header"
	}

	// fetch existing password for this id
	secret, err := store.FetchSecret(id)
	if err != nil {
		log.Println("validate.Secret():", err.Error())
		return http.StatusInternalServerError, "Internal error\n"
	}

	if secret != passwordString {
		return http.StatusUnauthorized, "Unauthorized\n"
	}

	newSecret := data.NewPassword()

	if err := store.StoreSecret(id, newSecret); err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	response := fmt.Sprintf("{%s}\n", secret)

	return http.StatusOK, response
}
