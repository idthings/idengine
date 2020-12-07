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
	envPort     string
	envToken    string
	expectHost  string
	expectToken string
}{
	{
		comment:     "test defaults, no env vars set",
		envHost:     "",
		envPort:     "",
		envToken:    "",
		expectHost:  "http://127.0.0.1",
		expectToken: "root",
	},
	{
		comment:     "test host set from env",
		envHost:     "myhost1",
		envPort:     "",
		envToken:    "",
		expectHost:  "myhost1",
		expectToken: "root",
	},
	{
		comment:     "test everything set from env",
		envHost:     "myvaulthost",
		envPort:     "9999",
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
