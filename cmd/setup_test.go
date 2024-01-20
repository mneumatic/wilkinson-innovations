package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

type testingHandler struct{}

func (mh *testingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
