package discover

import (
	"github.com/preemptcordon/server/obj"
)

func BuildDiscoverFeed(user obj.UserObj) {
	feed := BuildBaseFeed()
	settings = user.DiscoverSettings
	for entity := range feed {

	}
}
