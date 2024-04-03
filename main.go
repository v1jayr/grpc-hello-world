package main

import (
	"fmt"
	"log"
	"net"

	"github.com/v1jayr/grpc-hello-world/lib/grpc/proto"
	"github.com/v1jayr/grpc-hello-world/lib/server"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	helloWorldImpl := server.HelloWorldServer{}
	proto.RegisterHelloWorldServer(grpcServer, helloWorldImpl)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9091))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)
}
