package main

import (
	"fmt"
	"github.com/leonardodk/bookings/pkg/config"
	"github.com/leonardodk/bookings/pkg/handlers"
	"github.com/leonardodk/bookings/pkg/render"
	"log"
	"net/http"
	"time"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// going to be a handler function
// in order for a funciton to respond to a request for a web browser/for
// it to be a handler, it needs  response writer, and a request.
// this is the Home page handler

func main() {


	// change to true when in production
	app.Inproduction = false

	session = scs.New() // refers to session defined outside of main and thus available to middleware.go
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Inproduction // usually want this to be true but this is currently running on local host

	app.Session = session 

	// create template cache
	tc, err := render.CreateTemplateCache()

	// error check
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// add template cache to app config variable (AVR)
	app.TemplateCache = tc

	// false = set to dev mode
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	// passes a pointer to appconfic over to render.go
	// this way render.go can be set using the appconfig created here in main.
	render.NewTemplate(&app)

	// handlers is the package, Repo the variable, Home the method

	fmt.Printf("Starting on port No: %s\n***\n", portNumber)

	// _ = http.ListenAndServe(PortNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
