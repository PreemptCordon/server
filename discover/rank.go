
func evalone(user UserObj, post SearchResult){
	options = user.DiscoverOptions
	result = 0
	for config in options {
		for user in config.getusers(){
			if user.reaction(post) {
				result.score += config.weight * config.direction
			}
		}
	}
	switch (options.scale){
		case binary {
			if (result < 0){
				resultpage.pop(post)
			}
		}
		case linear {
			result = result
		}
		case logarithmic {
			result = log(result)
		}
		case exponential {
			result = 2^result
		}
	}
	return result
}