package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type UserService struct {
	Handler http.Handler
}

func StartUserServiceWithChiRouter() *UserService {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/allusers", GetAllUsers)
	r.Get("/allauthors", GetAllAuthors)
	r.Get("/allbooks", GetAllBooks)
	r.Post("/takebook", TakeBook)
	r.Post("/returnbook", ReturnBook)
	r.NotFound(NotFound)
	return &UserService{Handler: r}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func TakeBook(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func ReturnBook(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}
