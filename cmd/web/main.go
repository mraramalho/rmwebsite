package main

import (
	"log"
	"net/http"

	"github.com/mrramalho/rmwebsite/config"
	"github.com/mrramalho/rmwebsite/pkg/handlers"
	"github.com/mrramalho/rmwebsite/pkg/render"
)

const webPort = "8087"

func main() {
	var app config.AppConfig

	// create templates
	tc, err := render.CreateTemplates()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.HomeHandler)
	http.HandleFunc("/about", handlers.Repo.AboutHandler)
	http.HandleFunc("/contact", handlers.Repo.ContactHandler)

	log.Println("Starting server on port", webPort)
	err = http.ListenAndServe(":8087", nil)
	if err != nil {
		log.Fatal(err)
	}
}
