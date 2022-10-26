package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	UseCache      bool
	InProduction  bool
	Session       *scs.SessionManager
}
