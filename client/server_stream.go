package main

import (
	"context"
	"log"
	"time"

	"github.com/TechnoDiktator/grpcTutorial/proto"
)

func CallSayHelloServerStreaming(client proto.GreetServiceClient) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Call the SayHelloServerStreaming method on the client
	in := &proto.NameList{
		Names: []string{"Alice", "Bob", "Charlie"},
	}

	stream, err := client.SayHelloServerStream(ctx, in)
	if err != nil {
		log.Fatalf("Error calling SayHelloServerStreaming: %v", err)
	}

	// Read messages from the stream
	for {
		response, err := stream.Recv()
		if err != nil {
			log.Printf("Stream closed: %v", err)
			break
		}
		log.Printf("Response from SayHelloServerStreaming: %s", response.Message)
	}
}
