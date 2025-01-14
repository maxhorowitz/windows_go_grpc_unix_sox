#!/bin/bash

# Ensure the script stops on errors
set -e

# Run the Go server
echo "Starting the gRPC server..."
go run server/main.go

read -p "Press Enter to exit..."
