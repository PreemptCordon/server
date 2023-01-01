package cache

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/preemptcordon/server/obj"
)

func MakeInvite(user obj.UserObj) string {
	uuid := uuid.New()
	pipe := RedisClient.Pipeline()
	key := "invite" + uuid.String()
	pipe.Set(key, user.Email, time.Hour)
	_, err := pipe.Exec()
	if err != nil {
		panic(err)
	}
	return uuid.String()
}
func GetInvite(code string) (user string, err error) {
	val, err := RedisClient.Get("invite" + code).Result()
	if err == redis.Nil {
		return "", err
	}
	return val, nil
}
