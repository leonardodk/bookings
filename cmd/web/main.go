package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/leonardodk/bookings/pkg/config"
	"github.com/leonardodk/bookings/pkg/handlers"
	"github.com/leonardodk/bookings/pkg/render"
)

const portNumber = ":8080"

// create a config var and a scs.session var available to the whole package "main"
var app config.AppConfig
var session *scs.SessionManager
func main() {


	// set this to true when in production
	app.InProduction = false
	// this sets up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	// pass the session to app config and thus everywhere else

	app.Session = session

	// populate the TC and put it in the config var
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("couldn't create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false // false to put in dev mode
	// takes a pointer to an Appconfig and puts it inside a repository and returns the repository
	repo := handlers.NewRepo(&app)

	// gives the repository type back to the handlers package
	handlers.NewHandlers(repo)

	// gives the appconfig to the render package
	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	//_ = http.ListenAndServe("127.0.0.1:8080", nil)
	// http.ListenAndServe("tcp4://localhost:8080", http.DefaultServeMux)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)

}
