package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var handler testingHandler

	h := NoSurf(&handler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler. Type is: %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var handler testingHandler

	h := SessionLoad(&handler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler. Type is: %T", v))
	}
}
