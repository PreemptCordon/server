package wiki

import (
	"github.com/preemptcordon/server/cache"
	"github.com/preemptcordon/server/obj"
)

func linkdetect(text obj.Markdown) int {
	// return number of links in a string
	return 1
}
func UserLimitCheck(user obj.UserObj) bool {
	return true
}
func PageLimitCheck(page obj.WikiObj) bool {
	return false
}
func RateLimitEditCheck(user obj.UserObj, page obj.WikiObj) bool {

	userlimit := UserLimitCheck(user)
	pagelimit := PageLimitCheck(page)
	if userlimit {
		return false
	}
	if pagelimit {
		return false
	}
	return true
}
func UserLimitApply(user obj.UserObj, penalty int) {
	// apply limit to bucket
}

func RateLimitApply(user obj.UserObj, page obj.WikiObj, draft obj.Section) {
	linkpenalty := linkdetect(draft.Text)
	if cache.TryEdit(user, page) {
		page.Version = draft
	}
	UserLimitApply(user, linkpenalty)
}
