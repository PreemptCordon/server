package discover

import (
	"github.com/preemptcordon/server/obj"
)

func DiscoverRank(user obj.UserObj, post obj.UnrankedResult) obj.ResultEntity {
	// options := user.Settings.Ranking.DiscoverSort
	var result obj.ResultEntity
	// for config := range options {
	// 	for user := range config.getusers() {
	// 		if user.reaction(post) {
	// 			result.Score += config.weight * config.direction
	// 		}
	// 	}
	// }
	// switch options.Scale {
	// case obj.BinaryWeight:
	// 	if result.Score < 0 {
	// 		resultpage.Results.Pop(post)
	// 	}
	// case obj.LinearWeight:
	// 	result.Score = result.Score
	// case obj.LogarithmicWeight:
	// 	result.Score = int(math.Log(float64(result.Score)))
	// case obj.ExponentialWeight:
	// 	result.Score = 2 ^ result.Score
	// }
	return result
}
