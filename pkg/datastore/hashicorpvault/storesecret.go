package hashicorpvault

import (
	"context"
	"log"
)

// StoreSecret stores an id's secret
func (d *Datastore) StoreSecret(ctx context.Context, id string, secret string) error {
	log.Println("hashicorpvault.StoreSecret(): store for id ", id)
	return nil
}
