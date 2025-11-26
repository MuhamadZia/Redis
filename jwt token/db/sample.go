package db

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func CreateSample(ctx context.Context, rdb *redis.Client) {
	// sample
	token := "jwt-token-abc123"
	userLevel := "admin"

	// save token
	err := rdb.HSet(ctx, "auth:"+token, map[string]string{
		"user_level": userLevel,
	}).Err()
	if err != nil {
		panic(err)
	}

	rdb.Expire(ctx, "auth:"+token, 2*time.Minute)
}
