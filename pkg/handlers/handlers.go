package handlers

import (
	"net"
	"net/http"

	"github.com/AMSteffensen/bookings/pkg/config"
	"github.com/AMSteffensen/bookings/pkg/models"
	"github.com/AMSteffensen/bookings/pkg/render"
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

// Home is the homepage handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		remoteIP = r.RemoteAddr
	}
	ip := net.ParseIP(remoteIP)
	if ip != nil && ip.IsLoopback() && ip.To4() == nil {
		remoteIP = "127.0.0.1"
	}
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello, again."

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "booking.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "search.tmpl", &models.TemplateData{})
}
