#!/bin/bash

# Ensure the script stops on errors
set -e

# Clear the pb directory and create an empty one
PB_DIR="./pb"
if [ -d "$PB_DIR" ]; then
    rm -rf "$PB_DIR" # Remove the existing pb directory
fi
mkdir -p "$PB_DIR" # Create a new, empty pb directory

# Generate Go code for the reverse.proto file
protoc \
    --go_out=pb \
    --go-grpc_out=pb \
    protos/reverse.proto

echo "Code generation completed successfully!"
