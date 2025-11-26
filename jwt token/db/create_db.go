package db

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func CreateRedisDB(ctx context.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// defer rdb.Close()

	return rdb
}
