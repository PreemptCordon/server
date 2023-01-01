package transparency

import "github.com/preemptcordon/server/obj"

func (obj.UnrankedResult) PotentialVisit(obj.SearchTerms) {
	// increment the result's transparency counter, add the search terms used
}
