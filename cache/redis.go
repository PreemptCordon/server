package cache

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

var RedisClient *redis.Client
var ctx context.Context

func ConnectRedis(redisuri string) {
	ctx = context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: redisuri,
		DB:   0,
	})
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
	RedisClient = client

}
