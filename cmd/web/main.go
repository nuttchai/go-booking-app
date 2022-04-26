package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/nuttchai/WebApp-Golang/pkg/config"
	"github.com/nuttchai/WebApp-Golang/pkg/handlers"
	"github.com/nuttchai/WebApp-Golang/pkg/render"
)

// NOTE: It is convention that web folder is typically where Web Applications offten have their main function

/* NOTE: go.mod file tells the compiler that the application uses go modules.
it's like the package.json file used in Node.js dependency management */
const portNumber string = ":8080"

// NOTE: If we put variables here, it will be available for entire "main" package
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// change this to true when we are in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true                  // NOTE: should cookie stay available/persist after we close the browser?
	session.Cookie.SameSite = http.SameSiteLaxMode // NOTE: it tells how strict what site that cookie applies to
	session.Cookie.Secure = app.InProduction       // NOTE: set it to true, it will insist cookies need to be encrypted (in development mode, we use localhost:8080 which is not encrypted connection)

	// NOTE: make it available everywhere
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	_, _ = fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	/* NOTE Old Version:
	ListenAndServe: listen HTTP with port and handler (nil)
	It returns error, but we ignore it in this case by putting _ as variable */
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
