package datastore

import (
	"context"
	"errors"
	"fmt"
	"log"
)

// StoreSecret stores
func (d *Datastore) StoreSecret(ctx context.Context, id string, secret string) error {

	log.Println("datastore.StoreSecret():", id, secret)

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
