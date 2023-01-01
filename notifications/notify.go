package notifications

import (
	_ "github.com/preemptcordon/server/obj"
)

func notify(user UserObject, message Markdown) bool {
	options := user.Notifications
	for i, preference := range options {
		if preference {
			preference(message)
			return true
		}
	}
	return false
}
