export interface RequestRideProps {
    pickup: [number, number],
    destination: [number, number],
}

export interface Location {
    latitude: number,
    longitude: number,
}

export interface RouteInfo {
    route: {
        geometry: {
            coordinates: Location[]
        }[]
    }
}