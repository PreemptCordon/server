package cache

import (
	"time"

	"github.com/preemptcordon/server/obj"
)

type LimitData struct {
	Bucket int           `redis:"bucket"`
	Rate   time.Duration `redis:"rate"`
}

func TryEdit(user obj.UserObj, page obj.WikiArticle) bool {
	var limitinfo LimitData
	hashset := RedisClient.HGetAll(ctx, "limit"+page.Key.String())
	err := hashset.Scan(&limitinfo)
	if err != nil {
		panic(err)
	}
	recenteditors := ViewRecentEdits(page)
	if len(recenteditors) > limitinfo.Bucket {
		return false
	}
	thisedit := "limit" + page.Key.String() + time.Now().String()
	RedisClient.Set(ctx, thisedit, user.ID, limitinfo.Rate)

	return true
}
func SetLimit(page obj.WikiArticle, bucket, rate int) {
	RedisClient.HSet(ctx, "limit"+page.Key.String(), "bucket", bucket)
	RedisClient.HSet(ctx, "limit"+page.Key.String(), "rate", rate)
}
func ViewRecentEdits(page obj.WikiArticle) []obj.UserObj {
	recenteditors, err := RedisClient.Get(ctx, "limit"+page.Key.String()+"*").Result()
	if err != nil {
		panic(err)
	}
	var result []obj.UserObj
	user := new(obj.UserObj)
	user.ID = string(recenteditors) // technically this is an int... I'll change it later.
	result = append(result, *user)
	// for editor := range recenteditors {
	// 	user := new(obj.UserObj)
	// 	user.ID = string(editor) // technically this is an int... I'll change it later.
	// 	result = append(result, user)
	// }
	return result
}
