package main

import (
	"fmt"
	"log"
	"time"

	pb "github.com/sikozonpc/ride-sharing/shared/proto/driver"
	"google.golang.org/grpc"
)

type DriverGrpcHandler struct {
	pb.UnimplementedDriverServiceServer

	service *Service
}

func NewGrpcHandler(s *grpc.Server, service *Service) {
	var drivers []*Driver
	for i, route := range predefinedRoutes {
		drivers = append(drivers, &Driver{
			DriverId: fmt.Sprintf("driver-%d", i),
			Route:    route,
			Index:    0,
		})
	}
	handler := &DriverGrpcHandler{
		service: service,
	}

	pb.RegisterDriverServiceServer(s, handler)
}

func (h *DriverGrpcHandler) FindNearbyDrivers(reqStream pb.DriverService_FindNearbyDriversServer) error {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-reqStream.Context().Done():
			log.Println("Client disconnected from gRPC stream")
			return reqStream.Context().Err()
		case <-ticker.C:
			drivers := h.service.FindNearbyDrivers()

			// Stream updated drivers
			err := reqStream.Send(&pb.StreamDriversResponse{NearbyDrivers: drivers})
			if err != nil {
				log.Println("Error streaming drivers:", err)
				return err
			}
		}
	}
}
