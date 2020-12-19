package redis

import (
	"context"
	"os"
)

// Info sets our env vars
func (d *Datastore) info() {

	var value string

	value = os.Getenv("IDENGINE_DB_HOST")
	if len(value) > 0 {
		d.host = value
	} else {
		d.host = "127.0.0.1"
	}

	value = os.Getenv("IDENGINE_DB_PORT")
	if len(value) > 0 {
		d.port = value
	} else {
		d.port = "6379"
	}
	value = os.Getenv("IDENGINE_DB_DATABASE")
	if len(value) > 0 {
		d.database = value
	} else {
		d.database = "0"
	}

	d.ctx = context.Background()
}
