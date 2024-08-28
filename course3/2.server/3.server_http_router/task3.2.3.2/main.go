package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := makeRouter()
	http.ListenAndServe("localhost:8080", r)
}

func makeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/group1/{id}", firstAnswer)
	r.Get("/group2/{id}", secondAnswer)
	r.Get("/group3/{id}", thirdAnswer)
	r.NotFound(usualAnswer)
	return r
}

func firstAnswer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("group1, Hello world " + id))
}

func secondAnswer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("group2, Hello world " + id))
}

func thirdAnswer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("group3, Hello world " + id))
}

func usualAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}
