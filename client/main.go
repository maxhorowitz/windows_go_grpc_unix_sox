// client/main.go
package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "pb" // Replace with your actual protobuf package

	"google.golang.org/grpc"
)

const socketPath = "/tmp/grpc_unix_socket"

func main() {
	conn, err := grpc.Dial(
		socketPath,
		grpc.WithInsecure(), // No TLS over Unix sockets
		grpc.WithContextDialer(func(ctx context.Context, addr string) (net.Conn, error) {
			return net.Dial("unix", addr)
		}),
	)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewYourServiceClient(conn) // Replace with your actual service name

	// Example request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.YourMethodName(ctx, &pb.YourRequest{}) // Replace with your actual method and request
	if err != nil {
		log.Fatalf("Error calling YourMethodName: %v", err)
	}
	log.Printf("Response from server: %v", resp)
}
