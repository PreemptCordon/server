package options

type Option {
	decision Decision,
	authorRecommends int[Option],
	wiki Wiki,
}

func AddOption(user UserObj, decision Decision, about Wiki, alternatives int[Option]) bool {
	if (decision.debator != user){
		return false
	}
	decision.options[user] = Option{
		decision=decision,
		authorRecommends=alternatives,
		wiki=about
	}
}