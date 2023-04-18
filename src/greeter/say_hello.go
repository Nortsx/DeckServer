package greeter

import (
	"context"
	"github.com/Nortsx/deckserver/protos/generated/greeter"
)

type Server struct{}

func (server *Server) SayHello(ctx context.Context, in *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	response := in.Message
	if err := in.ValidateAll(); err != nil {
		response = err.Error()
	}
	return &greeter.HelloResponse{Message: response}, nil
}
