package discover

import (
	"math"

	"github.com/preemptcordon/server/obj"
)

func evalone(user obj.UserObj, post obj.SearchResult) obj.SearchResult {
	options := user.DiscoverOptions
	result := obj.SearchResult
	resultpage := obj.ResultPage{}
	for config := range options {
		for user := range config.getusers() {
			if user.reaction(post) {
				result.score += config.weight * config.direction
			}
		}
	}
	switch options.scale {
	case obj.BinaryWeight:
		if result < 0 {
			resultpage.Results.Pop(post)
		}
	case obj.LinearWeight:
		result = result
	case obj.LogarithmicWeight:
		result = math.Log(result)
	case obj.ExponentialWeight:
		result = 2 ^ result
	}
	return result
}
