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

// Connect attempts a connection to a Vault node
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

	return true
}
