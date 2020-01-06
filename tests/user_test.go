package tests

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-one/user"
	"log"
	"testing"
)

const PORT = "9002"

func TestUser(t *testing.T)  {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := user.NewUserServiceClient(conn)
	resp, err := client.Search(context.Background(), &user.UserRequest{
		Name: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Test err: %v", err)
	}

	log.Printf("resp: %s", resp.GetName())
}