package main

import (
	"RpcDemo/rpcroute"
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

const (
	server      = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := rpcroute.NewRouteClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &rpcroute.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	log.Printf("Get: %s", r.Message)

}
