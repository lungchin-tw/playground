package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"playground/grpc/pb"
	"playground/grpc/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Println("Start up server on port:", *port)

	authService := service.NewAuthService()

	interceptor := service.NewServerAuthInterceptor(
		map[string]bool{
			"/CPUService/DemoPermission":      false,
			"/CPUService/DemoTokenPermission": true,
		},
		authService.Secret(),
	)

	options := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	grpcServer := grpc.NewServer(options...)

	pb.RegisterCPUServiceServer(grpcServer, service.NewCPUService())
	pb.RegisterAuthServiceServer(grpcServer, authService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", *port))
	if err != nil {
		log.Panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Panic(err)
	}
}
