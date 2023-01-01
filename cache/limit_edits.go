package cache

import (
	"time"

	"github.com/preemptcordon/server/obj"
)

type LimitData struct {
	Bucket int           `redis:"bucket"`
	Rate   time.Duration `redis:"rate"`
}

func TryEdit(user obj.UserObj, page obj.WikiObj) bool {
	var limitinfo LimitData
	if err := RedisClient.HGetAll("limit" + page.Key.String()).Scan(&limitinfo); err != nil {
		panic(err)
	}
	recenteditors := ViewRecentEdits(page)
	if len(recenteditors) > limitinfo.Bucket {
		return false
	}
	thisedit := "limit" + page.Key.String() + time.Now().String()
	RedisClient.Set(thisedit, user.ID, limitinfo.Rate)

	return true
}
func SetLimit(page obj.WikiObj, bucket, rate int) {
	RedisClient.HSet("limit"+page.Key.String(), "bucket", bucket)
	RedisClient.HSet("limit"+page.Key.String(), "rate", rate)
}
func ViewRecentEdits(page obj.WikiObj) []obj.UserObj {
	recenteditors, err := RedisClient.Get("limit" + page.Key.String() + "*").Result()
	if err != nil {
		panic(err)
	}
	var result []obj.UserObj
	user := new(obj.UserObj)
	user.ID = string(recenteditors) // technically this is an int... I'll change it later.
	result = []obj.UserObj{
		user,
	}
	// for editor := range recenteditors {
	// 	user := new(obj.UserObj)
	// 	user.ID = string(editor) // technically this is an int... I'll change it later.
	// 	result = append(result, user)
	// }
	return result
}
