package choicetypes

func ResolveThresholdVote(decision Decision) DecisionResult {
	required = decision.threshold
	have := 0
	result := DecisionResult {
		decision,
		false
	}
	for (vote in votes) {
		if vote.for {
			have += 1
		}
		if have > threshold {
			result.approved = true
			return result
		}
	}
	return result
}