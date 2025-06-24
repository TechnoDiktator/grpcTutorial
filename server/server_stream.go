package main

import (
	"time"

	"github.com/TechnoDiktator/grpcTutorial/proto"
)

func (s *helloServer) SayHelloServerStream(req *proto.NameList, stream proto.GreetService_SayHelloServerStreamServer) error {
	// Implement your server streaming logic here
	for _, name := range req.Names {
		response := &proto.HelloResponse{
			Message: "Hello, " + name + "!",
		}

		time.Sleep(1 * time.Second) // Simulate some processing delay
		if err := stream.Send(response); err != nil {
			return err
		}
	}
	return nil
}
