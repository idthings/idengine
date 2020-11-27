package datastore

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// StoreSecret stores
func (d *Datastore) StoreSecret(id string, secret string, expirationDays int) error {

	log.Println("datastore.StoreSecret():", id, secret)

	if len(id) < 1 {
		return errors.New("datastore.StoreSecret(): invalid id")
	}

	if len(secret) < 1 {
		return errors.New("datastore.StoreSecret(): invalid secret")
	}

	key := fmt.Sprintf(identitySecretsKeyFormat, id)

	expiration := time.Duration(expirationDays) * (time.Hour * 24)

	_, err := d.client.Set(d.ctx, key, secret, expiration).Result()
	if err != nil {
		return err
	}

	// TODO: limit set size to 10

	return nil
}

// FetchSecret fetches
func (d *Datastore) FetchSecret(id string) (string, error) {

	log.Println("datastore.FetchSecret():", id)

	key := fmt.Sprintf(identitySecretsKeyFormat, id)

	secret, err := d.client.Get(d.ctx, key).Result()
	if err != nil {
		return "", err
	}

	return secret, nil
}
