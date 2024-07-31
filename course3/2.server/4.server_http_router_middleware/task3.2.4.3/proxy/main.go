package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target := "http://hugo:1313"
	url, err := url.Parse(target)

	if err != nil {
		log.Fatalf("error with url parsing: %v", err)
	}

	router := httputil.NewSingleHostReverseProxy(url)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if shouldProxy(r) {
			router.ServeHTTP(w, r)
		} else {
			answer(w, r)
		}
	})

	fmt.Println("starting server")
	http.ListenAndServe(":8080", nil)
}

func answer(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

func shouldProxy(r *http.Request) bool {
	return r.URL.Path != "/api/"
}
