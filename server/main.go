// server/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/maxhorowitz/windows_go_grpc_unix_sox/pb" // Replace with your actual protobuf package

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const socketPath = "/tmp/grpc.sock"

// Define your gRPC server
type server struct {
	pb.UnimplementedReverseServiceServer
}

func (s *server) GetReverse(ctx context.Context, req *pb.GetReverseRequest) (*pb.GetReverseResponse, error) {
	original := req.GetOriginal()
	if original == "" {
		return &pb.GetReverseResponse{Result: ""}, nil
	}
	runes := []rune(original) // Convert string to rune slice to handle Unicode characters
	n := len(runes)
	for i := 0; i < n/2; i++ {
		// Swap runes
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return &pb.GetReverseResponse{Result: string(runes)}, nil // Convert rune slice back to string}
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(socketPath)
		os.Exit(1)
	}()

	grpcServer := grpc.NewServer()
	pb.RegisterReverseServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	fmt.Printf("gRPC server listening on %s\n", socketPath)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
