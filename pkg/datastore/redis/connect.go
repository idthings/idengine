package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
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
		return false
	}

	fmt.Println("datastore.Connect(): connected to redis.")
	return true
}
