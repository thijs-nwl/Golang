package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//The first step is to use golang’s http module to get the response
	res, err := http.Get("https://www.citibikenyc.com/stations/json")
	if err != nil {
		panic(err.Error())
	}

	//Assuming you didnt see a panic call, the response to this http call
	//is being stored in the res variable. Next, we need to read the http body
	//into a byte array for parsing/processing (using golangs ioutil library)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	s, err := getStations([]byte(body))
	fmt.Print(s.StationBeanList[0])
}

func getStations(body []byte) (*StationAPIResponse, error) {
	var s = new(StationAPIResponse)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

//Now that you have the http body, we are going to extract the JSON
//parameters into go structs so they can be easily accessed in other go
//programs. There are many different ways to do this, but I am going to use
//two separate golang structures Station and StationAPIResponse
type Station struct {
	Id                    int64   `json:"id"`
	StationName           string  `json:"stationName"`
	AvailableDocks        int64   `json:"availableDocks"`
	TotalDocks            int64   `json:"totalDocks"`
	Latitude              float64 `json:"latitude"`
	Longitude             float64 `json:"longitude"`
	StatusValue           string  `json:"statusValue"`
	StatusKey             int64   `json:"statusKey"`
	AvailableBikes        int64   `json:"availableBikes"`
	StAddress1            string  `json:"stAddress1"`
	StAddress2            string  `json:"stAddress2"`
	City                  string  `json:"city"`
	PostalCode            string  `json:"postalCode"`
	Location              string  `json:"location"`
	Altitude              string  `json:"altitude"`
	TestStation           bool    `json:"testStation"`
	LastCommunicationTime string  `json:"lastCommunicationTime"`
	LandMark              string  `json:"landMark"`
}

type StationAPIResponse struct {
	ExecutionTime   string    `json:"executionTime"`
	StationBeanList []Station `json:"stationBeanList"`
}
