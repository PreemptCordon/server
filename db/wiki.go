package db

import (
	"context"

	"github.com/preemptcordon/server/obj"
)

func LoadWiki(id string) obj.WikiArticle {
	var result obj.WikiArticle
	err := DBPool.QueryRow(context.Background(), "select from articles where ID = %s", id).Scan(&result)
	if err != nil {
		panic(err)
	}
	return result
}
