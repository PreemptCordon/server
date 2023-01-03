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
	opt, err := redis.ParseURL(redisuri)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
	RedisClient = client

}
