package hashicorpvault

import (
	"context"
	"fmt"
	"log"
)

// StoreSecret stores an id's secret
func (d *Datastore) StoreSecret(ctx context.Context, id string, secret string) error {
	log.Println("hashicorpvault.StoreSecret(): store for id ", id)

	path := fmt.Sprintf(defaultIdentitySecretKeyFormat, id)

	options := map[string]interface{}{
		"cas": 0,
	}

	data := map[string]interface{}{
		"id":     id,
		"secret": secret,
	}

	payload := map[string]interface{}{
		"options": options,
		"data":    data,
	}

	l := d.client.Logical()
	response, err := l.Write(path, payload)
	if err != nil {
		return err
	}

	fmt.Println("hashicorpvault.StoreSecret(): ", payload)
	fmt.Println("hashicorpvault.StoreSecret(): ", response)

	return nil
}
