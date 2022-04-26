package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// NOTE: "next" is a common name for middleware argument
// WriteToConsole is a middleware example function
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		/* NOTE: next is something "moving to the next..."
		It might move to next middleware or another part of the file where we return mux */
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/", // NOTE: "/" applies to ENTIRE SITE for a cookie
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
