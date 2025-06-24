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


func (s *helloServer) SayHelloServerStream(req *proto.NameList, stream proto.GreetService_SayHelloServerStreamServer) error {
	// Implement your server streaming logic here
	for _, name := range req.Names {
		response := &proto.HelloResponse{
			Message: "Hello, " + name + "!",
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
	return nil
}
