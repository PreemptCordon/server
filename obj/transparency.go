package obj

type TransparencyReport interface {
	PotentialVisitDenied(SearchTerms)
}

func (s *UnrankedResult) PotentialVisitDenied(t SearchTerms) {
	//increment visit count
}
