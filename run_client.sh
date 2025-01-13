#!/bin/bash

# Ensure the script stops on errors
set -e

# Run the Go client
echo "Starting the gRPC client..."
go run client/main.go
