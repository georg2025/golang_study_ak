package controller

import (
	service "geoservice/internal/service"

	"github.com/go-chi/jwtauth"
)

type ControllerOption func(*Controller)

type Controller struct {
	Servicer service.GeoServicer
}

func NewController(token *jwtauth.JWTAuth, options ...ControllerOption) *Controller {
	service := service.NewGeoService(service.WithToken(token))
	controller := &Controller{Servicer: service}
	return controller
}
