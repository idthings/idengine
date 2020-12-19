package redis

import (
	"context"
	"errors"
	"fmt"
)

// StoreSecret stores
func (d *Datastore) StoreSecret(ctx context.Context, id string, secret string) error {

	//log.Println("datastore.StoreSecret():", id, secret) // use for development only

	if len(id) < 1 {
		return errors.New("datastore.StoreSecret(): invalid id")
	}

	if len(secret) < 1 {
		return errors.New("datastore.StoreSecret(): invalid secret")
	}

	key := fmt.Sprintf(defaultIdentitySecretKeyFormat, id)

	// set the secret value at key, with not expiry
	_, err := d.client.Set(d.ctx, key, secret, 0).Result()
	if err != nil {
		return err
	}

	return nil
}
