package main

import (
	"github.com/mneumantic/wilkinson-innovations/internal/config"
	"github.com/mneumantic/wilkinson-innovations/internal/handlers"
	"github.com/mneumantic/wilkinson-innovations/internal/models"
	"github.com/mneumantic/wilkinson-innovations/internal/render"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/order", handlers.Repo.Order)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Post("/contact", handlers.Repo.PostContact)
	mux.Get("/support", handlers.Repo.Support)
	mux.Get("/privacy-policy", handlers.Repo.PrivacyPolicy)
	mux.Get("/terms-of-service", handlers.Repo.TermsOfService)

	//  CHI NotFound & MethodNotAllowed ERROR HANDLING
	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.RenderTemplate(w, r, "error.gohtml", &models.Template{
			Title:      "Page Not Found | Wilkinson Innovations",
			Production: app.Production,
		})
		w.WriteHeader(404)
	})

	mux.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		render.RenderTemplate(w, r, "error.gohtml", &models.Template{
			Title:      "Page Not Found | Wilkinson Innovations",
			Production: app.Production,
		})
		w.WriteHeader(405)
	})

	fileServer := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return mux
}
