package auth

func delete(user UserObj) {
	logout(user)
	db.users[user].delete()
}
