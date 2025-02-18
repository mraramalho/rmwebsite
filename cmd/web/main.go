package main

import (
	"log"
	"net/http"

	"github.com/mrramalho/rmwebsite/config"
	"github.com/mrramalho/rmwebsite/pkg/handlers"
	"github.com/mrramalho/rmwebsite/pkg/render"
)

const webPort = ":8087" // Add the colon before the port number

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

	log.Println("Starting server on port", webPort)
	srv := &http.Server{
		Addr:    webPort,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
