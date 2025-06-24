package main

import (
	"context"
	"log"
	"time"

	"github.com/TechnoDiktator/grpcTutorial/proto"
)

func CallSayHello(client proto.GreetServiceClient) {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the SayHello method on the client
	response, err := client.SayHello(ctx, &proto.NoParam{})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	// Print the response message
	log.Printf("Response from SayHello: %s", response.Message)
}
