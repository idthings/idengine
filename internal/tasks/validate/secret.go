package validate

import (
	"fmt"
	"net/http"
)

// FetchSecretInterface fetches a secret from persistent storage
type FetchSecretInterface interface {
	FetchSecret(id string) (string, error)
}

// Secret runs
func Secret(store FetchSecretInterface, r *http.Request) (int, string) {

	id := "id"

	secret, err := store.FetchSecret(id)
	if err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	response := fmt.Sprintf("{%s,%s}\n", id, secret)

	return http.StatusOK, response
}
