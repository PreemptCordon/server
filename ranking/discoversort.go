package ranking

import (
	"github.com/preemptcordon/server/obj"
	"github.com/preemptcordon/server/search"
)

func DiscoverSearch(user obj.UserObj, terms obj.SearchTerms, options obj.DiscoverOptions) []obj.ResultPage {
	startresults := search.SearchFilter(user, terms)
	results := []obj.ResultPage{}
	for item := range startresults {
		item.rankself()
		results = append(results, item)
	}
	return results
}
