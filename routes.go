package main

import (
	"net/http"

	"github.com/preemptcordon/server/user"
	"github.com/preemptcordon/server/wiki"

	"github.com/go-chi/chi/v5"
)

func router() http.Handler {
	r := chi.NewRouter()
	// live JWTs
	r.Group(func(r chi.Router) {
		// r.Use(jwtauth.Verifier(tokenAuth))
		r.Get("/profile:*", user.ProfileHandler)
		r.Get("/article:*", wiki.ArticleHandler)
	})
	// refresh redis queue
	r.Group(func(r chi.Router) {
		// r.Use(redisauth.Verifier(tokenAuth))
		r.Get("/auth/refresh", func(w http.ResponseWriter, r *http.Request) {})
	})
	r.Group(func(r chi.Router) {
		r.Get("/rss/:feed", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/dav/:feed", func(w http.ResponseWriter, r *http.Request) {})
	})
	// unauth'd logins
	r.Group(func(r chi.Router) {
		r.Get("/auth/login", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/auth/register", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/auth/recover", func(w http.ResponseWriter, r *http.Request) {})
	})
	return r
}
