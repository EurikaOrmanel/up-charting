package cache

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

type CacheConfig struct {
	RedisClient *redis.Client
}

var ctx = context.Background()

func Config() CacheConfig {
	host := os.Getenv("REDDIS_HOST")
	port := os.Getenv("REDDIS_PORT")
	url := host + ":" + port
	redisClient := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "",
		DB:       0,
	})
	return CacheConfig{redisClient}
}
