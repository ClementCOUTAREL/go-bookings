package handlers

import (
	"fmt"
	"net/http"

	"github.com/ccoutarel/bookings/pkg/config"
	"github.com/ccoutarel/bookings/pkg/models"
	"github.com/ccoutarel/bookings/pkg/render"
)

// Repositroy is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repositroy used by the handlers
var Repo *Repository

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home renders the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

// About renders the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})
}

// Colonels renders the colonels room page
func (m *Repository) Colonels(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "colonels_quarters.page.html", &models.TemplateData{})
}

// Majors renders the majors room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors_suite.page.html", &models.TemplateData{})
}

// Contact renders the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

// Reservation renders the check reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "reservation.page.html", &models.TemplateData{})
}

// PostReservation handles the form send to search availability
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

// MakeReservation renders the make reservation page
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make_reservation.page.html", &models.TemplateData{})
}

// PostMakeReservation handles the form send to make reservation
func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Posted to make reservationy"))
}
