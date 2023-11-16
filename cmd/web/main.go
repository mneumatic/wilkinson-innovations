package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"wilkinson-innovations/pkg/configs"
	"wilkinson-innovations/pkg/data"
	"wilkinson-innovations/pkg/handlers"
	"wilkinson-innovations/pkg/render"

	"github.com/alexedwards/scs/v2"
)

var app configs.AppConfig
var session *scs.SessionManager

func main() {
	// Set Production / Development
	app.Production = false

	// Set Sessions
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Production // Sets Cookie to Secure true in Production.

	app.Session = session
	app.Testimonials = data.TestimonialData()
	app.Products = data.ProductData()

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// Check if PORT is available.
	// If "false" set port
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	// Format port string
	addr := fmt.Sprintf(":%s", port)


	srv := &http.Server{
		Addr:         addr,
		Handler:      routes(&app),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("main: running server on port", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("main: couldn't start server: %v\n", err)
	}
}