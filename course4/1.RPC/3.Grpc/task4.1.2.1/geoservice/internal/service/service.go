package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	cache "geoservice/internal/cache"
	repository "geoservice/internal/repository"
	models "geoservice/models"

	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt"
)

type ServiceOption func(*Service)

type Service struct {
	Database repository.UserRepository
	Token    *jwtauth.JWTAuth
}

func WithToken(token *jwtauth.JWTAuth) ServiceOption {
	return func(c *Service) {
		c.Token = token
	}
}

func NewService(options ...ServiceOption) (*Service, error) {
	dataBase, err := cache.NewCache()

	if err != nil {
		return nil, err
	}

	service := &Service{Database: dataBase}

	for _, option := range options {
		option(service)
	}

	return service, nil
}

type Servicer interface {
	RegisterUser(user models.User) (int, error)
	LoginUser(user models.User) (int, string, error)
	GetByID(id string) (models.User, error)
}

func (s *Service) RegisterUser(user models.User) (int, error) {
	_, ok, _ := s.Database.GetByName(context.Background(), user.Username)

	if ok {
		return http.StatusInternalServerError, fmt.Errorf("username already exist")
	}

	passwordHash := hashPassword([]byte(user.Password))
	user.Password = passwordHash
	err := s.Database.Create(context.Background(), user)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (s *Service) LoginUser(user models.User) (int, string, error) {
	databaseUser, ok, _ := s.Database.GetByName(context.Background(), user.Username)

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
	_, tokenString, _ := s.Token.Encode(claims)

	return http.StatusOK, tokenString, nil
}

func (s *Service) GetByID(id string) (models.User, error) {
	return s.Database.GetByID(context.Background(), id)
}

func hashPassword(password []byte) string {
	hash := sha256.New()
	hash.Write(password)
	return hex.EncodeToString(hash.Sum(nil))
}
