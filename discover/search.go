package discover

import (
	"github.com/preemptcordon/server/obj"
)

func DiscoverSearch(user obj.UserObj, terms SearchTerms, options obj.DiscoverOptions) {
	startresults := SearchFilter(user, terms)
	results := obj.ResultsPage
	for item := range startresults {
		score = evalone(item)
		results.insert(item, score)
	}
	return results
}
