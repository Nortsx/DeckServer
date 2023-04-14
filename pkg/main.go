package main

import (
	"github.com/Nortsx/deckserver/pkg/greeter/api"
	"github.com/Nortsx/deckserver/pkg/greeter/grpcserver"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	port := setConfig()

	server := grpc.NewServer()
	grpcServer := &grpcserver.Server{}
	api.RegisterGreeterServer(server, grpcServer)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func setConfig() (port string) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	return viper.GetString("PORT")
}
