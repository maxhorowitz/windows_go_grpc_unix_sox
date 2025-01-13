#!/bin/bash

# Ensure the script stops on errors
set -e

# Install Go protobuf compiler plugin
echo "Installing protoc-gen-go..."
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Install Go gRPC compiler plugin
echo "Installing protoc-gen-go-grpc..."
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

echo "Dependencies installed successfully!"
