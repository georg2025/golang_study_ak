package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	pb "grpcservice/go"

	"google.golang.org/grpc"
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

type GeoRPC struct {
	pb.UnimplementedGeoServiceServer
}

func (g *GeoRPC) SearchAnswer(ctx context.Context, req *pb.RequestAddress) (*pb.ResponceAddress, error) {
	log.Println("we got search request")

	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", req.Lat, req.Lng)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var address ResponseAddress

	err = json.Unmarshal(body, &address)
	if err != nil {
		return nil, err
	}

	return &pb.ResponceAddress{
		Address: &pb.Address{
			HouseNumber: address.Address.House_number,
			Road:        address.Address.Road,
			Suburb:      address.Address.Suburb,
			City:        address.Address.City,
			State:       address.Address.State,
			Country:     address.Address.Country,
		}}, nil
}

func (g *GeoRPC) GeocodeAnswer(ctx context.Context, req *pb.Address) (*pb.GetCoords, error) {
	log.Println("we got geocode request")

	parts := []string{}
	parts = append(parts, strings.Split(req.HouseNumber, " ")...)
	parts = append(parts, strings.Split(req.Road, " ")...)
	parts = append(parts, strings.Split(req.Suburb, " ")...)
	parts = append(parts, strings.Split(req.City, " ")...)
	parts = append(parts, strings.Split(req.State, " ")...)
	parts = append(parts, strings.Split(req.Country, " ")...)

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
		return nil, err
	}

	answer, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var coords []GetCoords
	err = json.Unmarshal(answer, &coords)
	if err != nil {
		return nil, err
	}

	var grpcRequest pb.GetCoords

	for _, i := range coords {
		coord := &pb.Coords{
			Lat: i.Lat,
			Lon: i.Lon,
		}

		grpcRequest.Coords = append(grpcRequest.Coords, coord)
	}

	return &grpcRequest, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Error with server starting", err)
	}

	server := grpc.NewServer()
	pb.RegisterGeoServiceServer(server, &GeoRPC{})

	log.Println("Server is online")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
