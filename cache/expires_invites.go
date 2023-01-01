package cache

import (
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/preemptcordon/server/obj"
)

func MakeInvite(user obj.UserObj) string {
	uuid := uuid.New()
	pipe := RedisClient.Pipeline()
	key := "invite" + uuid.String()
	pipe.Set(ctx, key, user.Email, time.Hour)
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return uuid.String()
}
func GetInvite(code string) (user string, err error) {
	val, err := RedisClient.Get(ctx, "invite"+code).Result()
	if err == redis.Nil {
		return "", err
	}
	return val, nil
}
