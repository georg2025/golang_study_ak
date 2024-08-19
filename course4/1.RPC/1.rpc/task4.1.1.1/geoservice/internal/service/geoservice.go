package service

import (
	"net/http"
	"net/rpc"

	models "geoservice/models"
)

type GeoServiceOption func(*GeoService)

type GeoService struct {
	Client *rpc.Client
}

func WithRPCClient(client *rpc.Client) GeoServiceOption {
	return func(g *GeoService) {
		g.Client = client
	}
}

func NewGeoService(options ...GeoServiceOption) (*GeoService, error) {
	service := &GeoService{}

	for _, option := range options {
		option(service)
	}

	return service, nil
}

type GeoProvider interface {
	SearchAnswer(coordinates models.RequestAddressSearch) (int, models.ResponseAddress, error)
	GeocodeAnswer(address models.Address) (int, []models.GetCoords, error)
}

func (g *GeoService) SearchAnswer(coordinates models.RequestAddressSearch) (int, models.ResponseAddress, error) {
	args := &coordinates
	var address models.ResponseAddress

	err := g.Client.Call("GeoPRC.SearchAnswer", args, address)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	return http.StatusOK, address, nil
}

func (g *GeoService) GeocodeAnswer(address models.Address) (int, []models.GetCoords, error) {
	args := &address
	var coords []models.GetCoords

	err := g.Client.Call("GeoRPC.GeocodeAnswer", args, &coords)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, coords, nil
}
