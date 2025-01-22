package main

import (
	"fmt"

	"github.com/mmcloughlin/geohash"
	pb "github.com/sikozonpc/ride-sharing/shared/proto/driver"
)

// Predefined routes for drivers
// Get this coordinates from Google Maps for example and build a route
var predefinedRoutes = [][][]float64{
	{
		{37.768727753110106, -122.41345597077878},
		{37.77019784198334, -122.41298146942745},
		{37.77163599059948, -122.41125013468515},
		{37.773790702602305, -122.41168660785345},
	},
	{
		{37.78938865879484, -122.42206098118852},
		{37.79112418447625, -122.42238063604239},
		{37.79160600785717, -122.42249902672565},
		{37.79111015076977, -122.42238063603901},
		{37.791797800743794, -122.42253454393163},
		{37.79353326986209, -122.4228837964593},
		{37.7939729779581, -122.42294891128813},
		{37.793739091013016, -122.42474844980799},
		{37.79289241408059, -122.42455310517508},
		{37.792172029382556, -122.4244643121601},
		{37.79264916808487, -122.42616913804746},
	},
	{
		{37.78647766728455, -122.42321282905907},
		{37.78374742447772, -122.42269398620836},
		{37.78300293033823, -122.4225475612199},
		{37.78291035023184, -122.42089295885036},
		{37.78310564009152, -122.41910523857551},
		{37.783340947012974, -122.41759218036147},
		{37.78242075518227, -122.41739703541508},
		{37.78149494160786, -122.41715787460058},
	},
	{
		{37.78149494160786, -122.41715787460058},
		{37.78242075518227, -122.41739703541508},
		{37.783340947012974, -122.41759218036147},
		{37.78310564009152, -122.41910523857551},
		{37.78291035023184, -122.42089295885036},
		{37.78374742447772, -122.42269398620836},
		{37.78647766728455, -122.42321282905907},
		{37.78300293033823, -122.4225475612199},
	},
}

type Driver struct {
	DriverId string
	Route    [][]float64
	Index    int // Current position in the route
}

type Service struct {
	drivers []*Driver
}

func NewService() *Service {
	var drivers []*Driver
	for i, route := range predefinedRoutes {
		drivers = append(drivers, &Driver{
			DriverId: fmt.Sprintf("driver-%d", i),
			Route:    route,
			Index:    0,
		})
	}

	return &Service{
		drivers,
	}
}

func (s *Service) FindNearbyDrivers() []*pb.Driver {
	var drivers []*pb.Driver

	for _, driver := range s.drivers {
		// Get current position and move to the next waypoint
		lat := driver.Route[driver.Index][0]
		lng := driver.Route[driver.Index][1]
		driver.Index = (driver.Index + 1) % len(driver.Route) // Loop back at end of route

		// Update driver protobuf object
		pbDriver := &pb.Driver{
			DriverId: driver.DriverId,
			Location: &pb.Location{
				Latitude:  float32(lat),
				Longitude: float32(lng),
			},
			Geohash: geohash.Encode(lat, lng),
		}
		drivers = append(drivers, pbDriver)
	}

	return drivers
}
