// Package pkg ...
package pkg

import "github.com/go-redis/redis/v8"

// Connection struct
type Connection struct {
	Pool *redis.Client
}

// NewConnection create redis pool connections
func NewConnection(redisServer string) *Connection {
	client := redis.NewClient(
		&redis.Options{
			Addr:     redisServer,
			Password: "", //
			PoolSize: 100,
		},
	)
	return &Connection{Pool: client}
}
