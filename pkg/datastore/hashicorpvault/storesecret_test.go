package hashicorpvault

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreSecret(t *testing.T) {

	const testComment = "Test storing secret in Vault"

	cluster := createVaultTestCluster(t)
	defer cluster.Cleanup()

	ds := Datastore{}
	ds.client = cluster.Cores[0].Client

	ctx := context.Background()
	err := ds.StoreSecret(ctx, "id", "secret")
	assert.Equal(t, nil, err)

}
