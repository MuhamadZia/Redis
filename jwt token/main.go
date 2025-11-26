package main

import (
	"context"

	"jwt-token-redis/db"
	"jwt-token-redis/funcs"
)

var ctx = context.Background()

func main() {
	// create rdb
	rdb := db.CreateRedisDB(ctx)

	// create sample
	db.CreateSample(ctx, rdb)

	// get test 1
	tokenTest := "jwt-token-abc123"
	userLvl1, err := funcs.GetUserLevel(ctx, tokenTest, rdb)
	if err != nil {
		panic(err)
	}

	println("user lvl:", userLvl1)

	// logout
	err = funcs.RemoveToken(ctx, tokenTest, rdb)
	if err != nil {
		panic(err)
	}

	// get test 1
	userLvl2, err := funcs.GetUserLevel(ctx, tokenTest, rdb)
	if err != nil {
		panic(err)
	}

	println("user lvl:", userLvl2)
}
