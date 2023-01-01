package cache

import (
	"regexp"

	"github.com/preemptcordon/server/obj"
)

func Sanitize(input obj.Markdown) obj.Markdown {
	htmlregex := regexp.MustCompile(`(?i)<[^>]*>`)
	output := string(htmlregex.ReplaceAll([]byte(input), []byte("")))
	return obj.Markdown(output)
}
func CachePage(key string, value obj.Markdown) {
	store := Sanitize(value)
	RedisClient.Set("cachepage"+key, string(store), 0) // will set expiration later...
}
func LoadPage(key string) obj.Markdown {
	stored, err := RedisClient.Get("cachepage" + key).Result()
	if err != nil {
		panic(err)
	}
	result := obj.Markdown(stored)
	return result
}
