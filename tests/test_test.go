package tests

import (
	"context"
	"google.golang.org/grpc"
	"grpc-test-one/test"
	"log"
	"testing"
)

const PORT = "9001"

func TestTest(t *testing.T)  {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := test.NewTestServiceClient(conn)
	resp, err := client.Search(context.Background(), &test.TestRequest{
		Name: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Test err: %v", err)
	}

	log.Printf("resp: %s", resp.GetName())
}