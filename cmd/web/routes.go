package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/mrramalho/rmwebsite/config"
	"github.com/mrramalho/rmwebsite/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.Get("/", handlers.Repo.HomeHandler)
	router.Get("/about", handlers.Repo.AboutHandler)

	return router
}
