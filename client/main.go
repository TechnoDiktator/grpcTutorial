package main

import (
	"log"

	"github.com/TechnoDiktator/grpcTutorial/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080" // Port is the port on which the gRPC server will listen.

)

func main() { // Create a connection to the gRPC server
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	
	log.Println("gRPC server is running on port , connection made", port)
	defer conn.Close()


	// Create a client for the GreetService
	client := proto.NewGreetServiceClient(conn)

	// names := proto.NameList{
	// 	Names: []string{"Alice", "Bob", "Charlie"},
	// }
	//CallSayHello(client)
	
	CallSayHelloServerStreaming(client)
}
