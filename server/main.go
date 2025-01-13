// server/main.go
package main

import (
	"context"
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
	pb.UnimplementedReverseServiceServer
}

func (s *server) GetReverse(ctx context.Context, req *pb.GetReverseRequest) (*pb.GetReverseResponse, error) {
	return nil, nil
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
	pb.RegisterReverseServiceServer(grpcServer, &server{})

	fmt.Printf("gRPC server listening on %s\n", socketPath)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
