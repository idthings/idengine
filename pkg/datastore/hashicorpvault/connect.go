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
//   true if we connect and find Vault unsealed
//   false otherwise
//
// implicitly this method returns "are we ready to process"
func (d *Datastore) Connect() bool {

	d.info() // set defaults or use env vars

	config := &api.Config{
		Address: d.host,
	}

	var err error
	d.client, err = api.NewClient(config)
	if err != nil {
		log.Println("hashicorpvault.Connect(): configuration failed. ", err.Error())
		return false
	}

	d.client.SetToken(d.token)

	sys := d.client.Sys()
	health, err := sys.Health()
	if err != nil {
		log.Println("hashicorpvault.Connect(): connection failed. ", err.Error())
		return false
	}

	if health.Sealed {
		log.Println("hashicorpvault.Connect(): connection failed. Vault is currently sealed.")
		return false
	}

	log.Printf("hashicorpvault.Connect(): successful connection to Vault node: %s\n", health.ClusterName)
	return true
}
