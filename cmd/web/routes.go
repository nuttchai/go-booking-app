package main

import (
	"net/http"

	// github.com/bmizerany/pat
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nuttchai/go-booking-app/pkg/config"
	"github.com/nuttchai/go-booking-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// NOTE: pat usage
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// NOTE: To make the complier know where to handle static items from local, like images
	// /static is the given directory
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
