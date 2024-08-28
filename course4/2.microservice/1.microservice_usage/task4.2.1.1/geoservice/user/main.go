package main

import (
	"context"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "user/go"
	repository "user/internal/repository"
)

const serveAddress = ":50053"

type UserRPC struct {
	Database *repository.PostgressDataBase
	pb.UnimplementedUserServiceServer
}

func (u *UserRPC) GetByID(ctx context.Context, req *pb.ID) (*pb.UserName, error) {
	userId := strconv.Itoa(int(req.Id))

	userName, err := u.Database.GetByID(context.Background(), userId)
	if err != nil {
		return nil, err
	}

	return &pb.UserName{Username: userName.Username}, nil
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

	RPCserver := &UserRPC{
		Database: dataBase,
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, RPCserver)

	log.Println("Server is online")

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
