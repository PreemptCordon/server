package feed

func BuildBaseFeed(user UserObj) ResultsPage {
	events := user.following[:400]
	return feed
}
