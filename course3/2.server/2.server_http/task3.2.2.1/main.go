package main

import (
	"fmt"
	"net/http"
)

func handleConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	http.HandleFunc("/", handleConnection)
	http.ListenAndServe(":8080", nil)
}
