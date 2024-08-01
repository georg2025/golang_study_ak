package main

import (
	"context"
	"fmt"
	controller "golibrary/internal/controller"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	Address = "localhost:8080"
)

func main() {
	handler := controller.StartUserServiceWithChiRouter()
	server := &http.Server{
		Addr:         Address,
		Handler:      handler.Handler,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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

	err := shutDown(ctx, server)
	time.Sleep(shuttingDownTime)

	if err == nil {
		log.Println("Server stopped gracefully")
	}
}
