package discover

import (
	"fmt"

	"github.com/preemptcordon/server/obj"
)

func BuildBaseFeed() []obj.UnrankedResult {
	var result []obj.UnrankedResult
	// pull from some other feed sort, without circular imports...
	return result
}
func BuildDiscoverFeed(user obj.UserObj) []obj.ResultPage {
	var result []obj.ResultPage
	feed := BuildBaseFeed()
	for _, entity := range feed {
		fmt.Println(entity)
		// if entity.Author.Settings.Discovery.Traversable {
		// 	// we can use this author to look for other authors

		// 	// somehow find them, then:
		// 	resultuser := entity.Author
		// 	if resultuser.Settings.Discovery.Discoverable {
		// 		// we can append this user to results
		// 	}
		// }

	}
	return result
}
