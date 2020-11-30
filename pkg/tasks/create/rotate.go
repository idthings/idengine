package create

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/idthings/idengine/pkg/data"
	"log"
	"net/http"
	"strings"
)

const (
	authHeaderName = "X-idThings-Password"
)

// RotateSecretInterface rotates
type RotateSecretInterface interface {
	FetchSecret(ctx context.Context, id string) (string, error)
	StoreSecret(ctx context.Context, id string, secret string) error
}

// RotateSecret runs
func RotateSecret(ctx context.Context, store RotateSecretInterface, r *http.Request) (int, string) {

	var i identity

	// extract guid from http request, the last element
	elements := strings.Split(r.URL.Path, "/") // always returns at least one element
	i.ID = elements[len(elements)-1]

	// valid guid string format
	_, err := uuid.Parse(i.ID)
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
	secret, err := store.FetchSecret(ctx, i.ID)
	if err != nil {
		log.Println("validate.Secret():", err.Error())
		return http.StatusInternalServerError, "Internal error\n"
	}

	if secret != passwordString {
		return http.StatusUnauthorized, "Unauthorized\n"
	}

	i.Secret = data.NewPassword() // our new secret

	if err := store.StoreSecret(ctx, i.ID, i.Secret); err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	var format string
	if val, ok := r.URL.Query()["format"]; ok { // returns a slice
		format = val[0]
	}

	var response string

	switch format {
	case "stream":
		response = fmt.Sprintf("{%s,%s}", i.ID, i.Secret)

	default:
		response = responseAsJSONString(i)
	}

	//	response := fmt.Sprintf("{%s}\n", secret)

	return http.StatusOK, response
}
