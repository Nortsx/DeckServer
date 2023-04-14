package grpcserver

import (
	"context"
	"github.com/Nortsx/deckserver/pkg/greeter/api"
)

type Server struct{}

func (server *Server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	response := in.Message
	if err := in.ValidateAll(); err != nil {
		response = err.Error()
	}
	return &api.HelloResponse{Message: response}, nil
}
