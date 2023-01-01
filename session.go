package main

func StartSession(user UserObj){
	longterm = uuid()
	sessionkey = "session:" + user.key + ":" + longterm
	RedisClient.Set(sessionkey, value)
	Cookie {"longtermclient", user.key + ":" + longterm}
}
func RefreshShortterm(request){
	shortterm = JWT({
		user: user.key,
		expires: shortttl
	})
	Cookie {"shorttermclient", shortterm}
}
func CheckSession(request){
	shortterm, err := loadcookie("shorttermclient")
	if shortterm != nil {
		if JWT.eval(shortterm) {
			return true
		}
	}
	longterm, err := loadcookie("longtermclient")
	if longterm != nil {
		result := RedisClient.get("session:" + longterm)
		if result != nil {
			RefreshShortterm()
		}
	}
	return false
}
func EndSession(user UserObj){
	RedisClient.del("Session:"+user.key+"*")
	cookie.del("longtermclient")
	cookie.del("shorttermclient")
}