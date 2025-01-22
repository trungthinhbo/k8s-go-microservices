package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pb "github.com/sikozonpc/ride-sharing/shared/proto/trip"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TripGrpcHandler struct {
	pb.UnimplementedTripServiceServer
}

func NewGrpcHandler(s *grpc.Server) {
	handler := &TripGrpcHandler{}

	pb.RegisterTripServiceServer(s, handler)
}

func (h *TripGrpcHandler) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {
	pickup := req.StartLocation
	destination := req.EndLocation

	url := fmt.Sprintf(
		"https://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=full&geometries=geojson",
		pickup.Longitude, pickup.Latitude,
		destination.Longitude, destination.Latitude,
	)

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch route: %v", err)
	}
	defer resp.Body.Close()

	// Read and parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read response: %v", err)
	}

	type OSRMResponse struct {
		Routes []struct {
			Distance float64 `json:"distance"`
			Duration float64 `json:"duration"`
			Geometry struct {
				Coordinates [][]float64 `json:"coordinates"`
			} `json:"geometry"`
		} `json:"routes"`
	}
	var routeResp OSRMResponse

	if err := json.Unmarshal(body, &routeResp); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse response: %v", err)
	}

	// parse OSRMResponse to pb.Route
	route := routeResp.Routes[0]
	geometry := route.Geometry.Coordinates
	coordinates := make([]*pb.Coordinate, len(geometry))
	for i, coord := range geometry {
		coordinates[i] = &pb.Coordinate{
			Latitude:  float32(coord[0]),
			Longitude: float32(coord[1]),
		}
	}

	return &pb.CreateTripResponse{
		Route: &pb.Route{
			Geometry: []*pb.Geometry{
				{
					Coordinates: coordinates,
				},
			},
			Distance: float32(route.Distance),
			Duration: float32(route.Duration),
		},
	}, nil
}
