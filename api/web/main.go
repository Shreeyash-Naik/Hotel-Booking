package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Shreeyash-Naik/Hotel-Booking/pkg/config"
	"github.com/Shreeyash-Naik/Hotel-Booking/pkg/handlers"
	"github.com/Shreeyash-Naik/Hotel-Booking/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var (
	app config.AppConfig
)

func main() {

	app.InProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.HttpOnly = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Cannot create template Cache")
	}

	app.UseCache = false
	app.TemplateCache = tc
	render.NewTemplates(&app)

	handlersRepo := handlers.NewRepo(&app)
	handlers.NewHandlers(handlersRepo)

	log.Println("App starting at port :8080")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: Routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
