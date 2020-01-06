package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-one/user"
	"log"
	"net"
)

type UserService struct {}

func (s *UserService) Search(ctx context.Context, r *user.UserRequest) (*user.UserResponse, error) {
	return &user.UserResponse{Name: r.Name}, nil
}


const PORT = "9002"

func main() {
	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, &UserService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	_ = server.Serve(lis)
}