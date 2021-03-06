package create

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/idthings/idengine/pkg/data"
	"log"
	"net/http"
)

type identity struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
}

// StoreSecretInterface stores a secret for an id in persistent storage
type StoreSecretInterface interface {
	StoreSecret(ctx context.Context, id string, secret string) error
}

// Identity runs
func Identity(store StoreSecretInterface, r *http.Request) (int, string) {

	var i identity
	i.ID = uuid.New().String()
	i.Secret = data.NewPassword()

	if err := store.StoreSecret(r.Context(), i.ID, i.Secret); err != nil {
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

	return http.StatusOK, response
}

func responseAsJSONString(i identity) string {

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(&i); err != nil {
		log.Println("identity.responseAsJSONString(): ", err.Error())
		return ""
	}

	return buf.String()
}
