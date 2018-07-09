package timings

import (
	"time"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"encoding/xml"
)

type Departures struct {
	XMLName xml.Name `xml:"ArrayOfNexTripDeparture"`
	DepartureTimes    []DepartureTime    `xml:"NexTripDeparture>DepartureTime"`
}

type DepartureTime string

func GetEarliestTime(routeId interface{}, stopId interface{}, direction string) time.Time{
	client := &http.Client{
	}
	getSchedulesEndpoint := "http://svc.metrotransit.org/NexTrip/"+routeId.(string)+"/"+direction+"/"+stopId.(string)
	reqDepartures, err := http.NewRequest("GET", getSchedulesEndpoint, nil)
	reqDepartures.Header.Add("Accept", "application/xml")
	resDepartures, err := client.Do(reqDepartures)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	departuresResBody, err := ioutil.ReadAll(resDepartures.Body)
	if err != nil {
		log.Fatal(err)
	}


	var departures Departures
	xml.Unmarshal([]byte(string(departuresResBody)), &departures)


	var earliestTime time.Time = time.Date(0001, 01, 01, 0, 0, 0, 0,time.UTC)

	for _, value := range departures.DepartureTimes {
		input := string(value) + "-0500"
		layout := "2006-01-02T15:04:05-0700"
		dTime, _ := time.Parse(layout, input)

		if !earliestTime.Equal(time.Date(0001, 01, 01, 0, 0, 0, 0, time.UTC)) {
			if dTime.Before(earliestTime) &&
				dTime.After(time.Now().UTC()) &&
				dTime.Day() == time.Now().UTC().Day() &&
				dTime.Month() == time.Now().UTC().Month() &&
				dTime.Year() == time.Now().UTC().Year(){
				earliestTime = dTime
			}
		} else {
			if dTime.After(time.Now().UTC()) &&
				dTime.Day() == time.Now().UTC().Day() &&
				dTime.Month() == time.Now().UTC().Month() &&
				dTime.Year() == time.Now().UTC().Year(){
				earliestTime = dTime
			}
		}
	}
	return earliestTime
}