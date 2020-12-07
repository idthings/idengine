package hashicorpvault

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	HostEnvVar  = "IDENGINE_VAULT_HOST"
	TokenEnvVar = "IDENGINE_VAULT_TOKEN"
)

var TestInfoItems = []struct {
	comment     string
	envHost     string
	envToken    string
	expectHost  string
	expectToken string
}{
	{
		comment:     "test defaults, no env vars set",
		envHost:     "",
		envToken:    "",
		expectHost:  "http://127.0.0.1:8200",
		expectToken: "developer",
	},
	{
		comment:     "test host set from env",
		envHost:     "myhost1",
		envToken:    "",
		expectHost:  "myhost1",
		expectToken: "developer",
	},
	{
		comment:     "test everything set from env",
		envHost:     "myvaulthost",
		envToken:    "sometoken",
		expectHost:  "myvaulthost",
		expectToken: "sometoken",
	},
}

func TestInfo(t *testing.T) {

	os.Unsetenv(HostEnvVar)
	os.Unsetenv(TokenEnvVar)

	for _, item := range TestInfoItems {

		if len(item.envHost) > 0 {
			os.Setenv(HostEnvVar, item.envHost)
			defer os.Unsetenv(HostEnvVar)
		}
		if len(item.envToken) > 0 {
			os.Setenv(TokenEnvVar, item.envToken)
			defer os.Unsetenv(TokenEnvVar)
		}

		var ds Datastore
		ds.info()

		assert.Equal(t, item.expectHost, ds.host, item.comment)
		assert.Equal(t, item.expectToken, ds.token, item.comment)

	}
}
