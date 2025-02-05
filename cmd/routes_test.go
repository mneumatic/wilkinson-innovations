package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/mneumantic/wilkinson-innovations/internal/config"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Error(fmt.Sprintf("Type is not *chi.Mux. Type is %T", v))
	}
}
