// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

package main
import (
	"encoding/json"
)

func UnmarshalWelcome(data []byte) (ScheduledFlight, error) {
	var r ScheduledFlight
	err := json.Unmarshal(data, &r)
	return r, err
}



type Payload struct {   
	ScheduledFlights []ScheduledFlight  `json:"scheduledFlights"`    
}

type ScheduledFlight struct {
	CarrierFSCode           string        `json:"carrierFsCode"`          
	FlightNumber            string        `json:"flightNumber"`           
	DepartureAirportFSCode  string        `json:"departureAirportFsCode"` 
	ArrivalAirportFSCode    string        `json:"arrivalAirportFsCode"`   
}
