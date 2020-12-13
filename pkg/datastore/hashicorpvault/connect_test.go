package hashicorpvault

import (
	"fmt"
	kv "github.com/hashicorp/vault-plugin-secrets-kv"
	"github.com/hashicorp/vault/api"
	vaulthttp "github.com/hashicorp/vault/http"
	"github.com/hashicorp/vault/sdk/helper/logging"
	"github.com/hashicorp/vault/sdk/logical"
	hashivault "github.com/hashicorp/vault/vault"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConnectWithoutVault(t *testing.T) {

	const testComment = "Test without Vault running"

	v := Datastore{}
	result := v.Connect()

	assert.Equal(t, false, result, testComment)
}

func TestConnectWithVault(t *testing.T) {

	const testComment = "Test with unsealed Vault"

	/*
	 * Setup our cluster and env vars
	 */

	cluster := createVaultTestCluster(t)
	defer cluster.Cleanup()

	// listener address is dynamic with TestCluster, so extract it now
	listener := cluster.Cores[0].Listeners[0]
	address := fmt.Sprintf("https://%s:%d", listener.Address.IP, listener.Address.Port)

	// do not fail on self-signed certs during testing
	os.Setenv("VAULT_SKIP_VERIFY", "true")
	defer os.Unsetenv("VAULT_SKIP_VERIFY")

	/*
	 * Begin testing our own code
	 */

	// our vault client implementation finds the cluster address in this env var
	os.Setenv("IDENGINE_VAULT_HOST", address)
	defer os.Unsetenv("IDENGINE_VAULT_HOST")

	v := Datastore{}
	v.host = address

	result := v.Connect()
	assert.Equal(t, true, result, testComment)
}

func createVaultTestCluster(t *testing.T) *hashivault.TestCluster {

	t.Helper()

	// See the url for log level Int's
	// https://github.com/hashicorp/go-hclog/blob/master/logger.go
	logger := logging.NewVaultLogger(6) // Off

	coreConfig := &hashivault.CoreConfig{
		LogicalBackends: map[string]logical.Factory{
			"kv": kv.Factory,
		},
		Logger: logger, // add our own logger, to reduce chatter
	}

	cluster := hashivault.NewTestCluster(t, coreConfig, &hashivault.TestClusterOptions{
		NumCores:    1, // default is 3 cores, but these test only require 1
		HandlerFunc: vaulthttp.Handler,
	})
	cluster.Start()

	// Create KV V2 mount
	if err := cluster.Cores[0].Client.Sys().Mount("kv", &api.MountInput{
		Type: "kv",
		Options: map[string]string{
			"version": "2",
		},
	}); err != nil {
		t.Fatal(err)
	}

	return cluster
}
