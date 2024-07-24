package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	adress = "localhost:8080"
)

func main() {
	r := makeRouter()
	http.ListenAndServe(adress, r)
}

func makeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/api/address/search", searchAnswer)
	r.Post("/api/address/geocode", geocodeAnswer)
	r.NotFound(usualAnswer)
	return r
}

func searchAnswer(w http.ResponseWriter, r *http.Request) {
	var coordinates RequestAddressSearch
	json.NewDecoder(r.Body).Decode(&coordinates)
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", coordinates.Lat, coordinates.Lng)

	resp, err := http.Get(url)

	if err != nil {
		http.Error(w, "url error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	var address ResponseAddress

	err = json.Unmarshal(body, &address)
	if err != nil {
		http.Error(w, "unmarshal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("you are in " + address.Address.City))
}

func geocodeAnswer(w http.ResponseWriter, r *http.Request) {
	var address Address
	json.NewDecoder(r.Body).Decode(&address)

	parts := []string{}
	parts = append(parts, strings.Split(address.House_number, " ")...)
	parts = append(parts, strings.Split(address.Road, " ")...)
	parts = append(parts, strings.Split(address.Suburb, " ")...)
	parts = append(parts, strings.Split(address.City, " ")...)
	parts = append(parts, strings.Split(address.State, " ")...)
	parts = append(parts, strings.Split(address.Country, " ")...)

	var sb strings.Builder
	for _, i := range parts {
		if i != "" {
			sb.WriteString("+")
			sb.WriteString(i)
		}
	}

	request := "https://nominatim.openstreetmap.org/search?q=" + strings.Trim(sb.String(), "+") + "&format=json"

	resp, err := http.Get(request)
	if err != nil {
		http.Error(w, "url error", http.StatusInternalServerError)
		return
	}

	answer, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	var coords []GetCoords

	err = json.Unmarshal(answer, &coords)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Your lattitude = " + coords[0].Lat + "; Your longitude = " + coords[0].Lon))
}

func usualAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

type ResponseAddress struct {
	Address Address `json:"address"`
}

type Address struct {
	House_number string `json:"house_number"`
	Road         string `json:"road"`
	Suburb       string `json:"suburb"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

type RequestAddressSearch struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type GetCoords struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}
