package routes

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
	"strings"
)

func GetRouteId(routeName string) interface{} {
	client := &http.Client{
	}
	getRoutesEndPoint := "http://svc.metrotransit.org/NexTrip/Routes"
	reqRoutes, err := http.NewRequest("GET",getRoutesEndPoint , nil)
	reqRoutes.Header.Add("Accept", "application/json")

	resRoutes, err := client.Do(reqRoutes)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(resRoutes.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result []map[string]interface{}
	json.Unmarshal([]byte(responseData), &result)

	var routeId interface{}

	for _, value := range result {

		if strings.Contains(value["Description"].(string),routeName)  {
			routeId = value["Route"]
			break
		}

	}
	return routeId
}