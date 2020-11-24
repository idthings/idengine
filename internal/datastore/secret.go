package datastore

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

// StoreSecret stores
func (d *Datastore) StoreSecret(id string, secret string) error {

	log.Println("datastore.StoreSecret():", id, secret)

	if len(id) < 1 {
		return errors.New("datastore.StoreSecret(): invalid id")
	}

	if len(secret) < 1 {
		return errors.New("datastore.StoreSecret(): invalid secret")
	}

	key := fmt.Sprintf(identitySecretsKeyFormat, id)
	timestamp := time.Now().UnixNano() / 1e6 // convert to milliseconds

	_, err := d.client.ZAdd(d.ctx, key, &redis.Z{
		Score:  float64(timestamp),
		Member: secret,
	}).Result()
	if err != nil {
		return err
	}

	// TODO: limit set size to 10

	return nil
}

// FetchSecret fetches
func (d *Datastore) FetchSecret(id string) (string, error) {

	log.Println("datastore.FetchSecret():", id)

	return "", nil
}
