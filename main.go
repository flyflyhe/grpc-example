package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-one/test"
	"log"
	"net"
)

type TestService struct {}

func (s *TestService) Search(ctx context.Context, r *test.TestRequest) (*test.TestResponse, error) {
	return &test.TestResponse{Name:r.Name}, nil
}


const PORT = "9001"

func main() {
	server := grpc.NewServer()
	test.RegisterTestServiceServer(server, &TestService{})

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	_ = server.Serve(lis)
}