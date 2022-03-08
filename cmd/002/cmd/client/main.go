package main

import (
	"context"
	"log"
	"time"

	"github.com/suzuito/sandbox-go/cmd/002/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		"127.0.0.1:8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewViewStorageServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, err := c.ReadFile(ctx, &pb.RequestReadFile{
		Org: "Hoge campany",
		Permissions: map[string]*pb.Permission{
			"storage.viewer": {
				Buckets: []string{"/foo"},
			},
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", r.String())
}
