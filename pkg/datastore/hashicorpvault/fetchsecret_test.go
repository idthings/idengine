package hashicorpvault

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFetchSecret(t *testing.T) {

	const testComment = "Test fetching secret in Vault"

	cluster := createVaultTestCluster(t)
	defer cluster.Cleanup()

	ds := Datastore{}
	ds.client = cluster.Cores[0].Client

	ctx := context.Background()
	err := ds.StoreSecret(ctx, "id", "secret12345")
	assert.Equal(t, nil, err)

	response, err := ds.FetchSecret(ctx, "id")
	assert.Equal(t, nil, err)
	assert.Equal(t, "secret12345", response)

	// test when there is no id
	response, err = ds.FetchSecret(ctx, "id-fake")
	assert.Equal(t, nil, err)
	assert.Equal(t, "", response)

}
