package funcs

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetUserLevel(ctx context.Context, token string, rdb *redis.Client) (string, error) {
	val, err := rdb.HGet(ctx, "auth:"+token, "user_level").Result()

	if err == redis.Nil {
		return "", fmt.Errorf("token is not found")
	}
	if err != nil {
		return "", err
	}

	return val, nil
}

func RemoveToken(ctx context.Context, token string, rdb *redis.Client) error {
	return rdb.Del(ctx, "auth:"+token).Err()
}
