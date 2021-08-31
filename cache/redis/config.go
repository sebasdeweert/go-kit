package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// Config holds the Redis configuration.
type Config struct {
	Addresses  []string
	Password   string
	DB         int
	Expiration time.Duration
}

// GetRingOptions returns a redis.RingOptions based on the configuration parameters.
func (c *Config) GetRingOptions() *redis.RingOptions {
	addresses := make(map[string]string, len(c.Addresses))

	for i, address := range c.Addresses {
		addresses[fmt.Sprintf("server%d", i+1)] = address
	}

	return &redis.RingOptions{
		DB:       c.DB,
		Addrs:    addresses,
		Password: c.Password,
	}
}
