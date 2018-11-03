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

//定義 listen port
const (
	port = ":50051"
)

//log輸出
var (
	out    = os.Stdout
	buf    bytes.Buffer
	logger = log.New(&buf, "", log.Lmicroseconds)

	logInfo = func(info string) {
		logger.SetOutput(out)
		logger.SetPrefix("INFO:")
		logger.Output(2, info)
	}
	logDebug = func(debug string) {
		logger.SetOutput(out)
		logger.SetPrefix("DEBUG:")
		logger.Output(2, debug)
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
