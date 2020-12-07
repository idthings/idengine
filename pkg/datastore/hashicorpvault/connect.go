package hashicorpvault

import (
	"github.com/hashicorp/vault/api"
	"log"
)

// Datastore to wrap HashiCorp Vault
type Datastore struct {
	host   string
	token  string
	client *api.Client
}

// Connect attempts to configure connection and test with health check
// Returns:
//   true if we can check the Vault status
//   false otherwise
func (d *Datastore) Connect() bool {

	d.info()

	config := &api.Config{
		Address: d.host,
	}

	var err error
	d.client, err = api.NewClient(config)
	if err != nil {
		log.Println("hashicorpvault.Connect(): connection failed. ", err.Error())
		return false
	}

	d.client.SetToken(d.token)

	sys := d.client.Sys()
	health, err := sys.Health()
	if err != nil {
		log.Println("hashicorpvault.Connect(): connection failed. ", err.Error())
		return false
	}

	log.Printf("Vault node: %s, Sealed: %t\n", health.ClusterName, health.Sealed)

	return true
}
