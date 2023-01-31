package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi(t *testing.T) {
	// Arrange
	r := chi.NewRouter()
	r.Get("/emailfinder/search/{term}", searchBodyByTerm)
	searchTerm := "apte"

	// Act
	req, err := http.NewRequest(
		"GET",
		"/emailfinder/search/"+searchTerm,
		nil)
	handleError(err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(r.ServeHTTP)

	handler.ServeHTTP(rr, req)

	// Assert
	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("Something went wrong, status code: %v but want %v", status, http.StatusOK)
	}
}

func BenchmarkIndexer(b *testing.B) {
	r := chi.NewRouter()
	r.Get("/emailfinder/search/{term}", searchBodyByTerm)

	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest("GET", "/emailfinder/search/apte", nil)
		handleError(err)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(r.ServeHTTP)

		handler.ServeHTTP(rr, req)
	}
}
