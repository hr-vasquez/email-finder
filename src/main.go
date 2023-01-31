package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

func main() {
	r := configuredChiRouter()

	r.Get("/emailfinder/search/{term}", searchBodyByTerm)

	err := http.ListenAndServe(":3000", r)
	handleError(err)
}

func configuredChiRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	return r
}

func searchBodyByTerm(w http.ResponseWriter, r *http.Request) {
	term := chi.URLParam(r, "term")

	if term == "" {
		w.WriteHeader(http.StatusBadRequest)
		logger.Println("You cannot pass an empty value to search")
		return
	}

	data := searchBody(term)

	w.Header().Set("Content-Type", "application/json")
	jData, errParsing := json.Marshal(data)
	handleError(errParsing)

	_, err := w.Write(jData)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
