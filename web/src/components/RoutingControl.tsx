import {useEffect, useState} from "react";
import {Polyline} from "react-leaflet";
import {RequestRideProps} from "@/types";


export function RoutingControl({start, destination, onRequestRide}: {
    onRequestRide: (props: RequestRideProps) => Promise<any>,
    start: [number, number],
    destination: [number, number] | null
}) {
    const [route, setRoute] = useState<[number, number][]>([])

    useEffect(() => {
        if (!destination) return

        const fetchRoute = async () => {
            const response = await onRequestRide({
                pickup: start,
                destination,
            })
            const data = await response.json()
            console.log(data)
            if (data.routes && data.routes.length > 0) {
                setRoute(data.routes[0].geometry.coordinates.map((coord: number[]) => [coord[1], coord[0]]))
            }
        }

        fetchRoute()
    }, [destination, start])

    if (!route) return null

    return <Polyline positions={route} color="blue"/>
}