package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/v1jayr/grpc-hello-world/lib/grpc/proto"
	"github.com/v1jayr/grpc-hello-world/lib/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultGrpcPort = 9090
	defaultHttpPort = 9080
)

func main() {
	proxyMode := flag.Bool("proxy", false, "start reverse proxy")
	grpcPort := flag.Uint("grpc-port", defaultGrpcPort, "grpc server port")
	httpPort := flag.Uint("http-port", defaultHttpPort, "grpc server port")
	flag.Parse()
	defer glog.Flush()

	if *proxyMode {
		fmt.Printf("starting HTTP gateway at port=%d\n", *httpPort)
		startHttpReverseProxy(*grpcPort, *httpPort)
	} else {
		fmt.Printf("starting gRPC server at port=%d\n", *grpcPort)
		startGrpcServer(*grpcPort)
	}
}

func startGrpcServer(port uint) {
	grpcServer := grpc.NewServer()
	helloWorldImpl := server.HelloWorldServer{}
	proto.RegisterHelloWorldServer(grpcServer, helloWorldImpl)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)
}

func startHttpReverseProxy(grpcPort, httpPort uint) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	option := runtime.WithErrorHandler(func(ctx context.Context, sm *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {

		b, errx := io.ReadAll(r.Body)
		if errx != nil {
			fmt.Printf("failed to read req, %v\n", errx)
		}
		fmt.Printf("[Method=%s] [URL=%s] Request=%s\n", r.Method, r.RequestURI, string(b))

		fmt.Printf("%v\n", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal error"))
	})

	mux := runtime.NewServeMux(option)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	endpoint := fmt.Sprintf("localhost:%d", grpcPort)
	if err := proto.RegisterHelloWorldHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux)
}
