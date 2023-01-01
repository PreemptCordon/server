package options


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