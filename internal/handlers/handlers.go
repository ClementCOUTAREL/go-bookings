package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ccoutarel/bookings/internal/config"
	"github.com/ccoutarel/bookings/internal/forms"
	"github.com/ccoutarel/bookings/internal/helpers"
	"github.com/ccoutarel/bookings/internal/models"
	"github.com/ccoutarel/bookings/internal/render"
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

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJson handles request for Availablitiy and sends JSON
func (m *Repository) AvailabilityJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	log.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// MakeReservation renders the make reservation page
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make_reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostMakeReservation handles the form send to make reservation
func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	// form.Has("first_name", r)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make_reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// MakeReservationJson handles request for Availablitiy and sends JSON
func (m *Repository) MakeReservationJson(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println("error creating JSON response")
	}

	log.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// ReservationSummary renders the reservation summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Can't get error from session")
		m.App.Session.Put(r.Context(), "error", "Cant get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation_summary.page.html", &models.TemplateData{
		Data: data,
	})
}
