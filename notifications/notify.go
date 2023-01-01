package notifications

func notify(user UserObject, message Markdown) bool {
	options = user.Notifications
	for preference in options {
		if preference {
			preference(message)
			return true
		}
	}
	return false
}
