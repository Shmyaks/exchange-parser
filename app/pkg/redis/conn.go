// Package redis ...
package redis

import "github.com/go-redis/redis"

// Connection struct
type Connection struct {
	Pool *redis.Client
}

// NewConnection create redis pool connections
func NewConnection(redisServer string) *Connection {
	client := redis.NewClient(&redis.Options{
		Addr:     redisServer,
		Password: "", //
		PoolSize: 100,
	})
	err := client.Ping().Err()
	if client.Ping().Err() != nil {
		panic(err)
	}
	return &Connection{Pool: client}
}
