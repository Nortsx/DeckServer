package main

import (
	generated "github.com/Nortsx/deckserver/protos/generated/greeter"
	"github.com/Nortsx/deckserver/src/greeter"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	port := setConfig()

	server := grpc.NewServer()
	grpcServer := &greeter.Server{}
	generated.RegisterGreeterServer(server, grpcServer)

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
