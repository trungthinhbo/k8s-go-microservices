package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

var GrpcAddr = ":9093"

func main() {
	// Start gRPC server
	lis, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	NewGrpcHandler(grpcServer)

	log.Printf("Starting gRPC server Trip service on port %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
