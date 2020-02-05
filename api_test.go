package main

import (
	"testing"
)

func TestFetchResponse(t  *testing.T){
	urlForTest := FetchResponse("https://api.darksky.net/forecast/58ffcf3ca186b5068bc9918ad2c16d8e/9.55,44.050000/")
	var resExample string = "exampleResponse.json"

	if urlForTest != resExample {
		t.Errorf("response is incorrect")
	}


}