package main

import (
	"encoding/json"
	"fmt"
	
)

var stats Payload

func unmarshalPractice() {

	url := "https://api.flightstats.com/flex/schedules/rest/v1/json/from/LCY/to/TXL/departing/2020/01/23?appId=95afa7e3&appKey=0d1c89f19008a8e2c69a378f31985ef3"
	res := FetchResponse(url)
	body := ResponseBodyToByte(res)
	json.Unmarshal(body, &stats)
	fmt.Println(stats)

}