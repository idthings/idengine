package redis

import (
	"context"
	"fmt"
	"log"
)

// FetchSecret fetches
func (d *Datastore) FetchSecret(ctx context.Context, id string) (string, error) {

	log.Println("datastore.FetchSecret():", id)

	key := fmt.Sprintf(defaultIdentitySecretKeyFormat, id)

	secret, err := d.client.Get(d.ctx, key).Result()
	if err != nil {
		return "", err
	}

	return secret, nil
}
