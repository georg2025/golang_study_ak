package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"

	"google.golang.org/grpc"

	pb "geoservice/go"
	models "geoservice/models"
)

type GeoServiceGRPC struct {
	ClientGeo  pb.GeoServiceClient
	ClientUser pb.UserServiceClient
	ClientAuth pb.AuthServiceClient
}

type GeoServiceFactory interface {
	MakeGeoServicer() (GeoServicer, error)
}

type GeoServiceGRPCOption func(*GeoServiceGRPC)

func MakeGeoServicer(options ...GeoServiceGRPCOption) (GeoServicer, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	Geoclient := pb.NewGeoServiceClient(conn)

	conn, err = grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	Authclient := pb.NewAuthServiceClient(conn)

	conn, err = grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	Userclient := pb.NewUserServiceClient(conn)

	service := &GeoServiceGRPC{
		ClientGeo:  Geoclient,
		ClientAuth: Authclient,
		ClientUser: Userclient,
	}

	for _, option := range options {
		option(service)
	}

	return service, nil
}

type GeoServicer interface {
	RegisterUser(user models.User) (int, error)
	LoginUser(user models.User) (int, string, error)
	SearchAnswer(coordinates models.RequestAddressSearch, token string) (int, models.ResponseAddress, error)
	GeocodeAnswer(address models.Address, token string) (int, []models.GetCoords, error)
	GetByID(id string) (models.User, error)
}

func (g *GeoServiceGRPC) RegisterUser(user models.User) (int, error) {
	passwordHash := hashPassword([]byte(user.Password))

	req := pb.User{
		Login:    user.Username,
		Password: passwordHash,
	}

	res, err := g.ClientAuth.RegisterUser(context.Background(), &req)
	if err != nil {
		return http.StatusNotFound, err
	}

	if res.IsRegistered != true {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (g *GeoServiceGRPC) LoginUser(user models.User) (int, string, error) {
	passwordHash := hashPassword([]byte(user.Password))

	req := pb.User{
		Login:    user.Username,
		Password: passwordHash,
	}

	res, err := g.ClientAuth.LoginUser(context.Background(), &req)
	if err != nil {
		return http.StatusNotFound, "", err
	}

	token := res.Token.Token

	return http.StatusOK, token, nil
}

func (g *GeoServiceGRPC) GetByID(id string) (models.User, error) {
	var user models.User

	userId, err := strconv.Atoi(id)
	if err != nil {
		return user, err
	}

	req := pb.ID{Id: int32(userId)}

	res, err := g.ClientUser.GetByID(context.Background(), &req)
	if err != nil {
		return user, err
	}

	user.Username = res.Username

	return user, nil
}

func (g *GeoServiceGRPC) SearchAnswer(coordinates models.RequestAddressSearch, token string) (int, models.ResponseAddress, error) {
	var address models.ResponseAddress

	permissionReq := pb.Token{
		Token: token,
	}
	permission, err := g.ClientAuth.CheckToken(context.Background(), &permissionReq)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	if !permission.IsValid {
		return http.StatusForbidden, address, err
	}

	req := pb.RequestAddress{
		Lat: coordinates.Lat,
		Lng: coordinates.Lng,
	}

	res, err := g.ClientGeo.SearchAnswer(context.Background(), &req)
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

func (g *GeoServiceGRPC) GeocodeAnswer(address models.Address, token string) (int, []models.GetCoords, error) {
	permissionReq := pb.Token{
		Token: token,
	}
	permission, err := g.ClientAuth.CheckToken(context.Background(), &permissionReq)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if !permission.IsValid {
		return http.StatusForbidden, nil, err
	}

	req := pb.Address{
		HouseNumber: address.House_number,
		Road:        address.Road,
		Suburb:      address.Suburb,
		City:        address.City,
		State:       address.State,
		Country:     address.Country,
	}

	res, err := g.ClientGeo.GeocodeAnswer(context.Background(), &req)
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

func hashPassword(password []byte) string {
	hash := sha256.New()
	hash.Write(password)
	return hex.EncodeToString(hash.Sum(nil))
}
