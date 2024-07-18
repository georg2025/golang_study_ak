package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := makeRouter()
	http.ListenAndServe("localhost:8080", r)
}

func makeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/1", firstAnswer)
	r.Get("/2", secondAnswer)
	r.Get("/3", thirdAnswer)
	r.NotFound(usualAnswer)
	return r
}

func firstAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}

func secondAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world 2"))
}

func thirdAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world 3"))
}

func usualAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}
