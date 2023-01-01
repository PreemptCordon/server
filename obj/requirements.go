package obj

type RequirementsObj struct {
	prompt     WikiArticle
	maintainer EntityObj
	threshold  int // min number of petitioning voters to add an option to a decision
}
