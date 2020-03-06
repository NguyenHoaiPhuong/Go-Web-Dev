package utils

import "github.com/go-redis/redis"

const (
	redisServer string = "localhost"
	redisPort   string = "32769"
)

var (
	client *redis.Client
)

func init() {
	// Init redis client
	client = redis.NewClient(&redis.Options{
		Addr:     redisServer + ":" + redisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// RedisClient returns redis client
func RedisClient() *redis.Client {
	return client
}
