#!/bin/bash

# Ensure the script stops on errors
set -e

# Check if a string argument is provided
if [ $# -eq 0 ]; then
  echo "Usage: $0 <string_argument>"
  exit 1
fi

# Get the string argument
ARG=$1

# Run the Go client with the argument
echo "Starting the gRPC client with argument: $ARG"
go run client/main.go "$ARG"
