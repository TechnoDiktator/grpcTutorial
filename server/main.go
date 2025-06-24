package main

import (
	"log"
	"net"
	pb "github.com/TechnoDiktator/grpcTutorial/proto" // Import your generated protobuf package
	"google.golang.org/grpc"
)

const ( // Port is the port on which the gRPC server will listen.
	Port = ":8080"
)

type helloServer struct {
	// Implement your gRPC service methods here
	pb.GreetServiceServer // Assuming you have a GreetServiceServer interface in your proto file

}


func main() {

	lis, err := net.Listen("tcp", Port)

	if err != nil {
		panic(err)
		log.Fatal("Failed to listen on port ans start the server", Port, ":", err)
		return
	}

	grpcServer := grpc.NewServer()

	// Register your gRPC service here
	// For example:
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("gRPC server is starting on port %s %s", Port, lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to start gRPC server on port", Port, ":", err)
		return

	}

	log.Println("gRPC server is running on port", Port)
}
