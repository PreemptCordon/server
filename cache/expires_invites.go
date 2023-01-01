package cache

func MakeInvite(user UserObj) {
	uuid := uuid4()
	pipe := RedisClient.Pipeline()
	key := "invite"+uuid
	err := pipe.Set(ctx, key, user.Email).Result()
	pipe.Expire(ctx,)key, time.Hour)
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return uuid
}
func GetInvite(code string) (user string) {
	val, err := RedisClient.Get(ctx, "invite"+code).Result()
	if err == redis.Nil {
		return false
	}
	return val
}
