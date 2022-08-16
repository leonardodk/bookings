package handlers

import (
	"github.com/leonardodk/bookings/pkg/config"
	"github.com/leonardodk/bookings/pkg/models"
	"github.com/leonardodk/bookings/pkg/render"
	"net/http"
)

// is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repositry used by the handlers
var Repo *Repository

// NewRepo creates a new respository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers to Repo
func NewHandlers(r *Repository) {
	Repo = r
}

// all this malarkey is so that the handlers functions can access the stuff in
// appconfig easily

// all handlers go here

// this is the Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	// send data to template
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// this is the About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, can you hear me?"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
