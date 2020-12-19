package redis

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	HostEnvVar = "IDENGINE_DB_HOST"
	PortEnvVar = "IDENGINE_DB_PORT"
	DBEnvVar   = "IDENGINE_DB_DATABASE"
)

var TestInitItems = []struct {
	comment        string
	envHost        string
	envPort        string
	envDatabase    string
	expectHost     string
	expectPort     string
	expectDatabase string
}{
	{
		comment:        "test defaults",
		envHost:        "",
		envPort:        "",
		envDatabase:    "",
		expectHost:     "127.0.0.1",
		expectPort:     "6379",
		expectDatabase: "0",
	},
	{
		comment:        "test host set from env",
		envHost:        "myhost1",
		envPort:        "",
		envDatabase:    "",
		expectHost:     "myhost1",
		expectPort:     "6379",
		expectDatabase: "0",
	},
	{
		comment:        "test everything set from env",
		envHost:        "myhost2",
		envPort:        "myport2",
		envDatabase:    "mydatabase2",
		expectHost:     "myhost2",
		expectPort:     "myport2",
		expectDatabase: "mydatabase2",
	},
}

func TestInfo(t *testing.T) {

	os.Unsetenv(HostEnvVar)
	os.Unsetenv(PortEnvVar)
	os.Unsetenv(DBEnvVar)

	for _, item := range TestInitItems {

		if len(item.envHost) > 0 {
			os.Setenv(HostEnvVar, item.envHost)
			defer os.Unsetenv(HostEnvVar)
		}
		if len(item.envPort) > 0 {
			os.Setenv(PortEnvVar, item.envPort)
			defer os.Unsetenv(PortEnvVar)
		}
		if len(item.envDatabase) > 0 {
			os.Setenv(DBEnvVar, item.envDatabase)
			defer os.Unsetenv(DBEnvVar)
		}

		var ds Datastore
		ds.info()

		assert.Equal(t, item.expectHost, ds.host, item.comment)
		assert.Equal(t, item.expectPort, ds.port, item.comment)
		assert.Equal(t, item.expectDatabase, ds.database, item.comment)

	}
}
