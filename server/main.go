// server/main.go
package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/maxhorowitz/windows_go_grpc_unix_sox/pb" // Replace with your actual protobuf package

	"google.golang.org/grpc"
)

const socketPath = "/tmp/grpc_unix_socket"

// Define your gRPC server
type server struct {
}

func (s *server) YourMethodName(req *pb.YourRequest, stream pb.YourService_YourMethodNameServer) error {
	// Implement your logic here
	return nil
}

func main() {
	// Clean up any existing socket file
	if _, err := os.Stat(socketPath); err == nil {
		os.Remove(socketPath)
	}

	// Create a listener for the Unix socket
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to listen on Unix socket: %v", err)
	}
	defer os.Remove(socketPath)

	grpcServer := grpc.NewServer()

	fmt.Printf("gRPC server listening on %s\n", socketPath)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
