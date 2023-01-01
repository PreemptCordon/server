package wiki

func linkdetect(string) int {
	// return number of links in a string
}

func RateLimitEditCheck(user UserObj, page Wiki){

	userlimit = UserLimitCheck(user)
	pagelimit = PageLimitCheck(page)
	if userlimit
}
func RateLimitApply(user UserObj, page Wiki, draft){
	linkpenalty = linkdetect(draft)
	if TryEdit(user, page) {
		wiki.version = draft
	}
	UserLimitApply(linkpenalty)
}