package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	gw "prisma-go/proto/proto/user"
	"prisma-go/services/user/config"
)

func run() error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcServerEndpoint := fmt.Sprintf("localhost:%d", config.GetAppConfig().GrpcPort)
	err := gw.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.GetAppConfig().Port), mux)
}

func main() {
	configPath := flag.String("config", "", "application config path")
	env := flag.String("env", "", "application env")
	flag.Parse()

	config.Load(*configPath, *env)
	defer glog.Flush()
	log.Printf("proxy is running... on port %v", config.GetAppConfig().Port)
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
