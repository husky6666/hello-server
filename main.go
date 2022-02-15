package main

import (
	"google.golang.org/grpc"
	"hello-server/pb_service"
	"hello-server/provider"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb_service.RegisterHelloServiceServer(server, provider.NewHelloProvider())
	_ = server.Serve(lis)
}
