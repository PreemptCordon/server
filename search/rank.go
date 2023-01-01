package search

func SearchRank(user UserObj, unsorted UnrankedResults) ResultsPage {
	ranking = user.SearchRank
	var results ResultsPage
	switch ranking {
		case discover {
			results = DiscoverSearch(user, unsorted)
		}
		case feedrank {
			results = FeedRank(unsorted)
		}
		case commentrank {
			results = CommentRank(unsorted)
		}
		case reversechron {
			results = ReverseChronRank(unsorted)
		}
	}
	return results
}
