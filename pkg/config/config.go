package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger // allows us to write to log
	Inproduction  bool
	Session       *scs.SessionManager
}
