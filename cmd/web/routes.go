package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/leonardodk/bookings/pkg/config"
	"github.com/leonardodk/bookings/pkg/handlers"
)

// this is a function that handles our routing to the correct handler
// we make a new patternsrevermux (basically a struct with a map of handlers matched with strings)
// add our handlers to it, then return it.

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
