package handlers

import (
	"fmt"
	"net/http"
	"wilkinson-innovations/pkg/configs"
	"wilkinson-innovations/pkg/models"
	"wilkinson-innovations/pkg/render"
)

var Repo *Repository

// Repository is the respository type
type Repository struct {
	App *configs.AppConfig
}

// NewRepo creates a new respository
func NewRepo(a *configs.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers, sets the new respository for the handlers.
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "index.tmpl", &models.Template{
		Title: "Wilkinson Innovations",
		Production: m.App.Production,
		Other: m.App.Testimonials,
		Products: m.App.Products,
	})
}

func (m *Repository) PostHome(w http.ResponseWriter, r *http.Request) {
	formDetails := struct {
		Name    string
		Company string
		Email   string
		Phone   string
		Message string
	}{
		Name: r.FormValue("name"),
		Company: r.FormValue("company"),
		Email: r.FormValue("email"),
		Phone: r.FormValue("phone"),
		Message: r.FormValue("message"),
	}

	fmt.Printf("Name: %s\nCompany: %s\nEmail: %s\nPhone: %s\nMessage: %s\n", formDetails.Name, formDetails.Company, formDetails.Email, formDetails.Phone, formDetails.Message)
	http.Redirect(w, r, "/?contact-successful=true", http.StatusSeeOther)
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.tmpl", &models.Template{
		Title: "About | Wilkinson Innovations",
		Production: m.App.Production,
	})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.tmpl", &models.Template{
		Title: "Contact Us | Wilkinson Innovations",
		Production: m.App.Production,
	})
}

func (m *Repository) PostContact(w http.ResponseWriter, r *http.Request) {
	formDetails := struct {
		Name    string
		Company string
		Email   string
		Phone   string
		Message string
	}{
		Name: r.FormValue("name"),
		Company: r.FormValue("company"),
		Email: r.FormValue("email"),
		Phone: r.FormValue("phone"),
		Message: r.FormValue("message"),
	}

	fmt.Printf("Name: %s\nCompany: %s\nEmail: %s\nPhone: %s\nMessage: %s\n", formDetails.Name, formDetails.Company, formDetails.Email, formDetails.Phone, formDetails.Message)
	http.Redirect(w, r, "/?contact-successful=true", http.StatusSeeOther)
}

func (m *Repository) PostFootbar(w http.ResponseWriter, r *http.Request) {
	formDetails := struct {
		Name    string
		Email   string
	}{
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
	}

	fmt.Printf("Name: %s\nCompany: %s\nEmail:", formDetails.Name, formDetails.Email)
	http.Redirect(w, r, "/?mailling-list-successful=true", http.StatusSeeOther)
}

func (m *Repository) PrivacyPolicy(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "privacy-policy.tmpl", &models.Template{
		Title: "Privacy Policy | Wilkinson Innovations",
		Production: m.App.Production,
	})
}

func (m *Repository) TermsOfService(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "terms-of-service.tmpl", &models.Template{
		Title: "Terms of Service | Wilkinson Innovations",
		Production: m.App.Production,
	})
}