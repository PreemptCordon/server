package auth

func RecoverStart(email string) {
	user = db.user[email]
	if user == nil {
		return false
	}
	recoverurl = cache.recover(user)
	send(user.email, "please visit this URL: "+recoverurl)
}
func RecoverFinalize(uri) {
	if cache.checkrecover(uri) {
		DisableTOTP(user)
		DisableWebAuthn(user)
		db.users.locked = false
		db.user.password = password
	}
	return false
}
func RecoverNotme(uri) {

}
