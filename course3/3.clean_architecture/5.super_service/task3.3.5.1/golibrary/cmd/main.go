package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	controller "golibrary/internal/controller"
)

const (
	address = "localhost:8080"
)

func main() {
	service, err := controller.StartUserService()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Problems with starting service")
	}

	r := StartChiRouter(service)
	server := &http.Server{
		Addr:         address,
		Handler:      r,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	sig := <-sigChan
	fmt.Printf("Recieved signal: %v. Starting shutting down\n", sig)

	shuttingDownTime := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shuttingDownTime)
	defer cancel()

	err = shutDown(ctx, server)
	time.Sleep(shuttingDownTime)

	if err == nil {
		log.Println("Server stopped gracefully")
	}
}

func shutDown(ctx context.Context, server *http.Server) error {
	return server.Shutdown(ctx)
}

func StartChiRouter(service controller.UserServicer) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/allusers", service.GetAllUsers)
	r.Get("/allauthors", service.GetAllAuthors)
	r.Get("/allbooks", service.GetAllBooks)
	r.Post("/takebook/{id}", service.TakeBook)
	r.Post("/returnbook/{id}", service.ReturnBook)
	r.Post("/addbook", service.AddBook)
	r.NotFound(service.NotFound)
	return r
}
