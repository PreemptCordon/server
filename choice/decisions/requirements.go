package decisions

/* requirements are option templates */

type RequirementsObj struct {
	parent Decision,
	prompt Wiki,
	maintainer EntityObj,
	threshold int, // min number of petitioning voters to add an option to a decision
}