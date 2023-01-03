package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/preemptcordon/server/cache"
	"github.com/preemptcordon/server/config"
	"github.com/preemptcordon/server/obj"
)

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func StartLongSessionHttp(user obj.UserObj, w http.ResponseWriter, r *http.Request) {
	sessionkey := cache.StoreLongSession(user)
	cookie := http.Cookie{
		Name:     "longterm",
		Value:    sessionkey,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
}
func RefreshShorttermHttp(user obj.UserObj, w http.ResponseWriter, r *http.Request) {
	claims := customClaims{
		Username: user.Handle,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    config.ServerConfig.Domain,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(config.ServerConfig.AppKey)
	if err != nil {
		panic(err)
	}
	cookie := http.Cookie{
		Name:     "shortterm",
		Value:    signedToken,
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
}
func CheckShorttermCookie(shortterm string) (*obj.UserObj, error) {
	var user obj.UserObj
	token, err := jwt.ParseWithClaims(
		shortterm,
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return config.ServerConfig.AppKey, nil
		},
	)
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return &user, errors.New("failed validating")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return &user, errors.New("session expired")
	}
	username := claims.Username
	user = obj.UserObj{
		Handle: username,
	}
	return &user, nil

}
func CheckSessionHttp(w http.ResponseWriter, r http.Request) (*obj.UserObj, error) {
	// user :=obj.UserObj{}
	var user *obj.UserObj
	shortterm, err := r.Cookie("shortterm")
	if err != nil {
		panic(err)
	}
	longterm, err := r.Cookie("longterm")
	if err != nil {
		panic(err)
	}
	if shortterm == nil && longterm == nil {
		return user, err
	}
	if shortterm != nil {
		user, err = CheckShorttermCookie(shortterm.String())
	}

	if user == nil && longterm != nil {
		user = cache.CheckLongSession(longterm.String())
		if &user != nil {
			RefreshShorttermHttp(*user, w, &r)
		}
	}
	return user, nil
}
func EndSession(user obj.UserObj, w http.ResponseWriter, r http.Request) {
	longterm, err := r.Cookie("longterm")
	if err != nil {
		panic(err)
	}
	cache.EndLongSession(longterm.String())
	shortcookie := http.Cookie{
		Name:     "shortterm",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &shortcookie)
	longcookie := http.Cookie{
		Name:     "longterm",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &longcookie)
}
