package shared

type Geometry struct {
	Coordinates [][]float64 `json:"coordinates"`
}

type Route struct {
	Geometry []Geometry `json:"geometry"`
}

type RouteInfo struct {
	Routes   []Route `json:"routes"`
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	// TODO: Add other fields
}
