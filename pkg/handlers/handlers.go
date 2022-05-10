package handlers

import (
	"net/http"

	"github.com/Shreeyash-Naik/Hotel-Booking/pkg/config"
	"github.com/Shreeyash-Naik/Hotel-Booking/pkg/models"
	"github.com/Shreeyash-Naik/Hotel-Booking/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(&w, "home.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.Get(r.Context(), "remote_ip")

	data := map[string]interface{}{
		"test":      "Test for about",
		"remote_ip": remoteIP,
	}

	td := models.TemplateData{
		Data: data,
	}
	render.RenderTemplate(&w, "about.html", &td)
}
