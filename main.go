package main

import (
	"fmt"
	"grpc_loader/cmd/client"
	"grpc_loader/cmd/server"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Specify service")
		return
	}
	switch os.Args[1] {
	case "server":
		server.RunServer()
	case "client":
		client.RunClient()
	default:
		fmt.Printf("Unknown service name %s", os.Args[1])
	}
}
