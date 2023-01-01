package mail
func background(){
	while time {
		for user in users {
			for email in user.mail {
				send(user,email)
			}
		}
	}
}