package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func connectRedis(ctx context.Context, redisuri string) {
	client := redis.NewClient(&redis.Options{
		Addr: redisuri,
		DB:   0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
	RedisClient = client

}
