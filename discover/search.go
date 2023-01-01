package discover

func DiscoverSearch(user UserObj, terms SearchTerms, options DiscoverOptions) {
	startresults := SearchFilter(user, terms)
	results ResultsPage
	for item in startresults {
		score = evalone(item)
		results.insert(item,score)
	}
	return results
}
