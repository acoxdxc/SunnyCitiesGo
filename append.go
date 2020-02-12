package main

import (
	"fmt"
	"encoding/json"
	"os"
)

func appendCities() ([]string) {
	var listOfPayloadFilenames []string = listOfURL()
   var _flightCodes []string
	for _, payloadFileName := range listOfPayloadFilenames{
		res := FetchResponse(payloadFileName)
		body := ResponseBodyToByte(res)


		var data Payload
		jsonErr := json.Unmarshal(body, &data)
		if jsonErr != nil {
			fmt.Println(jsonErr)
			os.Exit(1)
		}

		var flights []ScheduledFlight = data.ScheduledFlights
		for _, flight := range flights {
			var flightNumber string = flight.FlightNumber
			//fmt.Println(flightNumber)
			_flightCodes = append(_flightCodes, flightNumber)

		}

		
		
		//fmt.Println(body)
		//fmt.Println(data)
		//fmt.Println(jsonErr)
		//fmt.Println(data.CarrierFSCode) 
	}

	return _flightCodes
}


