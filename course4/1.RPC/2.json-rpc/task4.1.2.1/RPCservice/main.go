package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
)

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

type GeoRPC struct{}

func (g *GeoRPC) SearchAnswer(coordinates *RequestAddressSearch, reply *ResponseAddress) error {
	var address ResponseAddress
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", coordinates.Lat, coordinates.Lng)
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &address)
	if err != nil {
		return err
	}

	*reply = address
	return nil
}

func (g *GeoRPC) GeocodeAnswer(address *Address, reply *[]GetCoords) error {
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
	var coords []GetCoords

	resp, err := http.Get(request)
	if err != nil {
		return err
	}

	answer, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(answer, &coords)
	if err != nil {
		return err
	}

	*reply = coords
	return nil
}

func main() {
	geoServicer := new(GeoRPC)
	rpc.Register(geoServicer)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error with server starting", err)
	}

	log.Println("Server is online")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error with connection accepting", err)
		}

		go rpc.ServeConn(conn)
	}
}