package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handleConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		log.Fatalf("failed to load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	log.Printf("Server is running on port %s", port)
	http.HandleFunc("/", handleConnection)
	http.ListenAndServe(port, nil)
}
