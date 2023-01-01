package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/preemptcordon/server/obj"
)

func MakeRecovery(ctx context.Context, user obj.UserObj) string {
	uuid := uuid.New()
	key := "recovery" + uuid.String()
	_, err := RedisClient.Set(ctx, key, user.Email, time.Hour).Result()
	if err != nil {
		panic(err)
	}
	return uuid.String()
}
func GetRecovery(code string, user string) (invite string, err error) {
	val, err := RedisClient.Get(ctx, "recovery"+code).Result()
	if err == redis.Nil {
		return "", errors.New("no code")
	}
	return val, nil
}
