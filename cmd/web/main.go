package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mrramalho/rmwebsite/config"
	"github.com/mrramalho/rmwebsite/pkg/handlers"
	"github.com/mrramalho/rmwebsite/pkg/render"
)

const webPort = ":8087"
var app config.AppConfig

func main() {
	// change this to true when in production
	app.InProduction = false
	app.UseCache = false
	// create session with scs
	app.Session = scs.New()
	app.Session.Lifetime = 24 * time.Hour
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = app.InProduction

	// create templates
	tc, err := render.CreateTemplates()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc

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
