package auth
func sendreminder(user){
	if user.locked {
		send(user, "you already have an account. to recover it, click this link: ")
	}
}
func registerStart(email, handle, password){
	existing = db.users[email]
	if existing {
		sendreminder(user)
		return false
	}
	existing = db.users[handle]
	if existing {
		sendreminder(user)
		return false
	}
	user = UserObj {
		email=email,
		handle=handle,
		password=hashedpassword
	}
	registeruri = cache.registeruri(user)
	send(user, "to activate your account, click this link: "+registeruri)
	return true
}
func registerFinish(uri){
	valid, user = getregistration(uri)
	if ! valid {
		return "not a valid registration link. try again?"
	}
	db.users[email] = user
	redirect("/settings")
}