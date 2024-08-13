package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc/jsonrpc"
	"strings"

	pb "grpcservice/proto"

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

func (g *GeoRPC) SearchAnswer(ctx context.Context, req *pb.Request) (*pb.Responce, error) {
	var coordinates RequestAddressSearch

	err := json.Unmarshal([]byte(req.Info), &coordinates)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", coordinates.Lat, coordinates.Lng)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &pb.Responce{Info: string(body)}, nil
}

func (g *GeoRPC) GeocodeAnswer(ctx context.Context, req *pb.Request) (*pb.Responce, error) {
	var address Address

	err := json.Unmarshal([]byte(req.Info), &address)
	if err != nil {
		return nil, err
	}

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
		return nil, err
	}

	answer, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &pb.Responce{Info: string(answer)}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error with server starting", err)
	}

	server := grpc.NewServer()
	pb.RegisterGeoServiceServer(server, &GeoRPC{})

	log.Println("Server is online")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error with connection accepting", err)
		}

		go jsonrpc.ServeConn(conn)
	}
}
