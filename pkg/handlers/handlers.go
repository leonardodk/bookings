package handlers

// this file just contains handlers, the things that control what the user sees
// as long as this is in package main we can just run main.go from the TLD and we shouod be able to fnd everything

import (
	"net/http"

	"github.com/leonardodk/bookings/pkg/config"
	"github.com/leonardodk/bookings/pkg/models"
	"github.com/leonardodk/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// takes a pointer to an Appconfig and puts it inside a repository and returns the repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// takes a pointer to the repository made in main sets it as Repo wiht appconfig inside
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

// parsedTemplate, _ := template.ParseFiles("./templates/about.page.tmpl")

// 	err := parsedTemplate.Execute(w, nil)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
