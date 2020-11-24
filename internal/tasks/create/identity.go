package create

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/idthings/idengine/internal/data"
	"net/http"
)

// StoreSecretInterface stores a secret for an id in persistent storage
type StoreSecretInterface interface {
	StoreSecret(id string, secret string) error
}

// Identity runs
func Identity(store StoreSecretInterface, r *http.Request) (int, string) {

	id := uuid.New().String()
	secret := data.NewPassword()

	if err := store.StoreSecret(id, secret); err != nil {
		return http.StatusInternalServerError, err.Error()
	}

	response := fmt.Sprintf("{%s,%s}\n", id, secret)

	return http.StatusOK, response
}
