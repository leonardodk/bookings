package config

import (
	"html/template"
	"log"
	"github.com/alexedwards/scs/v2"
)

// holds teh application config, all site wide settings
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
