package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// NOTE: this package will ONLY be imported by other parts of the application

/* IMPORTANT: we need to make sure that this package, config, is imported where it needs to be by other parts
but it DOESN'T import anything else from the application itself. It only uses things that are build into the standard library

bacause if we try to make this package import another package that's from other parts of code, they will import each other all over the place which causes a problem called IMPORT CYCLE and the app will not compile */

// NOTE: the reason that the application wide config is useful is because it will be available to every part of the application which can used to avoid IMPORT CYCLE Problem

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
