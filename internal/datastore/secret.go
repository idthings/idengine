package datastore

import (
	"log"
)

// StoreSecret stores
func (d Datastore) StoreSecret(id string, secret string) error {

	log.Println("datastore.StoreSecret():", id, secret)

	return nil
}

func (d Datastore) FetchSecret(id string) (string, error) {
	return "", nil
}
