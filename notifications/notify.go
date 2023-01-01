package notifications

import (
	"fmt"

	"github.com/preemptcordon/server/obj"
)

func notify(user obj.UserObj, message obj.Markdown) bool {
	options := user.Notifications()
	for preference := range options {
		fmt.Println(preference)
		// if preference {
		// 	preference(message)
		// 	return true
		// }
	}
	return false
}
