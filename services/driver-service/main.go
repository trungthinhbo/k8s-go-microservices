package main

import (
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

var GrpcAddr = ":9092"
var HttpAddr = ":8082"

func main() {
	// Start HTTP server
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request from: %s to %s", r.RemoteAddr, r.URL.Path)
			writeJSON(w, http.StatusOK, map[string]string{"message": "Hello, Driver Service!"})
		})
		log.Println("Starting driver service HTTP on port 8082")
		log.Fatal(http.ListenAndServe(HttpAddr, nil))
	}()

	// Start gRPC server
	lis, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	service := NewService()
	NewGrpcHandler(grpcServer, service)

	log.Printf("Starting gRPC server on port %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
