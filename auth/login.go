package auth

import (
	"net/http"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	user = r.data.user
	user := db.users[user]
	if user == nil {
		return false
	}
	password_match = hashpasswd(user, r.data.password)
	if password_match == false {
		db.users[user].failures += 1
		if db.users[user].failures > 20 {
			user.locked = true
			recoverurl = cache.startrecover()
			send(user, "your account has been locked due to too many password attempts. To recover, click this link: "+recoverurl)
		}
		return false
	}
	token := MakeToken(user)
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		Name:     "shortwt",
		Value:    token,
	})
}
