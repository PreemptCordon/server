package search

import "github.com/preemptcordon/server/obj"

func CheckACL(article obj.WikiArticle, you obj.UserObj) bool {
	article.ACL
	return false
}
func CheckBlock(author obj.UserObj, you obj.UserObj) bool {
	if DBPool.QueryRow(Context.Background(), "select from blockrelation where (origin = %you and destination = %them) or (origin = %them and destination = %you)") {
		return false
	}
	return true
}
func CheckMute(page obj.WikiArticle, you obj.UserObj) bool {
	tags := page.Tags
	author := page.Controller
	return true
}
func CheckTagDistance(page obj.WikiArticle, category []obj.Category, distance int) bool {
	tags := page.Tags
	categories := page.Category
	// heck this is a graph depth search
	return true
}

func CheckModBlock(page obj.WikiArticle) bool {
	return true
}
