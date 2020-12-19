package hashicorpvault

import (
	"log"
	"os"
)

const (
	defaultHost  = "http://127.0.0.1:8200"
	defaultToken = "developer"
)

// info attempts to pull connection string from local env vars, or
// set reasonable defaults
func (d *Datastore) info() {

	var value string

	value = os.Getenv("IDENGINE_VAULT_HOST")
	if len(value) > 0 {
		d.host = value
		log.Println("hashicorpvault.info(): overriding vault host from env var IDENGINE_VAULT_HOST")
	} else {
		d.host = defaultHost
	}

	value = os.Getenv("IDENGINE_VAULT_TOKEN")
	if len(value) > 0 {
		d.token = value
		log.Println("hashicorpvault.info(): overriding vault host from env var IDENGINE_VAULT_TOKEN")
	} else {
		d.token = defaultToken
	}
}

// Info conforms
func (d *Datastore) Info() {}
