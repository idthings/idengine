package hashicorpvault

import (
	"context"
	"log"
)

// FetchSecret fetches the secret(s) for an id
func (d *Datastore) FetchSecret(ctx context.Context, id string) (string, error) {
	log.Println("hashicorpvault.FetchSecret(): store for id ", id)
	return "", nil
}
