package main

import (
	"fmt"
	"os"
	"time"
	"nextdeparture/routes"
	"nextdeparture/stops"
	"nextdeparture/timings"
)





func main() {

	/*routeName := "METRO Blue Line"
	direction := "south"
	stopName := "Target Field Station Platform 1"*/

	routeName := os.Args[1]
	stopName  := os.Args[2]
	var direction string
	switch os.Args[3] {
	case "south" :
		direction = "1"
	case "east" :
		direction = "2"
	case "north" :
		direction = "3"
	case "west" :
		direction = "4"
	}

	routeId := routes.GetRouteId(routeName)
	stopId := stops.GetStopId(routeId,stopName,direction)
	earliestTime := timings.GetEarliestTime(routeId,stopId,direction)

	nextDeparture := int(earliestTime.Sub(time.Now().UTC()).Minutes())
	if nextDeparture == 0 {
		fmt.Println(int(earliestTime.Sub(time.Now().UTC()).Seconds()), "Seconds")
	}else{
		fmt.Println(nextDeparture, "Minutes")
	}
}
