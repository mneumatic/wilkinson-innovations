package main

import (
	"net/http"
	"wilkinson-innovations/pkg/configs"
	"wilkinson-innovations/pkg/handlers"
	"wilkinson-innovations/pkg/models"
	"wilkinson-innovations/pkg/render"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *configs.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Post("/", handlers.Repo.PostHome)
	mux.Post("/mailling-list", handlers.Repo.PostFootbar)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/order", handlers.Repo.Order)
	mux.Get("/contact", handlers.Repo.Contact)
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
