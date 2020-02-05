package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)



type FlightCodeList struct{
	Name string
	Iata string
	City string

}

func mockCities()([6]string) {

	var mockWeather[6] string	
	mockWeather[0] = "Lisbon"
	mockWeather[1] = "Paris"
	mockWeather[2] = "Moscow"
	mockWeather[3] = "New York"
	mockWeather[4] = "Rome"
	mockWeather[5] = "Madrid"
	//fmt.Println(mockWeather)
	return mockWeather
}


var baseFlightURL string = "https://api.flightstats.com/flex/schedules/rest/v1/json/from/"
var departure string
var arrival string
var date string
//year/month/day

var fullFlightURL string = baseFlightURL + departure +"/to/" + arrival +"/departing/" + date +"?appId=5b95cdc1&appKey=20e09adffd6fc91daefcac0bceaf6755"


func unmarshalFlight()([]FlightCodeList){
	listOfCodes := make([]FlightCodeList,0)
	file, _ := ioutil.ReadFile("airports.json")
	err2 := json.Unmarshal(file, &listOfCodes)
	if err2 != nil {
		fmt.Println(err2)
			os.Exit(1)
		}
	//fmt.Println(listOfCodes)
	return listOfCodes
	
	}

func getFlightCodes()([]string) {
	var codesNeededForCities []string
	var finalCodesForCities []string
	cities := mockCities()
	codes := unmarshalFlight()

	for _,city := range cities{
		for _,code := range codes{
			if (city == code.City){
				codesNeededForCities = append(codesNeededForCities, code.Iata)
			}
		}
	}

	//codesNeededForCities contains empty spaces in the array. This loop returns finalCodesForCities, which returns
	//an array without these gaps
	for _, str := range codesNeededForCities {
        if str != "" {
            finalCodesForCities = append(finalCodesForCities, str)
        }
    }
    
	//fmt.Println(finalCodesForCities)
	return finalCodesForCities
	}

func listOfURL()([]string) {
	var flightUrlList []string
	flightList := getFlightCodes()
	departure = "YXU"
	var date string = "2020/03/15"
	for _,item := range flightList{
		arrival = item
		var fullFlightURL string = baseFlightURL + departure +"/to/" + arrival +"/departing/" + date +"?appId=5b95cdc1&appKey=20e09adffd6fc91daefcac0bceaf6755"
		
		flightUrlList = append(flightUrlList, fullFlightURL)
	}
	
	return flightUrlList
}

//var data ScheduledFlight
func UnmarshalBodyToPointerFlight(Body []byte, welcome *ScheduledFlight) {
    err2 := json.Unmarshal(Body, &welcome)
    if err2 != nil {
        fmt.Println(err2)
        os.Exit(1)
    }
}

func GetFlightData(url string) ([]ScheduledFlight) {
    res := FetchResponse(url)
	body := ResponseBodyToByte(res)
	data := make([]ScheduledFlight,0)
    json.Unmarshal(body, &data)
    return data
}

func oneFlightCode() (*Payload){
	var testURL string = "https://api.flightstats.com/flex/schedules/rest/v1/json/from/LCY/to/TXL/departing/2020/02/15?appId=95afa7e3&appKey=0d1c89f19008a8e2c69a378f31985ef3"
	
	res := FetchResponse(testURL)
	body := ResponseBodyToByte(res)
	var data *Payload
	json.Unmarshal(body, &data)

	fmt.Println(data)
	
	return data
}

func UnmarshalFlightStatsURL() []string {
    urlList := listOfURL()
    var flightsToSunnyCities []string
    for _, item := range urlList {

        var flightStats []ScheduledFlight = GetFlightData(item)
        fmt.Println(flightStats)

    }
    fmt.Println(flightsToSunnyCities)
    return flightsToSunnyCities
}

func oneFlightCodeMock() (*Payload){
	
	
	body, _ := ioutil.ReadFile("mockAPI_2.json")
	
	var data *Payload
	json.Unmarshal(body, &data)

	fmt.Println(data)
	
	return data
}

func multipleFlightMock() ([5]string) {
	 listOfJson := [5]string {"mockAPI_1.json", "mockAPI_2.json", "mockAPI_3.json", "mockAPI_4.json", "mockAPI_5.json"}

	 for _, item := range listOfJson{
		 var data *Payload 
		 body, _ := ioutil.ReadFile(item)
		 json.Unmarshal(body, &data)
		 fmt.Println(data)
		 
	 }
	 return listOfJson
}

func multipleFlightMockAppend() ([]Payload) {
	listOfJson := [5]string {"mockAPI_1.json", "mockAPI_2.json", "mockAPI_3.json", "mockAPI_4.json", "mockAPI_5.json"}
   var flightCodes []Payload
	for _, item := range listOfJson{
		var data Payload
		//var final string = data.FlightNumber
		
		body, _ := ioutil.ReadFile(item)
		json.Unmarshal(body, &data)
		flightCodes = append(flightCodes, data)
		
	}
	fmt.Println("The flight codes to sunny locations are:" , flightCodes)
	return flightCodes
}

func multipleFlightMockAppend2() ([]string) {
	listOfPayloadFilenames := [5]string {"mockAPI_1.json", "mockAPI_2.json", "mockAPI_3.json", "mockAPI_4.json", "mockAPI_5.json"}
   var _flightCodes []string
	for _, payloadFileName := range listOfPayloadFilenames{
		body, err := ioutil.ReadFile(payloadFileName)
		if err != nil {
	   		 fmt.Println(err)
			 os.Exit(1)
		}


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

func displayFlightCodes() {
	var flightCodes []string = multipleFlightMockAppend2()
	fmt.Println(flightCodes)

}


//To do next
// Data is given in format: DepartureLocationCode / FlightCode / ArrivalLocationCode
//Append list to show: Flight Codes from (DepartureLocationCode) to (ArrivialLocationCode) are (FlightCodes)