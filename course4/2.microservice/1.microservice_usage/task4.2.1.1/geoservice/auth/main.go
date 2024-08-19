package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "auth/go"
	repository "auth/internal/repository"
	models "auth/models"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
)

const serveAddress = ":50052"

const secret = "my secret"

type AuthRPC struct {
	pb.UnimplementedAuthServiceServer
	Database *repository.PostgressDataBase
	Secret   []byte
}

type MyClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

func (a *AuthRPC) RegisterUser(ctx context.Context, req *pb.User) (*pb.IsRegistered, error) {
	user := models.User{
		Username: req.Login,
		Password: req.Password,
	}

	_, ok, _ := a.Database.GetByName(context.Background(), user.Username)

	if ok {
		return nil, fmt.Errorf("username already exist")
	}

	err := a.Database.Create(context.Background(), user)

	if err != nil {
		return nil, err
	}

	return &pb.IsRegistered{IsRegistered: true}, nil
}

func (a *AuthRPC) LoginUser(ctx context.Context, req *pb.User) (*pb.IsLogined, error) {
	user := models.User{
		Username: req.Login,
		Password: req.Password,
	}

	databaseUser, ok, _ := a.Database.GetByName(context.Background(), user.Username)

	if !ok {
		return nil, fmt.Errorf("user dont exist")
	}

	if user.Password != databaseUser.Password {
		return nil, fmt.Errorf("invalid password")
	}

	claims := MyClaims{
		"bar",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring, err := token.SignedString(a.Secret)
	if err != nil {
		return nil, err
	}

	answer := &pb.IsLogined{
		IsLogined: true,
		Token: &pb.Token{
			Token: tokenstring,
		},
	}

	return answer, nil
}

func (a *AuthRPC) CheckToken(ctx context.Context, req *pb.Token) (*pb.IsValid, error) {
	token, err := jwt.ParseWithClaims(req.Token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.Secret, nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil {
		return nil, err
	} else if _, ok := token.Claims.(*MyClaims); ok {
		return &pb.IsValid{IsValid: true}, nil
	} else {
		return nil, fmt.Errorf("unknown claims type")
	}
}

func main() {
	listener, err := net.Listen("tcp", serveAddress)
	if err != nil {
		log.Fatal("Error with server starting", err)
	}

	dataBase, err := repository.StartPostgressDataBase(context.Background())
	if err != nil {
		log.Fatal("Cant start database", err)
	}

	RPCserver := &AuthRPC{
		Database: dataBase,
		Secret:   []byte(secret),
	}
	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, RPCserver)

	log.Println("Server is online")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
