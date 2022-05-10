package main

import (
	"net/http"

	"github.com/Shreeyash-Naik/Hotel-Booking/pkg/config"
	"github.com/Shreeyash-Naik/Hotel-Booking/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/home", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
