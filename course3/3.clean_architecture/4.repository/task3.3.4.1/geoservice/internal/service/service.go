package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	repository "geoservice/internal/repository"
	models "geoservice/models"

	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt"
)

type ServiceOption func(*GeoService)

type GeoService struct {
	Database repository.UserRepository
	Token    *jwtauth.JWTAuth
}

func WithToken(token *jwtauth.JWTAuth) ServiceOption {
	return func(c *GeoService) {
		c.Token = token
	}
}

func NewGeoService(options ...ServiceOption) (*GeoService, error) {
	dataBase, err := repository.StartPostgressDataBase(context.Background())

	if err != nil {
		return nil, err
	}

	service := &GeoService{Database: dataBase}

	for _, option := range options {
		option(service)
	}

	return service, nil
}

type GeoServicer interface {
	RegisterUser(user models.User) (int, error)
	LoginUser(user models.User) (int, string, error)
	SearchAnswer(coordinates models.RequestAddressSearch) (int, models.ResponseAddress, error)
	GeocodeAnswer(address models.Address) (int, []models.GetCoords, error)
}

// registerUser handle POST-requests on api/register
// @Summary Register
// @Tags Login
// @Description Register user
// @Accept  json
// @Produce  json
// @Param  input  body  User true  "username and password"
// @Success 200 {object} string
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/register [post]
func (c *GeoService) RegisterUser(user models.User) (int, error) {
	_, ok, _ := c.Database.GetByName(context.Background(), user.Username)

	if ok {
		return http.StatusInternalServerError, fmt.Errorf("username already exist")
	}

	passwordHash := hashPassword([]byte(user.Password))
	user.Password = passwordHash
	c.Database.Create(context.Background(), user)
	return http.StatusCreated, nil
}

// loginUser handle POST-requests on api/login
// @Summary Login
// @Tags Login
// @Description Login user
// @Accept  json
// @Produce  json
// @Param  input  body  User true  "username and password"
// @Success 200 {object} string
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/login [post]
func (c *GeoService) LoginUser(user models.User) (int, string, error) {
	databaseUser, ok, _ := c.Database.GetByName(context.Background(), user.Username)

	if !ok {
		return http.StatusForbidden, "", fmt.Errorf("user dont exist")
	}

	passwordHash := hashPassword([]byte(user.Password))
	if passwordHash != databaseUser.Password {
		return http.StatusForbidden, "", fmt.Errorf("invalid password")
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      jwtauth.ExpireIn(time.Hour),
	}
	_, tokenString, _ := c.Token.Encode(claims)

	return http.StatusOK, tokenString, nil
}

// searchAnswer handle POST-requests on api/address/search
// @Summary SearchCity
// @Tags Search
// @Description Search city Name by coords
// @Accept  json
// @Produce  json
// @Param  coordinates  body  RequestAddressSearch true  "Lattitude and Longitude"
// @Success 200 {object} string
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/address/search [post]
func (c *GeoService) SearchAnswer(coordinates models.RequestAddressSearch) (int, models.ResponseAddress, error) {
	var address models.ResponseAddress
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", coordinates.Lat, coordinates.Lng)
	resp, err := http.Get(url)

	if err != nil {
		return http.StatusInternalServerError, address, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	err = json.Unmarshal(body, &address)
	if err != nil {
		return http.StatusInternalServerError, address, err
	}

	return http.StatusOK, address, nil
}

// geocodeAnswer handle POST-requests on api/address/geocode
// @Summary SearchCoords
// @Tags Search
// @Description Search coords by address
// @Accept  json
// @Produce  json
// @Param  coordinates  body  Address true  "House number, road, suburb, city, state, country"
// @Success 200 {object} string
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/address/search [post]
func (c *GeoService) GeocodeAnswer(address models.Address) (int, []models.GetCoords, error) {
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
	var coords []models.GetCoords

	resp, err := http.Get(request)
	if err != nil {
		return http.StatusInternalServerError, coords, err
	}

	answer, err := io.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, coords, err
	}

	err = json.Unmarshal(answer, &coords)
	if err != nil {
		return http.StatusInternalServerError, coords, err
	}

	return http.StatusOK, coords, nil
}

func hashPassword(password []byte) string {
	hash := sha256.New()
	hash.Write(password)
	return hex.EncodeToString(hash.Sum(nil))
}
