package choicetypes

type LazySelection struct {
	user UserObj
	option Option
	magnitude int
}
type LazyArg
const (
	AgreeWillHelp LazyArg = iota,
	AgreeWontHelp,
	DisagreeNoAlt,
	DisagreeAlt,
)
func LazyCast(user UserObj, decision Decision, lazyarg LazyArg) {
	selection = LazySelection{}
	if lazyarg.AgreeWillHelp {
		selection.magnitude = 1
	}
	if lazyarg.DisagreeAlt {
		selection.magnitude = -1
	}
}

func LazyEval(decision Decision) DecisionResult {
	threshold := decision.threshold
	var Tally choice[count]
	var VolunteerPools [][]EntityObj
	for (vote in votes){
		if vote.agree or vote.volunteer {
			if vote.volunteer {
				VolunteerPools[vote.selection].append(vote.user)
			}
			Tally[vote.option] += vote.magnitude
		}

	}
}