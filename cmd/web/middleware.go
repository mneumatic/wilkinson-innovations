package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST Request
func NoSurf(next http.Handler) http.Handler {
	csrfHanlder := nosurf.New(next)

	csrfHanlder.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.Production,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHanlder
}

// SessionLoad loads and saves the session on evey request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}