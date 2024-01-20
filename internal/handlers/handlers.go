package handlers

import (
	"github.com/mneumantic/wilkinson-innovations/internal/config"
	"github.com/mneumantic/wilkinson-innovations/internal/forms"
	"github.com/mneumantic/wilkinson-innovations/internal/models"
	"github.com/mneumantic/wilkinson-innovations/internal/render"
	"log"
	"net/http"
)

var Repo *Repository

// Repository is the respository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new respository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers, sets the new respository for the handlers.
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "index.gohtml", &models.Template{
		Title:      "Wilkinson Innovations",
		Production: m.App.Production,
		Other:      m.App.Testimonials,
		Products:   m.App.Products,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.gohtml", &models.Template{
		Title:      "About | Wilkinson Innovations",
		Production: m.App.Production,
	})
}

func (m *Repository) Order(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "order.gohtml", &models.Template{
		Title:      "Order | Wilkinson Innovations",
		Production: m.App.Production,
	})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	var emptyContact models.Contact
	data := make(map[string]interface{})
	data["contact"] = emptyContact

	render.RenderTemplate(w, r, "contact.gohtml", &models.Template{
		Title: "Contact Us | Wilkinson Innovations",
		Form:  forms.New(nil),
		Data:  data,
	})
}

func (m *Repository) PostContact(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	contact := models.Contact{
		Name:    r.Form.Get("name"),
		Company: r.Form.Get("company"),
		Email:   r.Form.Get("email"),
		Phone:   r.Form.Get("phone"),
		Message: r.Form.Get("message"),
	}

	form := forms.New(r.PostForm)

	form.Required("name", "email", "message")
	form.MinLength("name", 3, r)
	form.IsEmail("email")
	form.MinLength("message", 100, r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["contact"] = contact

		render.RenderTemplate(w, r, "contact.gohtml", &models.Template{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "contact", contact)
	http.Redirect(w, r, "/?contact-successful=true", http.StatusSeeOther)
}

func (m *Repository) Support(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "support.gohtml", &models.Template{
		Title:      "Support | Wilkinson Innovations",
		Production: m.App.Production,
	})
}

func (m *Repository) PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "privacy-policy.gohtml", &models.Template{
		Title:      "Privacy Policy | Wilkinson Innovations",
		Production: m.App.Production,
	})
}

func (m *Repository) TermsOfService(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "terms-of-service.gohtml", &models.Template{
		Title:      "Terms of Service | Wilkinson Innovations",
		Production: m.App.Production,
	})
}
