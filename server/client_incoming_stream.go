package main

import (
	"context"
	"log"
	"time"

	"github.com/TechnoDiktator/grpcTutorial/proto"
)

func (s *helloServer) SayHelloClientStream(stream proto.GreetService_SayHelloClientStreamServer) error {
	// Create a slice to hold the received messages
	messages := make([]string, 0)
	// Receive messages from the stream
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == context.Canceled {
				log.Println("Stream canceled")
				return nil
			}
			log.Printf("Error receiving message: %v", err)
			return err
		}
		log.Printf("Received stream request message: %s", req.Name)
		messages = append(messages, req.Name)
		time.Sleep(1 * time.Second) // Simulate some processing delay
	}


}
