// client/main.go
package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/maxhorowitz/windows_go_grpc_unix_sox/pb" // Replace with your actual protobuf package

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

	client := pb.NewReverseServiceClient(conn)

	// Example request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if len(os.Args) != 2 {
		log.Fatalf("Must provide string to be reversed!")
		return
	}

	resp, err := client.GetReverse(ctx, &pb.GetReverseRequest{Original: os.Args[1]})
	if err != nil {
		log.Fatalf("Error calling GetReverse: %v", err)
	}
	log.Printf("Response from server: %v", resp)
}
