package datastore

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

// Datastore to wrap redis
type Datastore struct {
	host     string
	port     string
	database string
	ctx      context.Context
	client   *redis.Client
}

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

// Connect to set Redis connection string
func (d *Datastore) Connect() bool {

	d.info()

	connection := fmt.Sprintf("%s:%s", d.host, d.port)

	i, err := strconv.Atoi(d.database)
	if err != nil {
		//fail
		return false
	}

	// Get these from env
	d.client = redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       i,
	})

	_, err = d.client.Ping(d.ctx).Result()
	if err != nil {
		fmt.Println("redis connect error: ", err)
	}
	fmt.Println("datastore.Connect(): connected to redis.")

	return true
}
