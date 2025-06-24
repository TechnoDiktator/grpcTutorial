package main

import (
	"context"
	"log"
	"time"

	"github.com/TechnoDiktator/grpcTutorial/proto"
)

func CallSayHelloClientStreaming(client proto.GreetServiceClient) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	// Create a stream for client streaming
	stream, err := client.SayHelloClientStream(ctx)
	if err != nil {
		log.Fatalf("Error creating client stream: %v", err)
	}

	// Send multiple messages to the server
	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		request := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(request); err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
		time.Sleep(1 * time.Second) // Simulate some processing delay
	}

	// 	// Close the stream and receive the final response
	// 	response, err := stream.CloseAndRecv()
	// 	if err != nil {
	// 		log.Fatalf("Error receiving response: %v", err)
	// 	}

	// 	for _, message := range response.Messages {
	// 		log.Printf("Received message: %s", message)
	// 	}
	// }
}
