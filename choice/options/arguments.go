package options

/* arguments are per-option notes along a criteria (facet) */

type Option struct {
	parent     Decision
	maintainer EntityObj
	brief      Wiki
}

func AddArgument(user UserObj, their_option Option, criteria Criteria, argument Wiki) bool {
	var your_option Option
	for option in their_option.decision.options {
		if option.proposer == user {
			your_option = option
			break
		}
	}
	if your_option == nil {
		return false
	}
	their_option.criteria[criteria][your_option.key] = argument
	return true
}