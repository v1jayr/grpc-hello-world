package server

import (
	"context"
	"fmt"

	"github.com/v1jayr/grpc-hello-world/lib/grpc/proto"
)

type HelloWorldServer struct {
	proto.HelloWorldServer
}

func (s HelloWorldServer) Greet(ctx context.Context, req *proto.GreetRequest) (*proto.GreetResponse, error) {
	resp := &proto.GreetResponse{
		Message: fmt.Sprintf("Hello %s!", req.Name),
	}
	return resp, nil
}
