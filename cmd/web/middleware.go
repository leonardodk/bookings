package main

import (
	// "fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// func WriteToConsol(next http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
// 		fmt.Println("Hit the page")
// 		next.ServeHTTP(w, r)
// 	})
// }


//NoSurf ads CSRF prptection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.Inproduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad saves and loads the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}