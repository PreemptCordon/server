package cache

func TryEdit(user UserObj, page Wiki) bool {
	bucket, rate, err := RedisClient.Get("limit"+page.key).Result()
	recenteditors := ViewRecentEdits(page)
	if recenteditors.length > bucket {
		return false
	}
	pipe := RedisClient.Pipeline()
	thisedit = "limit"+page+timestamp
	pipe.Set(thisedit,user)
	pipe.Expire(thisedit,rate)

	return true
}
func SetLimit(page Wiki, bucket, rate) {
	RedisClient.set("limit"+page.key, {bucket, rate})
}
func ViewRecentEdits(page Wiki) []UserObj {
	recent := RedisClient.Get("limit"+page.key+"*")
	return recent
}