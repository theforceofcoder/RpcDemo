package main

import (
	"RpcDemo/rpcroute"
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var (
	out     = os.Stdout
	buf     bytes.Buffer
	logger  = log.New(&buf, "", log.Lmicroseconds)
	logInfo = func(info string) {
		logger.SetOutput(out)
		logger.SetPrefix("INFO:")
		logger.Output(2, info)
	}
)

type routeServer struct {
}

func (s *routeServer) SayHello(ctx context.Context, in *rpcroute.HelloRequest) (*rpcroute.HelloResponse, error) {
	logInfo("Request name: " + in.Name)
	return &rpcroute.HelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
	fmt.Println("Startup RPC Server")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	rpcroute.RegisterRouteServer(s, &routeServer{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve:  %v", err)
	}
}
