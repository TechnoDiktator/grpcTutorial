package main

import (
	"context"

	"github.com/TechnoDiktator/grpcTutorial/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *proto.NoParam) (*proto.HelloResponse, error) {
	// Implement your logic here
	// For example, return a simple greeting message
	return &proto.HelloResponse{
		Message: "Hello, World!",
	}, nil
}
