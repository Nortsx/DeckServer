package grpcserver

import (
	"../api"
	"context"
)

type Server struct{}

func (server *Server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloResponse, error) {
	if err := in.ValidateAll(); err != nil {
		return &api.HelloResponse{Message: err.Error()}, nil
	}
	return &api.HelloResponse{Message: in.Message}, nil
}
