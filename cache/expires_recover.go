package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/preemptcordon/server/obj"
)

func MakeRecovery(ctx context.Context, user obj.UserObj) string {
	uuid := uuid.New()
	key := "recovery" + uuid.String()
	_, err := RedisClient.Set(key, user.Email, time.Hour).Result()
	if err != nil {
		panic(err)
	}
	return uuid.String()
}
func GetRecovery(code string, user string) (invite string, err error) {
	val, err := RedisClient.Get("recovery" + code).Result()
	if err == redis.Nil {
		return "", errors.New("no code")
	}
	return val, nil
}
