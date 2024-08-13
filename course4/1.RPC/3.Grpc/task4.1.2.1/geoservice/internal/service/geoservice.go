package service

import (
	"context"
	"encoding/json"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	pb "geoservice/internal/proto"
	models "geoservice/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	Client *grpc.ClientConn
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
	client, err := grpc.NewClient("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

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

	args, err := json.Marshal(coordinates)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	client := pb.NewGeoServiceClient(g.Client)
	req := &pb.Request{Info: string(args)}

	res, err := client.SearchAnswer(context.Background(), req)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	err = json.Unmarshal([]byte(res.Info), &address)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	return http.StatusOK, address, nil
}

func (g *GeoServiceGRPC) GeocodeAnswer(address models.Address) (int, []models.GetCoords, error) {
	args, err := json.Marshal(address)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	client := pb.NewGeoServiceClient(g.Client)
	req := &pb.Request{Info: string(args)}

	res, err := client.GeocodeAnswer(context.Background(), req)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var coords []models.GetCoords

	err = json.Unmarshal([]byte(res.Info), &coords)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, coords, nil
}
