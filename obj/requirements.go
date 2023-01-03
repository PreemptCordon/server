package obj

type RequirementsObj struct {
	prompt     WikiArticle
	maintainer EntityInterface
	threshold  int // min number of petitioning voters to add an option to a decision
}
