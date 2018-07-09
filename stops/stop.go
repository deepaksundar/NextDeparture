package stops

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
	"strings"
)

func GetStopId(routeId interface{}, stopName string, direction string ) interface{}{
	client := &http.Client{
	}
	getStopsByRouteEndPoint := "http://svc.metrotransit.org/NexTrip/Stops/"+routeId.(string)+"/"+direction
	reqStops, err := http.NewRequest("GET", getStopsByRouteEndPoint, nil)
	reqStops.Header.Add("Accept", "application/json")
	resStops, err := client.Do(reqStops)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	stopsResBody, err := ioutil.ReadAll(resStops.Body)
	if err != nil {
		log.Fatal(err)
	}

	var stops []map[string]interface{}
	json.Unmarshal([]byte(stopsResBody), &stops)

	var stopId interface{}

	for _, value := range stops {

		if strings.Contains(value["Text"].(string), stopName) {
			stopId = value["Value"]
			break
		}
	}
	return stopId
}
