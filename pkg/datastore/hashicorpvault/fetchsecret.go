package hashicorpvault

import (
	"context"
	"fmt"
	"github.com/hashicorp/vault/api"
	"log"
)

// FetchSecret fetches the secret(s) for an id
func (d *Datastore) FetchSecret(ctx context.Context, id string) (string, error) {

	log.Println("hashicorpvault.FetchSecret(): fetch secrets for id ", id)

	path := fmt.Sprintf(defaultIdentitySecretKeyFormat, id)

	var secret *api.Secret

	l := d.client.Logical()
	secret, err := l.Read(path)
	if err != nil {
		return "", err
	}

	if secret != nil {
		if data, ok := secret.Data["data"]; ok {

			m := data.(map[string]interface{})
			value := m["secret"].(string)

			return value, nil
		}
	}

	return "", nil
}
