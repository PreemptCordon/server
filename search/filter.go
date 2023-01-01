package search

func SearchFilter(thisuser UserObj, terms SearchTerms) UnrankedResults {
	var firstpass []UnrankedResults
	for t := range terms.typesallowed {
		matches := db.select(t,terms.query)
		for match in matches {
			if ! CheckTagDistance(match, terms.tags){
				continue
			}
			if CheckBlock(match.user, thisuser) {
				continue
			}
			if ! CheckACL(match.acl, thisuser) {
				continue
			}
			if CheckMute(match, thisuser){
				continue
			}
			if CheckModBlock(match) {
				transparency.blockedsearch.incr(match)
				continue
			}
			firstpass.append(match)
		}
	}
	return firstpass
}
