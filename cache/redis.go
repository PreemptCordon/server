package cache

import (
	"github.com/go-redis/redis"
)

RedisClient := redis.NewClient(&redis.Options{
	Addr: ServerConfig.RedisURI,
	DB: 0,
})
var ctx = context.Background()