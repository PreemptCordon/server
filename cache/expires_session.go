package cache

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/preemptcordon/server/obj"
)

func StoreLongSession(user obj.UserObj) string {
	longterm := uuid.New()
	sessionkey := "session:" + user.Handle + ":" + longterm.String()
	RedisClient.Set(ctx, sessionkey, user.ID, time.Duration(30*24*time.Hour))

	return sessionkey
}
func CheckLongSession(sessionkey string) *obj.UserObj {
	result, err := RedisClient.Get(ctx, sessionkey).Result()
	if err != nil {
		panic(err)
	}
	user := obj.UserObj{
		ID: result,
	}
	return &user
}
func EndLongSession(sessionkey string) {
	// the trim is to prevent unauthenticated deleting of any session...
	RedisClient.Del(ctx, "session:"+strings.TrimPrefix("session:", sessionkey))
}
