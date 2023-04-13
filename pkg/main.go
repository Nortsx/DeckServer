package main

import (
	"../pkg/greeter/api"
	"../pkg/greeter/grpcserver"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	grpcServer := &grpcserver.Server{}
	api.RegisterGreeterServer(server, grpcServer)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
