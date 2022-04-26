package main

import (
	"net/http"

	// github.com/bmizerany/pat
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nuttchai/WebApp-Golang/pkg/config"
	"github.com/nuttchai/WebApp-Golang/pkg/handlers"
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

	return mux
}
