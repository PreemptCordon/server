package search

import (
	"github.com/preemptcordon/server/obj"
)

func SearchRank(user obj.UserObj, unsorted obj.UnrankedResult) obj.ResultPage {
	ranking := user.Settings.Ranking.SearchSort
	var results obj.ResultPage
	switch ranking {
	case obj.DiscoverSort:
		results = ranking.DiscoverSearch(user, unsorted)
	case obj.FeedSort:
		results = ranking.FeedRank(unsorted)
	case obj.CommentSort:
		results = ranking.CommentRank(unsorted)
	case obj.ReverseChron:
		results = ranking.ReverseChronRank(unsorted)
	}
	return results
}
