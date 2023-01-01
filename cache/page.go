package cache

func CachePage(key string, value Markdown) {
	store := Sanitize(value)
	RedisClient.Set("cachepage"+key, store)
}
func LoadPage(key string) Markdown {
	result, err := RedisClient.Get("cachepage" + key)
	return result
}
