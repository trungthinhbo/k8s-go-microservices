package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sikozonpc/ride-sharing/services/api-gateway/grpc_clients/driver_client"
	pb "github.com/sikozonpc/ride-sharing/shared/proto/driver"
	"log"
	"net/http"
)

func handleDriversWebSocket(w http.ResponseWriter, r *http.Request, client *driver_client.DriverServiceClient) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer func(conn *websocket.Conn) {
		_ = conn.Close()
	}(conn)

	log.Println("Client connected via WebSocket")

	// Channel to listen for client messages (e.g., user location updates)
	clientLocations := make(chan *pb.Location)
	go func() {
		for {
			// Read message from WebSocket
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading WebSocket message:", err)
				close(clientLocations)
				return
			}

			var location pb.Location
			if err := json.Unmarshal(msg, &location); err != nil {
				log.Println("Invalid location data:", err)
				continue
			}
			clientLocations <- &location
		}
	}()

	// Start streaming driver updates
	stream, err := client.StreamNearbyDrivers(r.Context())
	if err != nil {
		log.Println("Failed to start gRPC stream:", err)
		return
	}
	defer stream.CloseSend()

	// Handle sending client location updates to the driver service
	go func() {
		for loc := range clientLocations {
			err := stream.Send(&pb.FindNearbyDriversRequest{Location: loc})
			if err != nil {
				log.Println("Failed to send location to driver service:", err)
				return
			}
		}
	}()

	// Receive driver updates from driver service and forward them to the client
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Println("Error receiving driver updates:", err)
			return
		}

		driversJSON, err := json.Marshal(res.NearbyDrivers)
		if err != nil {
			log.Println("Error marshaling driver updates:", err)
			continue
		}

		// Send driver updates to the WebSocket client
		if err := conn.WriteMessage(websocket.TextMessage, driversJSON); err != nil {
			log.Println("Error sending driver updates:", err)
			return
		}
	}
}
