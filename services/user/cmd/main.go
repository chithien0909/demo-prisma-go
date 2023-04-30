package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"prisma-go/prisma/generated/prisma-client"
	gen "prisma-go/proto/proto/user"
	"prisma-go/services/user/config"
	"prisma-go/services/user/delivery/grpc/handler"
	"prisma-go/services/user/domain/user/usecase"
)

func main() {

	configPath := flag.String("config", "", "application config path")
	env := flag.String("env", "", "application env")
	flag.Parse()

	config.Load(*configPath, *env)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", config.GetAppConfig().GrpcPort))

	s := grpc.NewServer()

	client := prisma.New(nil)

	userUseCase := usecase.NewUserUseCase(client)
	userHandler := handler.NewUserHandler(userUseCase)

	// register the GreeterServerImpl on the gRPC server
	gen.RegisterUserServiceServer(s, userHandler)

	fmt.Printf("grpc server is running on port %v... \n", config.GetAppConfig().GrpcPort)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err while serve %v", err)
	}
}
