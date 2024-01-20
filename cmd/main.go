package main

import (
	"encoding/gob"
	"fmt"
	"github.com/mneumantic/wilkinson-innovations/internal/config"
	"github.com/mneumantic/wilkinson-innovations/internal/data"
	"github.com/mneumantic/wilkinson-innovations/internal/handlers"
	"github.com/mneumantic/wilkinson-innovations/internal/models"
	"github.com/mneumantic/wilkinson-innovations/internal/render"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	// Check if PORT is available.
	// If "false" set port
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	// Format port string
	addr := fmt.Sprintf("localhost:%s", port)

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

func run() error {
	// Session
	gob.Register(models.Contact{})

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
	return nil
}
