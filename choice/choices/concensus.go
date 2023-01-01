package choicetypes

func ConcensusEval(decision Decision) DecisionResult {
	result := DecisionResult{
		decision,
		approved=false,
	}
	for (vote in votes){
		if vote.block {
			return result
		}
	}
	result.approved = true
	return result
}