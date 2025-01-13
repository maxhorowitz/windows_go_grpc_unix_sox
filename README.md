# Go gRPC Unix Socket Example

This project demonstrates a simple gRPC server and client in Go that communicate with each other over a Unix socket. The server exposes a `GetReverse` RPC, which takes a string and returns its reversed version.

## Prerequisites

Before running this project, ensure the following requirements are met:

1. **Golang Installed**: Make sure Go is installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
   - Verify installation by running:
     ```bash
     go version
     ```

2. **Protocol Buffers Compiler (`protoc`)**: Install the Protocol Buffers compiler for generating Go code. Follow the installation guide at [https://grpc.io/docs/protoc-installation/](https://grpc.io/docs/protoc-installation/).

3. **Go Plugins for Protocol Buffers**: Installed via the included `install_deps.sh` script (explained below).

## Installation

1. Clone the repository:
   ```bash
   git clone git@github.com:maxhorowitz/windows_go_grpc_unix_sox.git
   cd windows_go_grpc_unix_sox
   ```

## Testing go-grpc over unix

1. Start the gRPC server
   ```bash
   chmod +x run_server.sh && ./run_server.sh
   ```

2. Make a gRPC client request to the server
   ```bash
   chmod +x run_client.sh && ./run_client.sh arg
   ```
   The `arg` is an arbitrary string which will be reversed and returned to the client.