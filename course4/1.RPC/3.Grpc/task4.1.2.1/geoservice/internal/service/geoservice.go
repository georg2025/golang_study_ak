package service

import (
	"context"
	"encoding/json"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"google.golang.org/grpc"

	pb "geoservice/internal/proto"
	models "geoservice/models"
)

type RPCFactory struct{}

type JRPCFactory struct{}

type GRPCFactory struct{}

type GeoServiceRPC struct {
	Client *rpc.Client
}

type GeoServiceJSONRPC struct {
	Client *rpc.Client
}

type GeoServiceGRPC struct {
	Client pb.GeoServiceClient
}

type GeoServiceFactory interface {
	MakeGeoProvider() (GeoProvider, error)
}

func (f *RPCFactory) MakeGeoProvider() (GeoProvider, error) {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		return nil, err
	}

	service := &GeoServiceRPC{Client: client}
	return service, nil
}

func (f *JRPCFactory) MakeGeoProvider() (GeoProvider, error) {
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		return nil, err
	}

	service := &GeoServiceJSONRPC{Client: client}
	return service, nil
}

func (f *GRPCFactory) MakeGeoProvider() (GeoProvider, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewGeoServiceClient(conn)

	service := &GeoServiceGRPC{Client: client}
	return service, nil
}

type GeoProvider interface {
	SearchAnswer(coordinates models.RequestAddressSearch) (int, models.ResponseAddress, error)
	GeocodeAnswer(address models.Address) (int, []models.GetCoords, error)
}

func (g *GeoServiceRPC) SearchAnswer(coordinates models.RequestAddressSearch) (int, models.ResponseAddress, error) {
	args := &coordinates
	var address models.ResponseAddress

	err := g.Client.Call("GeoPRC.SearchAnswer", args, address)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	return http.StatusOK, address, nil
}

func (g *GeoServiceRPC) GeocodeAnswer(address models.Address) (int, []models.GetCoords, error) {
	args := &address
	var coords []models.GetCoords

	err := g.Client.Call("GeoRPC.GeocodeAnswer", args, &coords)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, coords, nil
}

func (g *GeoServiceJSONRPC) SearchAnswer(coordinates models.RequestAddressSearch) (int, models.ResponseAddress, error) {
	var address models.ResponseAddress

	args, err := json.Marshal(coordinates)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	var reply string

	err = g.Client.Call("GeoPRC.SearchAnswer", string(args), &reply)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	err = json.Unmarshal([]byte(reply), &address)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	return http.StatusOK, address, nil
}

func (g *GeoServiceJSONRPC) GeocodeAnswer(address models.Address) (int, []models.GetCoords, error) {
	args, err := json.Marshal(address)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var reply string

	err = g.Client.Call("GeoRPC.GeocodeAnswer", string(args), &reply)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var coords []models.GetCoords

	err = json.Unmarshal([]byte(reply), &coords)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, coords, nil
}

func (g *GeoServiceGRPC) SearchAnswer(coordinates models.RequestAddressSearch) (int, models.ResponseAddress, error) {
	var address models.ResponseAddress

	req := pb.RequestAddress{
		Lat: coordinates.Lat,
		Lng: coordinates.Lng,
	}

	res, err := g.Client.SearchAnswer(context.Background(), &req)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	address.Address = models.Address{
		House_number: res.Address.HouseNumber,
		Road:         res.Address.Road,
		Suburb:       res.Address.Suburb,
		City:         res.Address.City,
		State:        res.Address.State,
		Country:      res.Address.Country,
	}

	return http.StatusOK, address, nil
}

func (g *GeoServiceGRPC) GeocodeAnswer(address models.Address) (int, []models.GetCoords, error) {
	req := pb.Address{
		HouseNumber: address.House_number,
		Road:        address.Road,
		Suburb:      address.Suburb,
		City:        address.City,
		State:       address.State,
		Country:     address.Country,
	}

	res, err := g.Client.GeocodeAnswer(context.Background(), &req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var coords []models.GetCoords

	for _, i := range res.Coords {
		coords = append(coords, models.GetCoords{
			Lat: i.Lat,
			Lon: i.Lon,
		})
	}

	return http.StatusOK, coords, nil
}
