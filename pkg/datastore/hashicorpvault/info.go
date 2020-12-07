package hashicorpvault

import (
	"os"
)

const (
	defaultHost  = "http://127.0.0.1"
	defaultToken = "root"
)

// Info attempts to pull connection string from local env vars, or
// set reasonable defaults
func (d *Datastore) info() {

	var value string

	value = os.Getenv("IDENGINE_VAULT_HOST")
	if len(value) > 0 {
		d.host = value
	} else {
		d.host = defaultHost
	}

	value = os.Getenv("IDENGINE_VAULT_TOKEN")
	if len(value) > 0 {
		d.token = value
	} else {
		d.token = defaultToken
	}
}
