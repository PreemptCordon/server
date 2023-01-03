package obj

type Argument struct {
	authorrole RoleEntityObj // who's arguing
	option     *Option       // what option we're talking about
	criteria   *Criteria     // what facet of the option are we evaluating
	statement  *Section      // what's our evaluation
}
