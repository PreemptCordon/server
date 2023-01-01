package auth

import "github.com/preemptcordon/server/obj"

func delete(user obj.UserObj) {
	logout(user)
	db.users[user].delete()
}
